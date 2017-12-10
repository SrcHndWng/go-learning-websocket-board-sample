package message

type Connected struct {
	Kind  int    `json:"kind"`
	Color string `json:"color"`
	Users []User `json:"users"`
}

func NewConnected(color string, users []User) *Connected {
	return &Connected{
		Kind:  KindConnected,
		Color: color,
		Users: users,
	}
}
