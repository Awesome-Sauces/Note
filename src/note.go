package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var file string

var app = tview.NewApplication()

var text = tview.NewTextView().
	SetDynamicColors(true).
	SetRegions(true).
	SetChangedFunc(func() {
		app.Draw()
	})

var position int = len(file)

var ResetColor string
var dir string

// BackGroundColorGlobal Text Customization variables

func overWrite() {
	d1 := []byte(file)
	err := os.WriteFile(os.Args[2], d1, 0644)
	check(err)
}

func main() {

	CheckArgs()

	var dat, err = os.ReadFile(os.Args[2])

	check(err)

	file = string(dat)

	switch BackGroundColorGlobal {
	case "[red]":
		text.SetBackgroundColor(tcell.ColorRed)
	case "[green]":
		text.SetBackgroundColor(tcell.ColorGreen)
	case "[yellow]":
		text.SetBackgroundColor(tcell.ColorYellow)
	case "[blue]":
		text.SetBackgroundColor(tcell.ColorBlue)
	case "[purple]":
		text.SetBackgroundColor(tcell.ColorPurple)
	case "[black]":
		text.SetBackgroundColor(tcell.NewRGBColor(30, 30, 30))
	case "[white]":
		text.SetBackgroundColor(tcell.ColorWhite)
	}

	text.SetTitleColor(tcell.ColorBlack)
	text.SetTitleAlign(tview.AlignCenter)

	text.SetBorder(true)
	text.SetTitle(TextColorGlobal + "[" + os.Args[2] + "]	 Click esc to quit!")
	text.SetBorderColor(tcell.ColorBlack)
	text.SetBorderAttributes(tcell.AttrDim)

	position = len(file)

	refresh(false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		if event.Key() == tcell.KeyLeft && event.Modifiers() == tcell.ModShift {
			setPos(len(regexp.MustCompile(`\S+\s*$`).ReplaceAllString(file[:position], "")))
			refresh(false)
			return event
		} else if event.Key() == tcell.KeyRight && event.Modifiers() == tcell.ModShift {
			setPos(len(file) - len(regexp.MustCompile(`^\s*\S+\s*`).ReplaceAllString(file[position:], "")))
			refresh(false)
			return event
		} else if event.Key() == tcell.KeyCtrlW {
			lastWord := regexp.MustCompile(`\S+\s*$`)
			newText := lastWord.ReplaceAllString(file[:position], "") + file[position:]
			position -= len(file) - len(newText)

			if position < 0 {
				position = 0
			}

			file = newText
			refresh(false)
			return event
		} else if event.Key() == tcell.KeyLeft {
			updatePos(false)
			refresh(false)
			return event
		} else if event.Key() == tcell.KeyRight {
			updatePos(true)
			refresh(false)
			return event
		} else if event.Key() == tcell.KeyUp {
			for i := position; i > 0; i-- {
				if file[i-1:i] == "\n" || file[i-1:i] == "\r\n" || file[i-1:i] == "\r" {
					setPos(i - 1)
					refresh(false)
					return event
				}
			}
			return event
		} else if event.Key() == tcell.KeyDown {
			for i := position; i > 0; i++ {
				if i+1 <= len(file) {
					if file[i:i+1] == "\n" || file[i:i+1] == "\r\n" || file[i:i+1] == "\r" {
						setPos(i + 1)
						refresh(false)
						return event
					}
				} else {
					return event
				}
			}
			return event
		} else if event.Key() == tcell.KeyEnter {
			file = file[0:position] + "\n" + file[position:]
			updatePos(true)
			refresh(false)
			return event
		} else if event.Key() == tcell.KeyDEL {

			if len(file) > 0 && position-1 != -1 {
				file = file[0:position-1] + file[position:]
				updatePos(false)
				refresh(false)

				return event
			}

			return event

		} else if event.Key() == tcell.KeyESC {
			overWrite()
			app.Stop()
		} else {

			if position == len(file) {
				file += string(event.Rune())
			} else {
				file = file[0:position] + string(event.Rune()) + file[position:]
			}

			updatePos(true)
			refresh(false)

			return event

		}

		return event
	})

	if err := app.SetRoot(text, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

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

	fmt.Println("Running Note v1.0.0:")
	fmt.Println("Using:\n- " + BackGround + " Background\n- " + Text + " Text Color")

	time.Sleep(500 * time.Millisecond)

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
