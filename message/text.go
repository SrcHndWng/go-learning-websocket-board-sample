package message

type Text struct {
	Kind   int    `json:"kind"`
	UserID string `json:"userId"`
	Text   string `json:"text"`
}
