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

        if e != "JSON" {
            totalFiles += 1
            _, ok := allHashes[hash]
            if ok {
                allHashes[hash] += 1
            } else {
                allHashes[hash] = 1
            }
        }

        // fmt.Printf("file=%s, h.Size=%d, hash=%s\n", h.Name, h.Size, hash)
    }

    fmt.Printf("totalFiles=%d, uniqueHashes=%\n", totalFiles, len(allHashes))

    return
}

func main() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    log.Printf("==============================================")
    defer log.Printf("==============================================")

    // Usage: main takeout.tgz
    // TODO - use an actual flag parser.
    // Likely, `main -i takeout.tgz -d path/to/db.sqlite`
    tarPath := os.Args[1]
    fmt.Println(tarPath)

    tarFile, err := os.Open(tarPath)
    if err != nil {
        panic(err)
    }
    defer tarFile.Close()

    ParseTakeoutGZIP(tarFile)
}