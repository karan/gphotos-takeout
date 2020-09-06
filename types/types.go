package types

type GeoData struct {
    Latitude      string
    Longitude     string
    Altitude      string
    LatitudeSpan  string
    LongitudeSpan string
}

type Metadata struct {
    Trashed          bool
    CreationTime     string
    ModificationTime string
    PhotoTakenTime   string
    GeoData          GeoData
    GeoDataExif      GeoData
}

type Photo struct {
    Hash       string `gorm:"primaryKey"`
    Extension  string
    Year       int
    Month      int
    Day        int
    Name       string
    Metadata   Metadata
    AlbumNames []string
}
