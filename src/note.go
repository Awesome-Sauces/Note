package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Simple error checking
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Initializing tview.Application and
// tview.Flex
var app = tview.NewApplication()
var flex = tview.NewFlex()

var pColorTheme *ColorTheme = &ColorTheme{name: "", keywords: make(map[string]cKeyWord)}

// NoteFile struct map
// Stores all essential data
var noteFiles = make(map[int]NoteFile)

// File window open in
// order of arguments
var current_file = 1

var ResetColor string
var dir string


// Note version for printing to terminal
var NOTE_VERSION string = "Note v1.3.2"
var NOTE_HELP string = NOTE_VERSION + "\n\n" +
		"Usage: note [-arg]\n" +
		"    note [--help]\n" +
		"    note [filename]\n" +
		"    note [-txt] [color]\n" +
		"    note [-bg] [color]\n" +
		"Experimental:\n" +
		"    note -script [SCRIPT_NAME]\n" +
		"Color list:\n" +
		"    [white]\n    [red]\n    [black]\n    [blue]\n    [green]\n    [purple]\n    [yellow]\n" 


func main() {
	// Checks args given to cli, if none are 
	// given then give a help message
	CheckArgs()

	// Start up Message
	// Sleep .5 seconds to give it
	// a "loading" effect
	fmt.Println("Running " + NOTE_VERSION +":")
	fmt.Println("Using:\n- " + BackGroundColorGlobal + " Background\n- " + TextColorGlobal + " Text Color")
	time.Sleep(500 * time.Millisecond)

	// Loop through noteFiles map and
	// Initializing all the tview.TextView
	for i, s := range noteFiles {
		s.autoCreateTextView()
		flex.AddItem(noteFiles[i].textView, 0, 1, false)
	}


	// Basic User input setup
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		if !noteFiles[current_file].textView.HasFocus(){app.SetFocus(noteFiles[current_file].textView)}
		
		// When user presses Right Arrow key + ctrl
		// Switch over to the file on right
		if event.Key() == tcell.KeyRight && event.Modifiers() == tcell.ModCtrl {
			if current_file + 1 > len(noteFiles){return event}
			noteFiles[current_file].textView.SetText(noteFiles[current_file].getStringFormat())
			current_file++
			app.SetFocus(noteFiles[current_file].textView)
			noteFiles[current_file].autoCreateTextView()
			return event
		// When user presses Left Arrow key + ctrl
		// Switch over to the file on left
		}else if event.Key() == tcell.KeyLeft && event.Modifiers() == tcell.ModCtrl {
			if current_file - 1 < 1 {return event}
			noteFiles[current_file].textView.SetText(noteFiles[current_file].getStringFormat())
			current_file--
			app.SetFocus(noteFiles[current_file].textView)
			noteFiles[current_file].autoCreateTextView()
			return event
		} 

		return event
	})


	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

/*

	Ignore this huge spaghetti code used for .json

*/


func LoadColorConfig() {
	// Open our jsonFile
	jsonFile, err := os.Open(dir)
	check(err)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var TxtClr TextColorBools
	var BackGroundClr BackGroundColorBools

	// we unmarshal our byteArray
	err = json.Unmarshal(byteValue, &TxtClr)
	check(err)
	err = json.Unmarshal(byteValue, &BackGroundClr)
	check(err)

	for i := 0; i < len(TxtClr.TextColor); i++ {
		BackGroundTextConfig[0] = TxtClr.TextColor[i].ColorRed
		BackGroundTextConfig[1] = TxtClr.TextColor[i].ColorGreen
		BackGroundTextConfig[2] = TxtClr.TextColor[i].ColorYellow
		BackGroundTextConfig[3] = TxtClr.TextColor[i].ColorBlue
		BackGroundTextConfig[4] = TxtClr.TextColor[i].ColorPurple
		BackGroundTextConfig[5] = TxtClr.TextColor[i].ColorCyan
		BackGroundTextConfig[6] = TxtClr.TextColor[i].ColorWhite
	}
	for i := 0; i < len(BackGroundClr.BackGroundColor); i++ {
		BackGroundColorConfig[0] = BackGroundClr.BackGroundColor[i].ColorRed
		BackGroundColorConfig[1] = BackGroundClr.BackGroundColor[i].ColorGreen
		BackGroundColorConfig[2] = BackGroundClr.BackGroundColor[i].ColorYellow
		BackGroundColorConfig[3] = BackGroundClr.BackGroundColor[i].ColorBlue
		BackGroundColorConfig[4] = BackGroundClr.BackGroundColor[i].ColorPurple
		BackGroundColorConfig[5] = BackGroundClr.BackGroundColor[i].ColorCyan
		BackGroundColorConfig[6] = BackGroundClr.BackGroundColor[i].ColorWhite
	}
}

