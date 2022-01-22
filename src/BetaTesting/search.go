package main

import (
	"fmt"
	"strings"
)

var NewlineChars map[int]NLC
var NLCChecks map[int]bool
var position int

type rightStr struct {
	// Where The String is located
	x    int
	y    int
	data string
}

type NLC struct {
	X int `default:"0"`
	Y int `default:"0"`
	Z int `default:"0"`
}

func main() {
	process()
}

func mkNLC(x int, y int, z int) NLC {
	val := NLC{X: x, Y: y, Z: z}
	return val
}

func process() {
	NewlineChars = make(map[int]NLC)
	localNLC := make(map[int]int)
	mapLoop := 0
	str := "123456789\n12345678\n12345"
	position = len(str)
	for iter := position; iter >= 1; iter-- {
		mapLoop++
		if mapLoop >= 3 {
			if strings.Contains(str[iter-1:iter], "\n") == true {
				localNLC[len(localNLC)] = iter
			}
		}

	}
	if localNLC[2] == 0 {
		localNLC[len(localNLC)] = 0
	}
	fmt.Println(localNLC)
	nlc := mkNLC(localNLC[0], localNLC[1], localNLC[2])
	NewlineChars[0] = nlc
}
