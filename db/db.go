package db

import (
    "log"

    "github.com/karan/gphotos-takeout/types"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func CreateDB(name string) (*gorm.DB, error) {
    log.Println("creating db")
    db, err := gorm.Open(sqlite.Open(name), &gorm.Config{
        SkipDefaultTransaction: true,
    })
    log.Println("migrating db")
    db.AutoMigrate(&types.Photo{})
    return db, err
}

func InsertPhoto(db *gorm.DB, photo *types.Photo) *gorm.DB {
    return db.Create(&photo)
}
