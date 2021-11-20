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
		Color:      "#0467d1",
		Head:       "shades",
		Tail:       "pixel",
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

func moveToFood(head Coord, board Board) string {
	minDist := 1000
	minDistIdx := 0
	for i := 0; i < len(board.Food); i++ {
		dist := magn(head, board.Food[i])
		if dist < minDist {
			minDist = dist
			minDistIdx = i

		}
	}
	nearFood := board.Food[minDistIdx]

	xDist := nearFood.X - head.X
	yDist := nearFood.Y - head.Y

	move := ""

	if math.Abs(float64(xDist)) > math.Abs(float64(yDist)) {
		if xDist < 0 {
			move = "left"
		} else {
			move = "right"
		}
	} else {
		if yDist < 0 {
			move = "down"
		} else {
			move = "up"
		}
	}
	return move
}

func coordInCoords(p Coord, ps []Coord) bool {
	var exists bool = false
	for i := 0; i < len(ps); i++ {
		if p == ps[i] {
			exists = true
			break
		}

	}
	return exists
}

func checkForLongerSnake(head Coord, ownLen int32, board Board, possMoves map[string]bool) {
	if len(board.Snakes) > 1 {
		for i := 0; i < len(board.Snakes); i++ {

			if ownLen <= board.Snakes[i].Length && head != board.Snakes[i].Body[0] {

				//Possible next moves of a snake that would eat me.
				var badZone = []Coord{
					{
						X: board.Snakes[i].Body[0].X + 1,
						Y: board.Snakes[i].Body[0].Y,
					},
					{
						X: board.Snakes[i].Body[0].X - 1,
						Y: board.Snakes[i].Body[0].Y,
					},
					{
						X: board.Snakes[i].Body[0].X,
						Y: board.Snakes[i].Body[0].Y + 1,
					},
					{
						X: board.Snakes[i].Body[0].X,
						Y: board.Snakes[i].Body[0].Y - 1,
					},
				}

				// check up
				up := Coord{
					X: head.X,
					Y: head.Y + 1,
				}
				if coordInCoords(up, badZone) == true {
					possMoves["up"] = false
				}

				// check down
				down := Coord{
					X: head.X,
					Y: head.Y - 1,
				}
				if coordInCoords(down, badZone) == true {
					possMoves["down"] = false
				}

				// check right
				right := Coord{
					X: head.X + 1,
					Y: head.Y,
				}
				if coordInCoords(right, badZone) == true {
					possMoves["right"] = false
				}

				// check left
				left := Coord{
					X: head.X - 1,
					Y: head.Y,
				}
				if coordInCoords(left, badZone) == true {
					possMoves["left"] = false
				}

			}
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

	var nextMove string

	safeMoves := []string{}
	for move, isSafe := range possibleMoves {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}
	checkForLongerSnake(state.You.Body[0], state.You.Length, state.Board, possibleMoves)
	safeishMoves := []string{}
	for move, isSafe := range possibleMoves {
		if isSafe {
			safeishMoves = append(safeishMoves, move)
		}
	}

	if len(safeMoves) == 0 {
		nextMove = "down"
		log.Printf("%s MOVE %d: No safe moves detected! Moving %s\n", state.Game.ID, state.Turn, nextMove)
	} else {
		if len(safeishMoves) != 0 {
			nextMove = safeishMoves[rand.Intn(len(safeishMoves))]
		} else {
			nextMove = safeMoves[rand.Intn(len(safeMoves))]
		}

		//Battle Mode - eat all the food
		if len(state.Board.Snakes) > 1 {
			if len(state.Board.Food) > 0 {
				foodMove := moveToFood(state.You.Head, state.Board)
				if len(safeishMoves) != 0 {
					for i := 0; i < len(safeishMoves); i++ {
						if foodMove == safeishMoves[i] {
							nextMove = foodMove
						}
					}
				} else {
					for i := 0; i < len(safeMoves); i++ {
						if foodMove == safeMoves[i] {
							nextMove = foodMove
						}
					}
				}

			}
		}

		//Solo mode - only eat when necessary
		if int(state.You.Health) <= (state.Board.Height + state.Board.Width) {
			log.Printf("%s Low health %d\n", state.Game.ID, state.Turn)
			if len(state.Board.Food) > 0 {
				foodMove := moveToFood(state.You.Head, state.Board)
				if possibleMoves[foodMove] == true {
					nextMove = foodMove
				}
			}
		}
		log.Printf("%s MOVE %d: %s\n", state.Game.ID, state.Turn, nextMove)
	}
	return BattlesnakeMoveResponse{
		Move: nextMove,
	}
}
