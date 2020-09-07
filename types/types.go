package types

import (
    "gorm.io/gorm"
)

type Album struct {
    gorm.Model
    Name string `gorm:"unique"`
}

type Photo struct {
    Hash      string `gorm:"primaryKey"`
    Extension string
    Year      string
    Month     string
    Day       string
    Name      string
    SizeBytes int64
    Albums    []Album `gorm:"many2many:photo_albums;"`

    // Metadata
    Trashed          bool
    CreationTime     string
    ModificationTime string
    PhotoTakenTime   string

    // Geodata
    Latitude      string
    Longitude     string
    Altitude      string
    LatitudeSpan  string
    LongitudeSpan string
}
