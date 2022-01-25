package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var CursorColorConfig map[int]string
var CursorTextConfig map[int]string

// TextColorBools struct
type TextColorBools struct {
	TextColor []TextColor `json:"TextColor"`
}

// TextColor struct
type TextColor struct {
	ColorRed    string `json:"ColorRed"`
	ColorGreen  string `json:"ColorGreen"`
	ColorYellow string `json:"ColorYellow"`
	ColorBlue   string `json:"ColorBlue"`
	ColorPurple string `json:"ColorPurple"`
	ColorCyan   string `json:"ColorCyan"`
	ColorWhite  string `json:"ColorWhite"`
}

// CursorColorBools struct which contains
// an array of users
type CursorColorBools struct {
	CursorColor []CursorColor `json:"CursorColor"`
}

// CursorColor struct which contains a name
// a type and a list of social links
type CursorColor struct {
	ColorRed    string `json:"ColorRed"`
	ColorGreen  string `json:"ColorGreen"`
	ColorYellow string `json:"ColorYellow"`
	ColorBlue   string `json:"ColorBlue"`
	ColorPurple string `json:"ColorPurple"`
	ColorCyan   string `json:"ColorCyan"`
	ColorWhite  string `json:"ColorWhite"`
}

func main() {
	CursorColorConfig = make(map[int]string)
	CursorTextConfig = make(map[int]string)
	loadColorConfig()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadColorConfig() {
	// Open our jsonFile
	jsonFile, err := os.Open("colorConf.json")
	check(err)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var TxtClr TextColorBools
	var CursorClr CursorColorBools

	// we unmarshal our byteArray
	err = json.Unmarshal(byteValue, &TxtClr)
	check(err)
	err = json.Unmarshal(byteValue, &CursorClr)
	check(err)

	for i := 0; i < len(TxtClr.TextColor); i++ {
		CursorTextConfig[0] = TxtClr.TextColor[i].ColorRed
		CursorTextConfig[1] = TxtClr.TextColor[i].ColorGreen
		CursorTextConfig[2] = TxtClr.TextColor[i].ColorYellow
		CursorTextConfig[3] = TxtClr.TextColor[i].ColorBlue
		CursorTextConfig[4] = TxtClr.TextColor[i].ColorPurple
		CursorTextConfig[5] = TxtClr.TextColor[i].ColorCyan
		CursorTextConfig[6] = TxtClr.TextColor[i].ColorWhite
	}
	for i := 0; i < len(CursorClr.CursorColor); i++ {
		CursorColorConfig[0] = CursorClr.CursorColor[i].ColorRed
		CursorColorConfig[1] = CursorClr.CursorColor[i].ColorGreen
		CursorColorConfig[2] = CursorClr.CursorColor[i].ColorYellow
		CursorColorConfig[3] = CursorClr.CursorColor[i].ColorBlue
		CursorColorConfig[4] = CursorClr.CursorColor[i].ColorPurple
		CursorColorConfig[5] = CursorClr.CursorColor[i].ColorCyan
		CursorColorConfig[6] = CursorClr.CursorColor[i].ColorWhite
	}
}
