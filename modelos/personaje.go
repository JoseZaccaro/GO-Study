package modelos

import "time"

type Personaje struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Status   string    `json:"status"`
	Species  string    `json:"species"`
	Type     string    `json:"type"`
	Gender   string    `json:"gender"`
	Origin   Location  `json:"origin"`
	Location Location  `json:"location"`
	Image    string    `json:"image"`
	Episode  []string  `json:"episode"`
	URL      string    `json:"url"`
	Created  time.Time `json:"created"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Personajes []Personaje

type Response struct {
	Info    Info       `json:"info"`
	Results Personajes `json:"results"`
}

type Info struct {
	Count int    `json:"count"`
	Pages int    `json:"pages"`
	Next  string `json:"next"`
	Prev  string `json:"prev"`
}
