package entity

type Game struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Genre       string  `json:"genre"`
	ReleaseDate string  `json:"release_date"`
}
