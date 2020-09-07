package main

import (
    "archive/tar"
    "compress/gzip"
    "crypto/sha1"
    "encoding/hex"
    "fmt"
    "io"
    "log"
    "os"
    "path/filepath"
    "regexp"
    "strings"

    "github.com/karan/gphotos-takeout/db"
    "github.com/karan/gphotos-takeout/types"
)

var (
    dateAlbumRegexp = regexp.MustCompile(`^([0-9]{4})-([0-9]{2})-([0-9]{2})$`)
)

func ComputeHash(reader io.Reader) string {
    sha := sha1.New()
    if _, err := io.Copy(sha, reader); err != nil {
        log.Fatal(err)
    }
    return hex.EncodeToString(sha.Sum(nil))
}

func CreateOrUpdatePhoto(reader io.Reader, h *tar.Header, dbConn *db.Connection) *types.Photo {
    p := types.Photo{}
    log.Printf("File: %q", h.Name)

    e := filepath.Ext(h.Name)
    if e == ".json" {
        // TODO: Handle JSON type separately
        // Takeout/Google Photos/2020-08-31/IMG_20200831_081118.jpg.json
        // Read EXIF metadata and save it in a take
        // Directory and filename are unique enough to match with a photo
        // trashed: true - delete this
        // creationTime, modificationTime, geoData, geoDataExif, photoTakenTime
        // Build a struct for metadata metadata, save in sqlite if not already present.
        return nil
    }

    hash := ComputeHash(reader)

    p = dbConn.FindPhoto(hash)
    p.Extension = e

    p.Hash = hash
    p.SizeBytes = h.Size

    // Eg: Takeout/Google Photos/2020-08-31/IMG_20200831_071508.jpg
    parts := strings.Split(h.Name, "/")
    p.Name = parts[len(parts)-1]
    albumName := parts[2]
    albumCapture := dateAlbumRegexp.FindStringSubmatch(albumName)

    log.Printf("albumCapture=%+v", albumCapture)
    if len(albumCapture) == 4 {
        // Photo type (not a user-created album)
        p.Year = albumCapture[1]
        p.Month = albumCapture[2]
        p.Day = albumCapture[3]
    } else {
        // Custom album
        // TODO: Parse date created time based on file
        p.Albums = append(p.Albums, types.Album{Name: albumName})
    }
    return &p
}

func ParseTakeoutGZIP(dbConn *db.Connection, reader io.Reader) (err error) {
    greader, err := gzip.NewReader(reader)
    if err != nil {
        return
    }

    totalFiles := 0
    treader := tar.NewReader(greader)

    for h, err := treader.Next(); err == nil; h, err = treader.Next() {
        p := CreateOrUpdatePhoto(treader, h, dbConn)
        if p != nil {
            dbConn.InsertPhoto(p)
        }
        totalFiles += 1
    }

    fmt.Printf("totalFiles=%d\n", totalFiles)

    return
}

func main() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    log.Printf("==============================================")
    defer log.Printf("==============================================")

    // Create Database
    dbConn, err := db.CreateDB("database.db")
    if err != nil {
        panic("failed to connect database")
    }

    // Open and process tar ball
    // Usage: main takeout.tgz
    // TODO - use an actual flag parser: `main -i takeout.tgz -d path/to/db.sqlite`
    tarPath := os.Args[1]
    fmt.Println(tarPath)

    tarFile, err := os.Open(tarPath)
    if err != nil {
        panic(err)
    }
    defer tarFile.Close()

    ParseTakeoutGZIP(dbConn, tarFile)
}
