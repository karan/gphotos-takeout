package main

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func createDB(name string) {
    log.Printf("Creating %s...\n", name)
    file, err := os.Create(name)
    if err != nil {
        log.Fatal(err.Error())
    }
    file.Close()
    log.Println("database created")

    sqliteDatabase, _ := sql.Open("sqlite3", "./database.db")
    return sqliteDatabase
}

func CreateTable(db *sql.DB) {
    // TODO: album names
    createTableSQL := `CREATE TABLE photo (
        "photohash" TEXT NOT NULL PRIMARY KEY AUTOINCREMENT,
        "year" integer,
        "month" integer,
        "day" integer,
        "name" TEXT,
        "exif_trashed" INTEGER,
        "exif_creationTime" TEXT,
        "exif_modificationTime" TEXT,
        "exif_photoTakenTime" TEXT,
        "exif_geoData" TEXT,
        "exif_geoDataExif" TEXT,
      );` // SQL Statement for Create Table

    log.Println("Create table...")
    statement, err := db.Prepare(createTableSQL) // Prepare SQL Statement
    if err != nil {
        log.Fatal(err.Error())
    }
    statement.Exec() // Execute SQL Statements
    log.Println("table created")
}
