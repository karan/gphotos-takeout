# gphotos-takeout

Parse a Google Takeout archive for Google Photos and store it in a database.

It's a WIP and contributions are welcome.

## How it works

Reads a Takeout tgz file, and (without extracting it) walks the tree to find properties of each photo that can be saved in the database. This includes optional EXIF metadata that is in a separate json file.

## Setup

TODO

## Build and Run

```
$ go run main.go path/to/takeout.tgz
```
