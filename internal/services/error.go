package services

type Error struct {
	Msg  string
	Code string
}

func (e Error) Error() string {
	return e.Code + " : " + e.Msg
}
