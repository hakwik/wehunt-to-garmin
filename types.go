package main

import (
	"encoding/xml"
	"time"
)

const (
	dateTimeFormat = "2006-01-02T15:04:05Z"
)

type Wpt struct {
	Lat        string `xml:"lat,attr"`
	Lon        string `xml:"lon,attr"`
	Time       string `xml:"time"`
	Name       string `xml:"name"`
	Cmt        string `xml:"cmt"`
	Desc       string `xml:"desc"`
	Sym        string `xml:"sym"`
	Type       string `xml:"type"`
	Extensions struct {
		WaypointExtension struct {
			DisplayMode string `xml:"gpxx:DisplayMode"`
		} `xml:"gpxx:WaypointExtension"`
		WaypointExtensionZ struct {
			DisplayMode string `xml:"wptx1:DisplayMode"`
		} `xml:"wptx1:WaypointExtension"`
		CreationTimeExtension struct {
			CreationTime string `xml:"ctx:CreationTime"`
		} `xml:"ctx:CreationTimeExtension"`
	} `xml:"extensions"`
}

func (w *Wpt) Garminify() {
	w.Time = time.Now().Format(dateTimeFormat)

	switch w.Type {
	case "pass":
		w.Sym = "Blind"
	case "tower":
		w.Sym = "Tree Stand"
	case "salt_lick":
		w.Sym = "Square Blue"
	case "parking":
		w.Sym = "Parking Area"
	case "bar":
		w.Sym = "Toll Booth"
	case "gathering_place":
		w.Sym = "Campground"
	case "food_place":
		w.Sym = "Pizza"
	default:
		w.Sym = "Flag, Blue"
	}

	w.Type = "user"
	w.Cmt = ""
	w.Desc = ""

	w.Extensions.WaypointExtension.DisplayMode = "SymbolAndName"
	w.Extensions.WaypointExtensionZ.DisplayMode = "SymbolAndName"
	w.Extensions.CreationTimeExtension.CreationTime = time.Now().Format(dateTimeFormat)
}

