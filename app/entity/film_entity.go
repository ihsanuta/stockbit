package entity

import (
	"fmt"
	"net/url"
)

type FilmLists struct {
	Search   []DataFilm `json:"Search"`
	Total    string     `json:"totalResults"`
	Response string     `json:"Response"`
}

type DataFilm struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type ParamsFilm struct {
	Title      string `json:"title" form:"title"`
	Pagination int    `json:"page" form:"page"`
}

func (p *ParamsFilm) ToParams() url.Values {
	values := make(url.Values)

	if p.Title != "" {
		values.Add("s", p.Title)
	}

	if p.Pagination > 0 {
		values.Add("page", fmt.Sprintf("%d", p.Pagination))
	}

	return values
}

type DataFilmDetail struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Rated    string `json:"Rated"`
	Released string `json:"Released"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Writer   string `json:"Writer"`
	Actors   string `json:"Actors"`
	Plot     string `json:"Plot"`
	Language string `json:"Language"`
	Country  string `json:"Country"`
	Awards   string `json:"Awards"`
	Poster   string `json:"Poster"`
	Ratings  []struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	} `json:"Ratings"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbID     string `json:"imdbID"`
	Type       string `json:"Type"`
	DVD        string `json:"DVD"`
	BoxOffice  string `json:"BoxOffice"`
	Production string `json:"Production"`
	Website    string `json:"Website"`
	Response   string `json:"Response"`
}
