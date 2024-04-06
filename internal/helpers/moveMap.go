package helpers

var standardKnightMoveMap map[int8]uint64
var standardRookMoveMap map[int8]uint64
var standardBishopMoveMap map[int8]uint64
var standardQueenMoveMap map[int8]uint64
var standardKingMoveMap map[int8]uint64
var standardWhitePawnMoveMap map[int8]uint64
var standardWhitePawnAttackMap map[int8]uint64
var standardBlackPawnMoveMap map[int8]uint64
var standardBlackPawnAttackMap map[int8]uint64

func init() {
	// Initialize Standard Move Maps
	standardKnightMoveMap = make(map[int8]uint64)
	standardRookMoveMap = make(map[int8]uint64)
	standardBishopMoveMap = make(map[int8]uint64)
	standardQueenMoveMap = make(map[int8]uint64)
	standardKingMoveMap = make(map[int8]uint64)
	standardWhitePawnMoveMap = make(map[int8]uint64)
	standardWhitePawnAttackMap = make(map[int8]uint64)
	standardBlackPawnMoveMap = make(map[int8]uint64)
	standardBlackPawnAttackMap = make(map[int8]uint64)

	// Populate Maps
	createStandardKnightMoves()
	createStandardRookMoves()
	createStandardBishopMoves()
	createStandardQueenMoves()
	createStandardKingMoves()
	createStandardWhitePawnMoves()
	createStandardBlackPawnMoves()
	createStandardWhitePawnAttacks()
	createStandardBlackPawnAttacks()

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
func GetStandardQueenMoves(index int8) uint64 {
	return standardQueenMoveMap[index]
}
func GetStandardKingMoves(index int8) uint64 {
	return standardKingMoveMap[index]
}
func GetStandardWhitePawnMoves(index int8) uint64 {
	return standardWhitePawnMoveMap[index]
}
func GetStandardWhitePawnAttacks(index int8) uint64 {
	return standardWhitePawnAttackMap[index]
}
func GetStandardBlackPawnMoves(index int8) uint64 {
	return standardBlackPawnMoveMap[index]
}
func GetStandardBlackPawnAttacks(index int8) uint64 {
	return standardBlackPawnAttackMap[index]
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
func createStandardQueenMoves() {
	for i := int8(0); i < 64; i++ {
		standardQueenMoveMap[i] = generateQueenMoves(i)
	}
}
func createStandardKingMoves() {
	for i := int8(0); i < 64; i++ {
		standardKingMoveMap[i] = generateKingMoves(i)
	}
}
func createStandardWhitePawnMoves() {
	for i := int8(0); i < 64; i++ {
		standardWhitePawnMoveMap[i] = generateWhitePawnMoves(i)
	}
}
func createStandardWhitePawnAttacks() {
	for i := int8(0); i < 64; i++ {
		standardWhitePawnAttackMap[i] = generateWhitePawnAttacks(i)
	}
}
func createStandardBlackPawnMoves() {
	for i := int8(0); i < 64; i++ {
		standardBlackPawnMoveMap[i] = generateBlackPawnMoves(i)
	}
}
func createStandardBlackPawnAttacks() {
	for i := int8(0); i < 64; i++ {
		standardBlackPawnAttackMap[i] = generateBlackPawnAttacks(i)
	}
}

func generateWhitePawnMoves(index int8) uint64 {
	var moveSet [64]bool
	newIndex := index + 8
	if index/8 == 1 {
		moveSet[newIndex+8] = true
	}
	if isWithinBoard(newIndex) {
		moveSet[newIndex] = true
	}
	// Handle promotions elsewhere
	return createBitBoard(moveSet)
}

func generateWhitePawnAttacks(index int8) uint64 {
	var moveSet [64]bool
	newRank := (index / 8) + 1
	newIndexLeft := index + 7
	newIndexRight := index + 9
	if isWithinBoard(newIndexLeft) && (newIndexLeft/8) == newRank {
		moveSet[newIndexLeft] = true
	}
	if isWithinBoard(newIndexRight) && (newIndexRight/8) == newRank {
		moveSet[newIndexRight] = true
	}
	// Handle en passant elsewhere
	return createBitBoard(moveSet)
}

func generateBlackPawnMoves(index int8) uint64 {
	var moveSet [64]bool
	newIndex := index - 8
	if index/8 == 7 {
		moveSet[newIndex-8] = true
	}
	if isWithinBoard(newIndex) {
		moveSet[newIndex] = true
	}
	// Handle promotions elsewhere
	return createBitBoard(moveSet)
}

func generateBlackPawnAttacks(index int8) uint64 {
	var moveSet [64]bool
	newRank := (index / 8) - 1
	newIndexLeft := index - 7
	newIndexRight := index - 9
	if isWithinBoard(newIndexLeft) && (newIndexLeft/8) == newRank {
		moveSet[newIndexLeft] = true
	}
	if isWithinBoard(newIndexRight) && (newIndexRight/8) == newRank {
		moveSet[newIndexRight] = true
	}
	// Handle en passant elsewhere
	return createBitBoard(moveSet)
}

func generateKingMoves(index int8) uint64 {
	// The change in index for each of the 8 king moves
	kingMoves := [8]int8{9, 8, 7, 1, -1, -7, -8, -9}
	// Check if the move changed went to expected rank
	// Moves thats would go off the board can cause wrong indexes to be selected
	kingRankChecks := map[int8]int8{
		9: 1, 8: 1, 7: 1, 1: 0, -1: 0, -7: -1, -8: -1, -9: -1,
	}
	var moveSet [64]bool
	var newIndex int8
	var rankChange int8
	for _, move := range kingMoves {
		newIndex = index + move
		rankChange = (newIndex / 8) - (index / 8)
		if isWithinBoard(newIndex) && kingRankChecks[move] == rankChange {
			moveSet[newIndex] = true
		}
	}
	return createBitBoard(moveSet)
}

func generateQueenMoves(index int8) uint64 {
	bishopMoves := generateBishopMoves(index)
	rookMoves := generateRookMoves(index)

	return (bishopMoves | rookMoves)
}

func generateBishopMoves(index int8) uint64 {
	var moveSet [64]bool

	// Moving Up-Right
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

	// Moving Up-Left
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
	// Check if the move changed went to expected rank
	// Moves thats would go off the board can cause wrong indexes to be selected
	knightRankChecks := map[int8]int8{
		17: 2, 15: 2, 10: 1, 6: 1, -6: -1, -10: -1, -15: -2, -17: -2,
	}
	var moveSet [64]bool
	var newIndex int8
	var rankChange int8
	for _, move := range knightMoves {
		newIndex = index + move
		rankChange = (newIndex / 8) - (index / 8)
		if isWithinBoard(newIndex) && knightRankChecks[move] == rankChange {
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
