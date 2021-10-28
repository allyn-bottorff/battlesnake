package main

import (
	"testing"
)

func equalMoveMaps(m1 map[string]bool, m2 map[string]bool) bool {
	if m1["up"] != m2["up"] || m1["down"] != m2["down"] || m1["left"] != m2["left"] || m1["right"] != m2["right"] {
		return false
	}
	return true
}

func TestBattlesnakeInfoResponse(t *testing.T) {
	testResponse := BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "MazerRackham",
		Color:      "0467d1",
		Head:       "default",
		Tail:       "default",
	}

	if testResponse != info() {
		t.Fatalf("info() return did not match test data")
	}
}

func TestCheckForBodyUp(t *testing.T) {
	var testBody = []Coord{
		{
			X: 5,
			Y: 5,
		},
		{
			X: 5,
			Y: 6,
		},
		{
			X: 5,
			Y: 7,
		},
	}

	possibleMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}

	expectedMoves := map[string]bool{
		"up":    false,
		"down":  true,
		"left":  true,
		"right": true,
	}

	checkForBody(testBody, possibleMoves)
	if equalMoveMaps(possibleMoves, expectedMoves) != true {
		t.Fatalf("Possible moves: %v not equal to expected moves. Head at %v", possibleMoves, testBody[0])

	}

}

func TestCheckForBodyDown(t *testing.T) {
	var testBody = []Coord{
		{
			X: 5,
			Y: 5,
		},
		{
			X: 5,
			Y: 4,
		},
		{
			X: 5,
			Y: 3,
		},
	}

	possibleMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}

	expectedMoves := map[string]bool{
		"up":    true,
		"down":  false,
		"left":  true,
		"right": true,
	}

	checkForBody(testBody, possibleMoves)
	if equalMoveMaps(possibleMoves, expectedMoves) != true {
		t.Fatalf("Possible moves: %v not equal to expected moves. Head at %v", possibleMoves, testBody[0])

	}

}

func TestCheckForBodyLeft(t *testing.T) {
	var testBody = []Coord{
		{
			X: 5,
			Y: 5,
		},
		{
			X: 4,
			Y: 5,
		},
		{
			X: 3,
			Y: 5,
		},
	}

	possibleMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}

	expectedMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  false,
		"right": true,
	}

	checkForBody(testBody, possibleMoves)
	if equalMoveMaps(possibleMoves, expectedMoves) != true {
		t.Fatalf("Possible moves: %v not equal to expected moves. Head at %v", possibleMoves, testBody[0])

	}

}

func TestCheckForBodyRight(t *testing.T) {
	var testBody = []Coord{
		{
			X: 5,
			Y: 5,
		},
		{
			X: 6,
			Y: 5,
		},
		{
			X: 7,
			Y: 5,
		},
	}

	possibleMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}

	expectedMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": false,
	}

	checkForBody(testBody, possibleMoves)
	if equalMoveMaps(possibleMoves, expectedMoves) != true {
		t.Fatalf("Possible moves: %v not equal to expected moves.", possibleMoves)

	}

}

func TestCheckForBodyTail(t *testing.T) {
	var testBody = []Coord{
		{
			X: 5,
			Y: 5,
		},
		{
			X: 6,
			Y: 5,
		},
		{
			X: 7,
			Y: 5,
		},
		{
			X: 4,
			Y: 5,
		},
	}

	possibleMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}

	expectedMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  false,
		"right": false,
	}

	checkForBody(testBody, possibleMoves)
	if equalMoveMaps(possibleMoves, expectedMoves) != true {
		t.Fatalf("Possible moves: %v not equal to expected moves.", possibleMoves)

	}

}

func TestCheckForWallsUp(t *testing.T) {
	testBoard := Board{
		Height: 11,
		Width:  11,
	}
	possibleMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}
	head := Coord{
		X: 5,
		Y: 10,
	}

	expectedMoves := map[string]bool{
		"up":    false,
		"down":  true,
		"left":  true,
		"right": true,
	}

	checkForWalls(head, testBoard, possibleMoves)

	if equalMoveMaps(possibleMoves, expectedMoves) != true {
		t.Fatalf("Possible moves: %v not equal to expected moves.", possibleMoves)
	}

}

func TestCheckForWallsDown(t *testing.T) {
	testBoard := Board{
		Height: 11,
		Width:  11,
	}
	possibleMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}
	head := Coord{
		X: 5,
		Y: 0,
	}

	expectedMoves := map[string]bool{
		"up":    true,
		"down":  false,
		"left":  true,
		"right": true,
	}

	checkForWalls(head, testBoard, possibleMoves)

	if equalMoveMaps(possibleMoves, expectedMoves) != true {
		t.Fatalf("Possible moves: %v not equal to expected moves.", possibleMoves)
	}

}

func TestCheckForWallsLeft(t *testing.T) {
	testBoard := Board{
		Height: 11,
		Width:  11,
	}
	possibleMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}
	head := Coord{
		X: 0,
		Y: 5,
	}

	expectedMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  false,
		"right": true,
	}

	checkForWalls(head, testBoard, possibleMoves)

	if equalMoveMaps(possibleMoves, expectedMoves) != true {
		t.Fatalf("Possible moves: %v not equal to expected moves.", possibleMoves)
	}

}

func TestCheckForWallsRight(t *testing.T) {
	testBoard := Board{
		Height: 11,
		Width:  11,
	}
	possibleMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}
	head := Coord{
		X: 10,
		Y: 5,
	}

	expectedMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": false,
	}

	checkForWalls(head, testBoard, possibleMoves)

	if equalMoveMaps(possibleMoves, expectedMoves) != true {
		t.Fatalf("Possible moves: %v not equal to expected moves.", possibleMoves)
	}

}

func TestCoordInSnake(t *testing.T) {
	body := []Coord{
		{
			X: 1,
			Y: 1,
		},
		{
			X: 2,
			Y: 1,
		},
	}
	testIn := Coord{
		X: 1,
		Y: 1,
	}
	testOut := Coord{
		X: 2,
		Y: 3,
	}

	if coordInSnake(testIn, body) != true {
		t.Fatalf("Coordinate in snake not caught")

	}

	if coordInSnake(testOut, body) != false {
		t.Fatalf("Coodinate not in snake caught")
	}

}

func TestCheckForSnakes(t *testing.T) {
	testP := Coord{
		X: 2,
		Y: 2,
	}

}
