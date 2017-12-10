package message

type UserLeft struct {
	Kind   int    `json:"kind"`
	UserID string `json:"userId"`
}

func NewUserLeft(userID string) *UserLeft {
	return &UserLeft{
		Kind:   KindUserLeft,
		UserID: userID,
	}
}
