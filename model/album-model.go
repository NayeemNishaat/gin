package model

// album represents data about a record album.
type Album struct {
	ID     string  `json:"id"` // Note: JSON tag
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
	URL    string  `json:"url"`
}

// albums slice to seed record album data.
var Albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99, URL: "https://i.scdn.co/image/ab67616d0000b273611ea3fb281f7956ffd33b77"},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99, URL: "https://i.scdn.co/image/ab67616d0000b2734d370365fffc95657ee5b1d6"},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99, URL: "https://i.scdn.co/image/ab67616d0000b2733cd4246cb09d8222c3d61106"},
}
