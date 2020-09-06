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

    _ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

// TODO: create a sqlite database

func ComputeHash(reader io.Reader) string {
    sha := sha1.New()
    if _, err := io.Copy(sha, reader); err != nil {
        log.Fatal(err)
    }
    return hex.EncodeToString(sha.Sum(nil))
}

func ParseTakeoutGZIP(reader io.Reader) (err error) {
    greader, err := gzip.NewReader(reader)
    if err != nil {
        return
    }

    // allExts := make(map[string]int)
    allHashes := make(map[string]int)
    totalFiles := 0
    treader := tar.NewReader(greader)
    for h, err := treader.Next(); err == nil; h, err = treader.Next() {
        e := filepath.Ext(h.Name)
        // _, ok := allExts[e]
        // if ok {
        //     allExts[e] += 1
        // } else {
        //     allExts[e] = 1
        // }

        hash := ComputeHash(treader)

        // TODO: handle photo type
        // Takeout/Google Photos/2020-08-31/IMG_20200831_071508.jpg
        // Calculate hash, size etc etc
        // Save it in database

        // TODO: Handle Instant Upload?
        // Takeout/Google Photos/Instant Upload/IMG_20200831_073924.jpg

        // TODO: save album names?

        // TODO: Parse date created time based on file and dir name

        // Build a struct for all photo metadata, save in sqlite if hash isn't already present.

        if e != "JSON" {
            totalFiles += 1
            _, ok := allHashes[hash]
            if ok {
                allHashes[hash] += 1
            } else {
                allHashes[hash] = 1
            }
        }

        // TODO: Handle JSON type separately
        // Takeout/Google Photos/2020-08-31/IMG_20200831_081118.jpg.json
        // Read EXIF metadata and save it in a take
        // Directory and filename are unique enough to match with a photo
        // trashed: true - delete this
        // creationTime, modificationTime, geoData, geoDataExif, photoTakenTime
        // Build a struct for metadata metadata, save in sqlite if not already present.

        // fmt.Printf("file=%s, h.Size=%d, hash=%s\n", h.Name, h.Size, hash)
    }

    fmt.Printf("totalFiles=%d, uniqueHashes=%d\n", totalFiles, len(allHashes))

    return
}

func main() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    log.Printf("==============================================")
    defer log.Printf("==============================================")

    // Create Database
    sqliteDatabase := CreateDB("database.db")
    defer sqliteDatabase.Close()
    CreateTable(sqliteDatabase)

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

    ParseTakeoutGZIP(tarFile)
}
