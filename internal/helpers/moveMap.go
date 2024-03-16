package helpers

var standardKingMoveMap map[int8]uint64
var standardQueenMoveMap map[int8]uint64
var standardRookMoveMap map[int8]uint64
var standardBishopMoveMap map[int8]uint64
var standardKnightMoveMap map[int8]uint64
var standardPawnMoveMap map[int8]uint64

func init() {
	// Initialize Standard Move Maps
	standardKingMoveMap = make(map[int8]uint64)
	standardKingMoveMap = make(map[int8]uint64)
	standardQueenMoveMap = make(map[int8]uint64)
	standardRookMoveMap = make(map[int8]uint64)
	standardBishopMoveMap = make(map[int8]uint64)
	standardKnightMoveMap = make(map[int8]uint64)
	standardPawnMoveMap = make(map[int8]uint64)

	// Populate Maps
	createStandardKnightMoves()
}

func GetStandardKnightMoves(index int8) uint64 {
	return standardKnightMoveMap[index]
}

func createStandardKnightMoves() {
	for i := int8(0); i < 64; i++ {
		standardKnightMoveMap[i] = generateKnightMoves(i)
	}
}

func generateKnightMoves(index int8) uint64 {

	// The change in index for each of the 8 knight moves
	knightMoves := [8]int8{17, 15, 10, 6, -6, -10, -15, -17}
	var moveSet [64]bool

	for _, move := range knightMoves {
		if isWithinBoard(index + move) {
			moveSet[index+move] = true
		}
	}

	// Create bitboard
	var moveBitBoard uint64 = 0
	for i := 63; i >= 0; i-- {
		if moveSet[i] {
			moveBitBoard += 1
		}
		moveBitBoard = moveBitBoard << 1
	}
	return moveBitBoard
}

func isWithinBoard(index int8) bool {
	if index < 0 || index > 63 {
		return false
	}
	return true
}
