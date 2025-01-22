package entity

type Order struct {
	ID        int     `json:"id"`
	UserID    int     `json:"user_id"`
	GameID    int     `json:"game_id"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
	CreatedAt string  `json:"created_at"`
}
