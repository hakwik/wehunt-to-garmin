wehunt-to-garmin
================

This is an extremely simple and dumb little program that takes a GPX file as exported from [WeHunt](https://wehunt.app/) and attempts to convert it to a GPX file as exported from the Garmin software.

It is certainly possible to import a WeHunt GPX directly into Garmin but that makes all different types of locations/waypoints get the same icon in the Garmin software and handhelds. This make navigation in the woods a hazzle. This piece of software attempts to calm that particular itch.

## Building
`go build .`

That's it. To get fancy and limit file size, run `go build -ldflags "-s -w" .`

## Running
Simply run the program and pass the name of the export WeHunt GPX as input. The program will create a new filename for the processed file.

## Dependencies
None outside the Go stdlib.
