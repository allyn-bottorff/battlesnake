package main

import (
	"log"
	"math"
	"math/rand"
)

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

func checkForBody(body []Coord, possMoves map[string]bool) {
	head := body[0]

	// check up
	for i := 1; i < len(body); i++ {
		if head.X == body[i].X && head.Y+1 == body[i].Y {
			possMoves["up"] = false
			break
		}
	}

	// check down
	for i := 1; i < len(body); i++ {
		if head.X == body[i].X && head.Y-1 == body[i].Y {
			possMoves["down"] = false
			break
		}
	}

	// check right

	for i := 1; i < len(body); i++ {
		if head.X+1 == body[i].X && head.Y == body[i].Y {
			possMoves["right"] = false
			break
		}
	}

	// check left
	for i := 1; i < len(body); i++ {
		if head.X-1 == body[i].X && head.Y == body[i].Y {
			possMoves["left"] = false
			break
		}
	}

}

func checkForWalls(head Coord, board Board, possMoves map[string]bool) {
	// check up
	if possMoves["up"] != false {
		if head.Y+1 >= board.Height-1 {
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
		if head.X+1 >= board.Width-1 {
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

	// Step 0: Don't let your Battlesnake move back in on it's own neck
	//myHead := state.You.Body[0] // Coordinates of your head
	//myNeck := state.You.Body[1] // Coordinates of body piece directly behind your head (your "neck")
	//if myNeck.X < myHead.X {
	//	possibleMoves["left"] = false
	//} else if myNeck.X > myHead.X {
	//	possibleMoves["right"] = false
	//} else if myNeck.Y < myHead.Y {
	//	possibleMoves["down"] = false
	//} else if myNeck.Y > myHead.Y {
	//	possibleMoves["up"] = false
	//}

	checkForBody(state.You.Body, possibleMoves)

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
