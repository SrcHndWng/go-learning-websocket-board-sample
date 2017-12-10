package message

type Clear struct {
	Kind   int    `json:"kind"`
	UserID string `json:"userId"`
}
