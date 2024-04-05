package helpers

var standardKnightMoveMap map[int8]uint64
var standardRookMoveMap map[int8]uint64
var standardBishopMoveMap map[int8]uint64

func init() {
	// Initialize Standard Move Maps
	standardKnightMoveMap = make(map[int8]uint64)
	standardRookMoveMap = make(map[int8]uint64)
	standardBishopMoveMap = make(map[int8]uint64)

	// Populate Maps
	createStandardKnightMoves()
	createStandardRookMoves()
	createStandardBishopMoves()
}

func GetStandardKnightMoves(index int8) uint64 {
	return standardKnightMoveMap[index]
}
func GetStandardRookMoves(index int8) uint64 {
	return standardRookMoveMap[index]
}
func GetStandardBishopMoves(index int8) uint64 {
	return standardBishopMoveMap[index]
}

func createStandardKnightMoves() {
	for i := int8(0); i < 64; i++ {
		standardKnightMoveMap[i] = generateKnightMoves(i)
	}
}
func createStandardRookMoves() {
	for i := int8(0); i < 64; i++ {
		standardRookMoveMap[i] = generateRookMoves(i)
	}
}
func createStandardBishopMoves() {
	for i := int8(0); i < 64; i++ {
		standardBishopMoveMap[i] = generateBishopMoves(i)
	}
}

func generateBishopMoves(index int8) uint64 {
	var moveSet [64]bool

	// Moving Up-Left
	newIndex := index + 9
	rank := index / 8
	for isWithinBoard(newIndex) {
		if (newIndex / 8) == (rank + 1) {
			moveSet[newIndex] = true
			rank = newIndex / 8
			newIndex += 9
		} else {
			break
		}
	}

	// Moving Up-Right
	newIndex = index + 7
	rank = index / 8
	for isWithinBoard(newIndex) {
		if (newIndex / 8) == (rank + 1) {
			moveSet[newIndex] = true
			rank = newIndex / 8
			newIndex += 7
		} else {
			break
		}
	}

	// Moving Down-Left
	newIndex = index - 7
	rank = index / 8
	for isWithinBoard(newIndex) {
		if (newIndex / 8) == (rank - 1) {
			moveSet[newIndex] = true
			rank = newIndex / 8
			newIndex -= 7
		} else {
			break
		}
	}

	// Moving Down-Right
	newIndex = index - 9
	rank = index / 8
	for isWithinBoard(newIndex) {
		if (newIndex / 8) == (rank - 1) {
			moveSet[newIndex] = true
			rank = newIndex / 8
			newIndex -= 9
		} else {
			break
		}
	}

	return createBitBoard(moveSet)
}

func generateKnightMoves(index int8) uint64 {
	// The change in index for each of the 8 knight moves
	knightMoves := [8]int8{17, 15, 10, 6, -6, -10, -15, -17}
	// Check if the move changed went to expected row
	// Moves thats would go off the board can cause wrong indexes to be selected
	knightRowChecks := map[int8]int8{
		17: 2, 15: 2, 10: 1, 6: 1, -6: -1, -10: -1, -15: -2, -17: -2,
	}
	var moveSet [64]bool
	var newIndex int8
	var rowChange int8
	for _, move := range knightMoves {
		newIndex = index + move
		rowChange = (newIndex / 8) - (index / 8)
		if isWithinBoard(newIndex) && knightRowChecks[move] == rowChange {
			moveSet[newIndex] = true
		}
	}

	return createBitBoard(moveSet)
}

func generateRookMoves(index int8) uint64 {
	var moveSet [64]bool

	// Checking each of the four directions of the rook at the same time
	indexUp := index
	indexDown := index
	indexLeft := index
	indexRight := index
	// Making sure not to include numbers "off the board"
	var rank int8 = index / 8

	for i := int8(0); i < 8; i++ {
		indexUp += 8
		indexDown -= 8
		indexLeft -= 1
		indexRight += 1

		if isWithinBoard(indexUp) {
			moveSet[indexUp] = true
		}
		if isWithinBoard(indexDown) {
			moveSet[indexDown] = true
		}
		if isWithinBoard(indexLeft) && (indexLeft/8) == rank {
			moveSet[indexLeft] = true
		}
		if isWithinBoard(indexRight) && (indexRight/8) == rank {
			moveSet[indexRight] = true
		}
	}

	return createBitBoard(moveSet)
}

func isWithinBoard(index int8) bool {
	if index < 0 || index > 63 {
		return false
	}
	return true
}
func createBitBoard(moveSet [64]bool) uint64 {
	var moveBitBoard uint64 = 0
	for i := 63; i >= 0; i-- {
		moveBitBoard = moveBitBoard << 1
		if moveSet[i] {
			moveBitBoard += 1
		}
	}
	return moveBitBoard
}
