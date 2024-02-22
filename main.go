package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"time"
)

var (
	verbose bool
)

func init() {
	flag.BoolVar(&verbose, "v", false, "use verbose logging")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Printf("%s INPUT-FILE [where INPUT-FILE is a GPX file as exported from WeHunt]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	filename := flag.Arg(0)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("could not open file:", err)
		os.Exit(1)
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("could not read file:", err)
		os.Exit(1)
	}

	var gpx *Gpx

	err = xml.Unmarshal(content, &gpx)
	if err != nil {
		fmt.Println("could not unmarshal xml:", err)
		os.Exit(1)
	}

	gpx.Garminify()

	for i, w := range gpx.Wpt {
		w.Garminify()
		if verbose {
			fmt.Println("Garminified waypoint", i+1)
		}
	}

	x, err := xml.MarshalIndent(gpx, "", "    ")
	xstr := []byte("<?xml version=\"1.0\"?>\n" + string(x))
	if err != nil {
		fmt.Println("could not marshal XML:", err)
	}

	t := time.Now().Format("2006-01-02_15-04-05")
	out := t + "-garminified-" + filename
	outfile, err := os.Create(out)
	if err != nil {
		fmt.Println("Could not create output file", outfile)
		os.Exit(1)
	}
	defer outfile.Close()
	outfile.Write(xstr)
	fmt.Println("Wrote file", out)
}

func findMax(values []string) string {
	return fmt.Sprintf("%f", slices.Max(extractAndParse(values)))
}

func findMin(values []string) string {
	return fmt.Sprintf("%f", slices.Min(extractAndParse(values)))
}

func extractAndParse(values []string) []float64 {
	floats := make([]float64, 0)
	for _, strval := range values {
		val, err := strconv.ParseFloat(strval, 32)
		if err != nil {
			panic("could not convert string to float: " + err.Error())
		}
		floats = append(floats, val)
	}
	return floats
}
