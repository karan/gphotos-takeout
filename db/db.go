package db

import (
    "log"

    "github.com/karan/gphotos-takeout/types"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type Connection struct {
    dbConn *gorm.DB
}

func CreateDB(name string) (*Connection, error) {
    log.Println("creating db")
    var err error
    dbConn, err := gorm.Open(sqlite.Open(name), &gorm.Config{
        SkipDefaultTransaction: true,
    })
    if err != nil {
        return nil, err
    }
    log.Println("migrating db")
    dbConn.AutoMigrate(&types.Photo{})

    return &Connection{dbConn: dbConn}, nil
}

func (c *Connection) FindPhoto(hash string) types.Photo {
    var p types.Photo
    c.dbConn.First(&p, hash)
    return p
}

func (c *Connection) InsertPhoto(photo *types.Photo) {
    c.dbConn.Save(photo)
}
