# gphotos-takeout

Parse a Google Takeout archive for Google Photos and store information about photos in a database.

It's a WIP and contributions are welcome.

## How it works

Reads a Takeout tgz file, and (without extracting it) walks the tree to find properties of each photo that can be saved in the database. This includes optional EXIF metadata that is in a separate json file.

The main program parses a tgz archive, and stores information about unique photos in a database (but not the photo content itself). It also stores photo metadata in a separate database.

Additional programs can be written against the resulting database for manipulating photos in the archive. For example, one program could be used to save all unique photos to a directory structure such as `YYYY/MM/DD`, or de-dupe album photos by creating symlinks.

## Setup

TODO

## Build and Run

```
$ go run main.go path/to/takeout.tgz
```
