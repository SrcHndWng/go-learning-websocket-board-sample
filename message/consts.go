package message

const (
	// KindConnected is sent when user connects
	KindConnected = iota + 1
	// KindUserJoined is sent when someone else joins
	KindUserJoined
	// KindUserLeft is sent when someone leaves
	KindUserLeft
	// KindStroke message specifies a drawn stroke by a user
	KindStroke
	// KindClear message is sent when a user clears the screen
	KindClear
	// KindColorSelect message is sent when a user selects color
	KindColorSelect
	// KindText messag is sent when a user input text
	KindText
)
