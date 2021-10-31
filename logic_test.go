package main

import (
	"testing"
)

func equalMoveMaps(m1 map[string]bool, m2 map[string]bool) bool {
	var matches bool
	if m1["up"] != m2["up"] || m1["down"] != m2["down"] || m1["left"] != m2["left"] || m1["right"] != m2["right"] {
		matches = false
	} else {
		matches = true
	}

	return matches
}

func TestBattlesnakeInfoResponse(t *testing.T) {
	testResponse := BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "MazerRackham",
		Color:      "#0467d1",
		Head:       "shades",
		Tail:       "pixel",
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
	up := Coord{
		X: 5,
		Y: 6,
	}

	checkForSelf(testBody, up)
	if checkForSelf(testBody, up) != true {
		t.Fatalf("Failed to find self up.")

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
	down := Coord{
		X: 5,
		Y: 4,
	}

	if checkForSelf(testBody, down) != true {
		t.Fatalf("Failed to find self, down")

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
	left := Coord{
		X: 4,
		Y: 5,
	}

	if checkForSelf(testBody, left) != true {
		t.Fatalf("Failed to find self, left.")
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
	right := Coord{
		X: 6,
		Y: 5,
	}

	if checkForSelf(testBody, right) != true {
		t.Fatalf("Failed to find self, right.")

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
	snake1 := []Coord{
		{
			X: 1,
			Y: 1,
		},
		{
			X: 1,
			Y: 2,
		},
		{
			X: 1,
			Y: 3,
		},
	}
	snake2 := []Coord{
		{
			X: 2,
			Y: 1,
		},
		{
			X: 2,
			Y: 2,
		},
		{
			X: 2,
			Y: 3,
		},
	}

	battlesnakes := []Battlesnake{
		{
			Body: snake1,
		},
		{
			Body: snake2,
		},
	}
	board := Board{
		Snakes: battlesnakes,
	}

	testP1 := Coord{
		X: 2,
		Y: 2,
	}
	testP2 := Coord{
		X: 3,
		Y: 3,
	}

	if checkForSnakes(board, testP1) != true {
		t.Fatalf("Failed to find point which is occupied by other snakes.")
	}

	if checkForSnakes(board, testP2) != false {
		t.Fatalf("Incorrectly found a point which is unoccupied by other snakes.")
	}

}
