package message

type UserJoined struct {
	Kind int  `json:"kind"`
	User User `json:"user"`
}

func NewUserJoined(userID string, color string) *UserJoined {
	return &UserJoined{
		Kind: KindUserJoined,
		User: User{ID: userID, Color: color},
	}
}
