package models

type BoardPosition struct {
	bPawn   int64
	bRook   int64
	bKnight int64
	bBishop int64
	bQueen  int64
	bKing   int64

	wPawn   int64
	wRook   int64
	wKnight int64
	wBishop int64
	wQueen  int64
	wKing   int64

	whiteTurn          bool
	castlingRightWhite bool
	castlingRightBlack bool
	enPassantSquare    int
}
