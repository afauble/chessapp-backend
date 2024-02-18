package data

type ChessGame struct {
	gameType        GameType
	playerOne       User
	playerTwo       User
	currentPosition BoardPosition
	moves           []string
}

type GameType struct {
	variant    string // maybe enum?
	timeFormat string // maybe enum?
}
