package main

import (
	"log"
	"math"
	"math/rand"
)

func coordInSnake(coord Coord, snake []Coord) bool {
	var exists bool = false
	for i := 0; i < len(snake); i++ {
		if coord == snake[i] {
			exists = true
			break
		}
	}
	return exists
}

func info() BattlesnakeInfoResponse {
	log.Println("INFO")
	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "MazerRackham",
		Color:      "0467d1",
		Head:       "default",
		Tail:       "default",
	}
}

func start(state GameState) {
	log.Printf("%s START\n", state.Game.ID)
}

// This function is called when a game your Battlesnake was in has ended.
// It's purely for informational purposes, you don't have to make any decisions here.
func end(state GameState) {
	log.Printf("%s END\n\n", state.Game.ID)
}

func checkForBodies(body []Coord, board Board, possMoves map[string]bool) {
	head := body[0]

	// check up
	up := Coord{
		X: head.X,
		Y: head.Y + 1,
	}
	if checkForSelf(body, up) == true || checkForSnakes(board, up) == true {
		possMoves["up"] = false
	}

	// check down
	down := Coord{
		X: head.X,
		Y: head.Y - 1,
	}
	if checkForSelf(body, down) == true || checkForSnakes(board, down) == true {
		possMoves["down"] = false
	}

	// check right
	right := Coord{
		X: head.X + 1,
		Y: head.Y,
	}
	if checkForSelf(body, right) == true || checkForSnakes(board, right) == true {
		possMoves["right"] = false
	}

	// check left
	left := Coord{
		X: head.X - 1,
		Y: head.Y,
	}
	if checkForSelf(body, left) == true || checkForSnakes(board, left) == true {
		possMoves["left"] = false
	}

}

func checkForSelf(body []Coord, testP Coord) bool {
	collides := true
	for i := 1; i < len(body); i++ {
		if testP.X == body[i].X && testP.Y == body[i].Y {
			break
		} else {
			collides = false
		}
	}
	return collides
}

func checkForSnakes(board Board, testP Coord) bool {

	collides := true

	for i := 0; i < len(board.Snakes); i++ {
		collides = coordInSnake(testP, board.Snakes[i].Body)
		if collides == true {
			break
		}

	}
	return collides
}

func checkForWalls(head Coord, board Board, possMoves map[string]bool) {
	// check up
	if possMoves["up"] != false {
		if head.Y+1 >= board.Height {
			possMoves["up"] = false
		}
	}
	// check down
	if possMoves["down"] != false {
		if head.Y-1 < 0 {
			possMoves["down"] = false
		}
	}
	// check right
	if possMoves["right"] != false {
		if head.X+1 >= board.Width {
			possMoves["right"] = false
		}
	}

	// check left
	if possMoves["left"] != false {
		if head.X-1 < 0 {
			possMoves["left"] = false
		}
	}

}

func magn(p1 Coord, p2 Coord) int {
	c2 := (p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y)
	c := math.Sqrt(float64(c2))
	return int(c)
}

func moveToFood(head Coord, possMoves map[string]bool, board Board) []string {
	maxDistIdx := 0
	for i := 0; i < len(board.Food); i++ {
		dist := magn(head, board.Food[i])
		if dist > maxDistIdx {
			maxDistIdx = i

		}
	}
	nearFood := board.Food[maxDistIdx]

	xDist := head.X - nearFood.X
	yDist := head.Y - nearFood.Y

	move := ""

	if math.Abs(float64(xDist)) > math.Abs(float64(yDist)) {
		if xDist < 0 {
			move := "left"
		}
	}

}

// This function is called on every turn of a game. Use the provided GameState to decide
// where to move -- valid moves are "up", "down", "left", or "right".
// We've provided some code and comments to get you started.
func move(state GameState) BattlesnakeMoveResponse {

	possibleMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}

	checkForBodies(state.You.Body, state.Board, possibleMoves)

	checkForWalls(state.You.Body[0], state.Board, possibleMoves)

	// TODO: Step 3 - Don't collide with others.
	// Use information in GameState to prevent your Battlesnake from colliding with others.

	// TODO: Step 4 - Find food.
	// Use information in GameState to seek out and find food.

	// Finally, choose a move from the available safe moves.
	// TODO: Step 5 - Select a move to make based on strategy, rather than random.
	var nextMove string

	safeMoves := []string{}
	for move, isSafe := range possibleMoves {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}

	if len(safeMoves) == 0 {
		nextMove = "down"
		log.Printf("%s MOVE %d: No safe moves detected! Moving %s\n", state.Game.ID, state.Turn, nextMove)
	} else {
		nextMove = safeMoves[rand.Intn(len(safeMoves))]
		log.Printf("%s MOVE %d: %s\n", state.Game.ID, state.Turn, nextMove)
	}
	return BattlesnakeMoveResponse{
		Move: nextMove,
	}
}
