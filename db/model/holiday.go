package model

import "time"

type Country struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Datetime struct {
	Year   int `json:"year"`
	Month  int `json:"month"`
	Day    int `json:"day"`
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
	Second int `json:"second"`
}

type Timezone struct {
	Offset          string `json:"offset"`
	Zoneabb         string `json:"zoneabb"`
	Zoneoffset      int    `json:"zoneoffset"`
	Zonedst         int    `json:"zonedst"`
	Zonetotaloffset int    `json:"zonetotaloffset"`
}

type Date struct {
	Iso      time.Time `json:"iso"`
	Datetime Datetime  `json:"datetime"`
	Timezone Timezone  `json:"timezone"`
}

type Holiday struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Country     Country  `json:"country"`
	Type        []string `json:"type"`
	Locations   string   `json:"locations"`
	States      string   `json:"states"`
	Date        Date     `json:"date,omitempty"`
}

type Holidays []Holiday

type Response struct {
	Holidays Holidays `json:"holidays"`
}

type Meta struct {
	Code int `json:"code"`
}

type ResponseRest struct {
	Meta     Meta     `json:"meta"`
	Response Response `json:"response"`
}