func FindColorConfig() (string, string) {
	var BackGround string
	var Text string
	for i := 0; i <= 6; i++ {
		if BackGroundColorConfig[i] == "true" {
			//fmt.Println(i)
			BackGround = ColorMap[i]
		}
		if BackGroundTextConfig[i] == "true" {
			//fmt.Println(i)
			Text = ColorMap[i]
		}

	}

	return BackGround, Text
}

func textColorChange(Arg1 string) {
	switch Arg1 {
	case "red":
		LoadColorConfig()
		ChangeColor("true", "false", "false", "false", "false", "false", "false", BackGroundColorConfig[0], BackGroundColorConfig[1], BackGroundColorConfig[2], BackGroundColorConfig[3], BackGroundColorConfig[4], BackGroundColorConfig[5], BackGroundColorConfig[6])
		break
	case "green":
		LoadColorConfig()
		ChangeColor("false", "true", "false", "false", "false", "false", "false", BackGroundColorConfig[0], BackGroundColorConfig[1], BackGroundColorConfig[2], BackGroundColorConfig[3], BackGroundColorConfig[4], BackGroundColorConfig[5], BackGroundColorConfig[6])
		break
	case "yellow":
		LoadColorConfig()
		ChangeColor("false", "false", "true", "false", "false", "false", "false", BackGroundColorConfig[0], BackGroundColorConfig[1], BackGroundColorConfig[2], BackGroundColorConfig[3], BackGroundColorConfig[4], BackGroundColorConfig[5], BackGroundColorConfig[6])
		break
	case "blue":
		LoadColorConfig()
		ChangeColor("false", "false", "false", "true", "false", "false", "false", BackGroundColorConfig[0], BackGroundColorConfig[1], BackGroundColorConfig[2], BackGroundColorConfig[3], BackGroundColorConfig[4], BackGroundColorConfig[5], BackGroundColorConfig[6])
		break
	case "purple":
		LoadColorConfig()
		ChangeColor("false", "false", "false", "false", "true", "false", "false", BackGroundColorConfig[0], BackGroundColorConfig[1], BackGroundColorConfig[2], BackGroundColorConfig[3], BackGroundColorConfig[4], BackGroundColorConfig[5], BackGroundColorConfig[6])
		break
	case "black":
		LoadColorConfig()
		ChangeColor("false", "false", "false", "false", "false", "true", "false", BackGroundColorConfig[0], BackGroundColorConfig[1], BackGroundColorConfig[2], BackGroundColorConfig[3], BackGroundColorConfig[4], BackGroundColorConfig[5], BackGroundColorConfig[6])
		break
	case "white":
		LoadColorConfig()
		ChangeColor("false", "false", "false", "false", "false", "false", "true", BackGroundColorConfig[0], BackGroundColorConfig[1], BackGroundColorConfig[2], BackGroundColorConfig[3], BackGroundColorConfig[4], BackGroundColorConfig[5], BackGroundColorConfig[6])
		break
	default:
		fmt.Println("No matching color was found")
		fmt.Println("Use note --help for help")
		os.Exit(1)
	}

}

