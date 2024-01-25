package game

type CreateGameDto struct {
	Seed       *int           `json:"seed,omitempty"`
	Difficulty GameDifficulty `json:"difficulty"`
	Rows       *int           `json:"rows,omitempty"`
	Cols       *int           `json:"cols,omitempty"`
	Mines      *int           `json:"mines,omitempty"`
}
