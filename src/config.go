package main

import (
	"os"
)

// TextColorBools struct
type TextColorBools struct {
	TextColor []TextColor `json:"TextColor"`
}

// BackGroundColorBools struct
type BackGroundColorBools struct {
	BackGroundColor []BackGroundColor `json:"BackGroundColor"`
}

// TextColor struct for colorConfig.json, Purely cosmetics
type TextColor struct {
	ColorRed    string `json:"ColorRed"`
	ColorGreen  string `json:"ColorGreen"`
	ColorYellow string `json:"ColorYellow"`
	ColorBlue   string `json:"ColorBlue"`
	ColorPurple string `json:"ColorPurple"`
	ColorCyan   string `json:"ColorCyan"`
	ColorWhite  string `json:"ColorWhite"`
}

// BackGroundColor struct for colorConfig.json, Purely cosmetics
type BackGroundColor struct {
	ColorRed    string `json:"ColorRed"`
	ColorGreen  string `json:"ColorGreen"`
	ColorYellow string `json:"ColorYellow"`
	ColorBlue   string `json:"ColorBlue"`
	ColorPurple string `json:"ColorPurple"`
	ColorCyan   string `json:"ColorCyan"`
	ColorWhite  string `json:"ColorWhite"`
}

// Background and TextColor 
// related variables
var BackGroundColorGlobal string
var TextColorGlobal string
var BackGroundColorConfig map[int]string
var BackGroundTextConfig map[int]string
var ColorMap map[int]string

func init() {

	// Setting up Maps for Text & BackGround color Customization
	BackGroundColorConfig = make(map[int]string)
	BackGroundTextConfig = make(map[int]string)
	ColorMap = make(map[int]string)

	// Adding the standard colors to ColorMap
	ColorMap[0] = "[red]"
	ColorMap[1] = "[green]"
	ColorMap[2] = "[yellow]"
	ColorMap[3] = "[blue]"
	ColorMap[4] = "[purple]"
	ColorMap[5] = "[black]"
	ColorMap[6] = "[white]"
	ResetColor = "[white]"

	// Find location of home dir then check if colorConfig.json exists
	dirname, err := os.UserHomeDir()
	dir = dirname + "/note/colorConfig.json"

	check(err)

	// Loading color value config
	LoadColorConfig()

	// Asigning values
	BackGroundColorGlobal, TextColorGlobal = FindColorConfig()
}