func backGroundChange(Arg1 string) {
	switch Arg1 {
	case "red":
		LoadColorConfig()
		ChangeColor(BackGroundTextConfig[0], BackGroundTextConfig[1], BackGroundTextConfig[2], BackGroundTextConfig[3], BackGroundTextConfig[4], BackGroundTextConfig[5], BackGroundTextConfig[6], "true", "false", "false", "false", "false", "false", "false")
		break
	case "green":
		LoadColorConfig()
		ChangeColor(BackGroundTextConfig[0], BackGroundTextConfig[1], BackGroundTextConfig[2], BackGroundTextConfig[3], BackGroundTextConfig[4], BackGroundTextConfig[5], BackGroundTextConfig[6], "false", "true", "false", "false", "false", "false", "false")
		break
	case "yellow":
		LoadColorConfig()
		ChangeColor(BackGroundTextConfig[0], BackGroundTextConfig[1], BackGroundTextConfig[2], BackGroundTextConfig[3], BackGroundTextConfig[4], BackGroundTextConfig[5], BackGroundTextConfig[6], "false", "false", "true", "false", "false", "false", "false")
		break
	case "blue":
		LoadColorConfig()
		ChangeColor(BackGroundTextConfig[0], BackGroundTextConfig[1], BackGroundTextConfig[2], BackGroundTextConfig[3], BackGroundTextConfig[4], BackGroundTextConfig[5], BackGroundTextConfig[6], "false", "false", "false", "true", "false", "false", "false")
		break
	case "purple":
		LoadColorConfig()
		ChangeColor(BackGroundTextConfig[0], BackGroundTextConfig[1], BackGroundTextConfig[2], BackGroundTextConfig[3], BackGroundTextConfig[4], BackGroundTextConfig[5], BackGroundTextConfig[6], "false", "false", "false", "false", "true", "false", "false")
		break
	case "black":
		LoadColorConfig()
		ChangeColor(BackGroundTextConfig[0], BackGroundTextConfig[1], BackGroundTextConfig[2], BackGroundTextConfig[3], BackGroundTextConfig[4], BackGroundTextConfig[5], BackGroundTextConfig[6], "false", "false", "false", "false", "false", "true", "false")
		break
	case "white":
		LoadColorConfig()
		ChangeColor(BackGroundTextConfig[0], BackGroundTextConfig[1], BackGroundTextConfig[2], BackGroundTextConfig[3], BackGroundTextConfig[4], BackGroundTextConfig[5], BackGroundTextConfig[6], "false", "false", "false", "false", "false", "false", "true")
		break
	default:
		fmt.Println("No matching color was found")
		fmt.Println("Use note --help for help")
		os.Exit(1)
	}

}

func ChangeColor(valueA string, valueB string, valueC string, valueD string, valueE string, valueF string, valueG string, valueH string, valueI string, valueJ string, valueK string, valueL string, valueM string, valueN string) {
	str := `{
  "TextColor": [
    {
      "ColorRed":"` + valueA + `",
      "ColorGreen":"` + valueB + `",
      "ColorYellow":"` + valueC + `",
      "ColorBlue":"` + valueD + `",
      "ColorPurple":"` + valueE + `",
      "ColorCyan" :"` + valueF + `",
      "ColorWhite" :"` + valueG + `"
    }
  ],
  "BackGroundColor": [
    {
      "ColorRed":"` + valueH + `",
      "ColorGreen":"` + valueI + `",
      "ColorYellow":"` + valueJ + `",
      "ColorBlue":"` + valueK + `",
      "ColorPurple":"` + valueL + `",
      "ColorCyan" :"` + valueM + `",
      "ColorWhite" :"` + valueN + `"
    }
  ]
}`

	d1 := []byte(str)

	err := os.WriteFile(dir, d1, 0644)
	if err != nil {
		return
	}
}
