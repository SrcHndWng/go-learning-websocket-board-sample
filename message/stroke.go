package message

type Stroke struct {
	Kind   int     `json:"kind"`
	UserID string  `json:"userId"`
	Points []Point `json:"points"`
	Finish bool    `json:"finish"`
	Color  string  `json:"color"`
}
