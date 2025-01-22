package entity

type Game struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	ReleaseDate string  `json:"release_date"`
	CreatedAt   string  `json:"created_at"`
	CategoryID  int     `json:"category_id"`
}