type Gpx struct {
	XMLName        xml.Name `xml:"gpx"`
	Xmlns          string   `xml:"xmlns,attr"`
	Xsi            string   `xml:"xmlns:xsi,attr"`
	Wptx1          string   `xml:"xmlns:wptx1,attr"`
	Gpxtrx         string   `xml:"xmlns:gpxtrx,attr"`
	Gpxtpx         string   `xml:"xmlns:gpxtpx,attr"`
	Gpxx           string   `xml:"xmlns:gpxx,attr"`
	Trp            string   `xml:"xmlns:trp,attr"`
	Adv            string   `xml:"xmlns:adv,attr"`
	Prs            string   `xml:"xmlns:prs,attr"`
	Tmd            string   `xml:"xmlns:tmd,attr"`
	Vptm           string   `xml:"xmlns:vptm,attr"`
	Ctx            string   `xml:"xmlns:ctx,attr"`
	Gpxacc         string   `xml:"xmlns:gpxacc,attr"`
	Gpxpx          string   `xml:"xmlns:gpxpx,attr"`
	Vidx1          string   `xml:"xmlns:vidx1,attr"`
	Creator        string   `xml:"creator,attr"`
	Version        string   `xml:"version,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Metadata       struct {
		Link struct {
			Chardata string `xml:",chardata"`
			Href     string `xml:"href,attr"`
			Text     string `xml:"text"`
		} `xml:"link"`
		Time   string `xml:"time"`
		Bounds struct {
			Maxlat string `xml:"maxlat,attr"`
			Maxlon string `xml:"maxlon,attr"`
			Minlat string `xml:"minlat,attr"`
			Minlon string `xml:"minlon,attr"`
		} `xml:"bounds"`
	} `xml:"metadata"`
	Wpt []*Wpt `xml:"wpt"`
	Trk struct {
		Name       string `xml:"name"`
		Extensions struct {
			TrackExtension struct {
				DisplayColor string `xml:"gpxx:DisplayColor"`
			} `xml:"gpxx:TrackExtension"`
		} `xml:"extensions"`
		Trkseg struct {
			Trkpt []struct {
				Lat string `xml:"lat,attr"`
				Lon string `xml:"lon,attr"`
			} `xml:"trkpt"`
		} `xml:"trkseg"`
	} `xml:"trk"`
}

func (g *Gpx) Garminify() {
	g.Xsi = "http://www.w3.org/2001/XMLSchema-instance"
	g.Wptx1 = "http://www.garmin.com/xmlschemas/WaypointExtension/v1"
	g.Gpxtrx = "http://www.garmin.com/xmlschemas/GpxExtensions/v3"
	g.Gpxtpx = "http://www.garmin.com/xmlschemas/TrackPointExtension/v1"
	g.Gpxx = "http://www.garmin.com/xmlschemas/GpxExtensions/v3"
	g.Trp = "http://www.garmin.com/xmlschemas/TripExtensions/v1"
	g.Adv = "http://www.garmin.com/xmlschemas/AdventuresExtensions/v1"
	g.Prs = "http://www.garmin.com/xmlschemas/PressureExtension/v1"
	g.Tmd = "http://www.garmin.com/xmlschemas/TripMetaDataExtensions/v1"
	g.Vptm = "http://www.garmin.com/xmlschemas/ViaPointTransportationModeExtensions/v1"
	g.Ctx = "http://www.garmin.com/xmlschemas/CreationTimeExtension/v1"
	g.Gpxacc = "http://www.garmin.com/xmlschemas/AccelerationExtension/v1"
	g.Gpxpx = "http://www.garmin.com/xmlschemas/PowerExtension/v1"
	g.Vidx1 = "http://www.garmin.com/xmlschemas/VideoExtension/v1"
	g.Creator = "Garmin Desktop App"
	g.Version = "1.1"
	g.SchemaLocation = "http://www.topografix.com/GPX/1/1 http://www.topografix.com/GPX/1/1/gpx.xsd http://www.garmin.com/xmlschemas/WaypointExtension/v1 http://www8.garmin.com/xmlschemas/WaypointExtensionv1.xsd http://www.garmin.com/xmlschemas/TrackPointExtension/v1 http://www.garmin.com/xmlschemas/TrackPointExtensionv1.xsd http://www.garmin.com/xmlschemas/GpxExtensions/v3 http://www8.garmin.com/xmlschemas/GpxExtensionsv3.xsd http://www.garmin.com/xmlschemas/ActivityExtension/v1 http://www8.garmin.com/xmlschemas/ActivityExtensionv1.xsd http://www.garmin.com/xmlschemas/AdventuresExtensions/v1 http://www8.garmin.com/xmlschemas/AdventuresExtensionv1.xsd http://www.garmin.com/xmlschemas/PressureExtension/v1 http://www.garmin.com/xmlschemas/PressureExtensionv1.xsd http://www.garmin.com/xmlschemas/TripExtensions/v1 http://www.garmin.com/xmlschemas/TripExtensionsv1.xsd http://www.garmin.com/xmlschemas/TripMetaDataExtensions/v1 http://www.garmin.com/xmlschemas/TripMetaDataExtensionsv1.xsd http://www.garmin.com/xmlschemas/ViaPointTransportationModeExtensions/v1 http://www.garmin.com/xmlschemas/ViaPointTransportationModeExtensionsv1.xsd http://www.garmin.com/xmlschemas/CreationTimeExtension/v1 http://www.garmin.com/xmlschemas/CreationTimeExtensionsv1.xsd http://www.garmin.com/xmlschemas/AccelerationExtension/v1 http://www.garmin.com/xmlschemas/AccelerationExtensionv1.xsd http://www.garmin.com/xmlschemas/PowerExtension/v1 http://www.garmin.com/xmlschemas/PowerExtensionv1.xsd http://www.garmin.com/xmlschemas/VideoExtension/v1 http://www.garmin.com/xmlschemas/VideoExtensionv1.xsd"

	lats := make([]string, 0)
	lons := make([]string, 0)
	for _, wpt := range g.Wpt {
		lats = append(lats, wpt.Lat)
		lons = append(lons, wpt.Lon)
	}

	g.Metadata.Bounds.Maxlat = findMax(lats)
	g.Metadata.Bounds.Minlat = findMin(lats)
	g.Metadata.Bounds.Maxlon = findMax(lons)
	g.Metadata.Bounds.Minlon = findMin(lons)
	g.Metadata.Link.Href = "https://www.garmin.com"
	g.Metadata.Link.Text = "Garmin International"
	g.Metadata.Time = time.Now().Format(dateTimeFormat)

	g.Trk.Extensions.TrackExtension.DisplayColor = "DarkGray"
}
