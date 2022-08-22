package main

import (
    _ "fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/rivo/tview"
)

type TextEditor struct {
	rows       map[int]string
	rowSpacing int
	cursor     position
}

type position struct {
	x int
	y int
}

func (txt *TextEditor) getRow(pos position) string {
	return txt.rows[Math.min(pos.y, len(txt.rows))]
}

func (txt *TextEditor) getChar(pos position) string {
	return txt.rows[Math.min(pos.y, len(txt.rows))][Math.max(pos.x-1, 0):Math.min(pos.x, len(txt.rows[Math.min(pos.y, len(txt.rows))]))]
}

func (txt *TextEditor) showCursor(rows map[int]string) string {
	text := ""
	//vFile = tview.Escape(v.text[0:Math.max(v.position-1, 0)]) + color + tview.Escape(v.text[Math.max(v.position-1, 0):v.position]) + "[:-]" + tview.Escape(v.text[v.position:])
	for iter, i := range rows {
		if iter == txt.cursor.y {

			i = i[0:Math.max(txt.cursor.x-1, 0)] + "[:#00aeff]" + i[Math.max(txt.cursor.x-1, 0):Math.max(txt.cursor.x, 0)] + "[:-]" + i[Math.max(txt.cursor.x, 0):]

			//		i[Math.max(txt.cursor.x-1, 0):Math.min(txt.cursor.x, len(txt.rows[Math.min(txt.cursor.y, len(txt.rows))]))]
		}
		text += i + "\n"
	}

	return text
}

func (txt *TextEditor) getFormat() string {
	final := ""

	for LOOP := 1; LOOP <= len(txt.rows); LOOP++ {
		index := LOOP
		element := tview.Escape(txt.rows[LOOP])

		space := ""
		mDigits := iterativeDigitsCount(len(txt.rows))
		sDigits := iterativeDigitsCount(index)
		NumberColor := "[" + pColorTheme.lnColor + ":" + pColorTheme.lnbgColor + ":" + pColorTheme.lnstyleColor + "]"

		for _, et := range pColorTheme.keywords {

			element = strings.ReplaceAll(element, et.name, "["+et.color+"]"+et.name+"[#ffffff]")

			//			return i
			//ping <- true
		}

		//lineColor := "[" + pColorTheme.lnColor + ":" + pColorTheme.lnbgColor + ":" + pColorTheme.lnstyleColor + "]"

		for i := 1; i < (mDigits-sDigits)+1; i++ {
			space += " "
		}

		lineNumbers := space + strconv.Itoa(index) + " [-:-:-]"

		txt.rowSpacing = len(lineNumbers) - (7 + len(strconv.Itoa(index)))

		final += NumberColor + lineNumbers + element + "\n"

	}

	return final
}

func (txt *TextEditor) getLocation() position {
	return position{x: txt.cursor.x + txt.rowSpacing, y: txt.cursor.y}
}

func (txt *TextEditor) finalize() string {
	final := ""

	for LOOP := 1; LOOP <= len(txt.rows); LOOP++ {
		//	fmt.Println(LOOP)
		index := LOOP
		element := txt.rows[LOOP]

		space := ""
		mDigits := iterativeDigitsCount(len(txt.rows))
		sDigits := iterativeDigitsCount(index)
		NumberColor := "[" + pColorTheme.lnColor + ":" + pColorTheme.lnbgColor + ":" + pColorTheme.lnstyleColor + "]"

		if index == txt.cursor.y {
			element = tview.Escape(element[0:Math.min(Math.max(txt.cursor.x, 0), len(txt.rows[txt.cursor.y]))]) + "[:#00aeff] " + "[:-]" + tview.Escape(element[Math.min(Math.max(txt.cursor.x, 0), len(txt.rows[txt.cursor.y])):])
		}

		for _, et := range pColorTheme.keywords {

			element = strings.ReplaceAll(element, et.name, "["+et.color+"]"+et.name+"[#ffffff]")

		}

		for i := 1; i < (mDigits-sDigits)+1; i++ {
			space += " "
		}

		lineNumbers := space + strconv.Itoa(index) + " [-:-:-]"

		txt.rowSpacing = len(lineNumbers) - (7 + len(strconv.Itoa(index)))

		final += NumberColor + lineNumbers + element + "\n"

	}

	return final
}

func (txt *TextEditor) setLocation(pos position, spacing bool) {
	if spacing {
		txt.cursor = position{x: pos.x - txt.rowSpacing, y: pos.y}
	} else {
		txt.cursor = position{x: pos.x, y: pos.y}
	}

	txt.checkUP()
}

func (txt *TextEditor) getNewLocation() (int, int) {
	return txt.cursor.x, txt.cursor.y
}

func (txt *TextEditor) newLine() {
	txt.rows[txt.cursor.y] = txt.rows[txt.cursor.y][0:Math.max(txt.cursor.x, 0)] + "\n" + txt.rows[txt.cursor.y][Math.max(txt.cursor.x, 0):]
	temp := txt.getText()
	list := strings.Split(temp, "\n")
	txt.rows = make(map[int]string)


	for index, element := range list {
		if index+1 != len(list) {
			txt.rows[index+1] = element
		}
	}

	txt.moveDown()

}

func (txt *TextEditor) getText() string {
	text := ""

	for LOOP := 1; LOOP <= len(txt.rows); LOOP++ {
		i := txt.rows[LOOP]
		text += i + "\n"	
	}

	return text
}

func (txt *TextEditor) loadText(text string) {
	
	yMap := make(map[int]string)

	word := strings.Split(text, "\n")

	for LOOP := 1; LOOP <= len(word); LOOP++ {
		element := word[LOOP-1]
		yMap[len(yMap)+1] = element
	}
	

	txt.rows = yMap

	txt.cursor = position{x: 0, y: 0}



}

func (txt *TextEditor) initCursor() {
	txt.cursor = position{x: 0, y: 0}
}

func (txt *TextEditor) deleteWord() {
	lastWord := regexp.MustCompile(`\S+\s*$`)
	newText := lastWord.ReplaceAllString(txt.rows[txt.cursor.y][:txt.cursor.x], "") + txt.rows[txt.cursor.y][txt.cursor.x:]
	txt.cursor.x = Math.max(txt.cursor.x-(len(txt.rows[txt.cursor.y])-len(newText)), 0)

	if txt.cursor.x < 0 {
		txt.cursor.x = 0
	}

	txt.rows[txt.cursor.y] = newText
}

func (txt *TextEditor) deleteChar() {
	if txt.cursor.x > 0 && len(txt.rows[txt.cursor.y]) > 0 {
		txt.rows[txt.cursor.y] = txt.rows[txt.cursor.y][0:Math.min(Math.max(txt.cursor.x-1, 0), len(txt.rows[txt.cursor.y]))] + txt.rows[txt.cursor.y][txt.cursor.x:]
		txt.moveLeft()
	} else if txt.cursor.x <= 0 {
		if len(txt.rows[txt.cursor.y]) > 0 && txt.cursor.y != 1 {
			row := txt.rows[txt.cursor.y]
			txt.rows = removeValue(txt.rows, txt.cursor.y)
			txt.rows[Math.max(txt.cursor.y-1, 0)] += row	
		}else if txt.cursor.y != 1{
			txt.rows = removeValue(txt.rows, txt.cursor.y)
		}

		txt.moveUp() 
	}
}

func (txt *TextEditor) addChar(e string) {
	txt.rows[txt.cursor.y] = txt.rows[txt.cursor.y][0:txt.cursor.x] + e + txt.rows[txt.cursor.y][txt.cursor.x:]
	txt.cursor.x++

	txt.checkUP()
}

func (txt *TextEditor) moveUp() {
	txt.cursor.y = Math.min(Math.max(0, txt.cursor.y-1), len(txt.rows))
	if txt.cursor.x > len(txt.rows[txt.cursor.y]) {
		txt.cursor.x = len(txt.rows[txt.cursor.y])
	}
	if txt.cursor.x < len(txt.rows[txt.cursor.y]) {
		txt.cursor.x = len(txt.rows[txt.cursor.y])
	}

	txt.checkUP()
}

func (txt *TextEditor) moveDown() {
	txt.cursor.y = Math.min(len(txt.rows), txt.cursor.y+1)
	if txt.cursor.x > len(txt.rows[txt.cursor.y]) {
		txt.cursor.x = len(txt.rows[txt.cursor.y])
	}
	if txt.cursor.x < len(txt.rows[txt.cursor.y]) {
		txt.cursor.x = len(txt.rows[txt.cursor.y])
	}

	txt.checkUP()
}

func (txt *TextEditor) moveRight() {
	if txt.cursor.x+1 > len(txt.rows[txt.cursor.y]) && txt.cursor.y != len(txt.rows) {
		txt.cursor.y = Math.min(txt.cursor.y+1, len(txt.rows))
		txt.cursor.x = 0
	} else {
		txt.cursor.x = Math.min(txt.cursor.x+1, len(txt.rows[txt.cursor.y]))
	}

	txt.checkUP()
}

func (txt *TextEditor) moveLeft() {
	if txt.cursor.x-1 < 0 && txt.cursor.y != 1 {
		txt.cursor.y = Math.max(txt.cursor.y-1, 1)
		txt.cursor.x = len(txt.rows[txt.cursor.y])
	} else {
		txt.cursor.x = Math.max(txt.cursor.x-1, 0)
	}

	txt.checkUP()
}

func (txt *TextEditor) checkUP(){
	if txt.cursor.x > len(txt.rows[txt.cursor.y]) {
		txt.cursor.x=len(txt.rows[txt.cursor.y])
	}else if txt.cursor.x <= 0{
		txt.cursor.x = 0
	}

	if txt.cursor.y > len(txt.rows){
		txt.cursor.y = len(txt.rows)
	}else if txt.cursor.y <= 0 {
		txt.cursor.y = 1
	}
}

func (txt *TextEditor) moveWordLeft() {
	txt.cursor.x = len(regexp.MustCompile(`\S+\s*$`).ReplaceAllString(txt.rows[txt.cursor.y][:txt.cursor.x], ""))
}

func (txt *TextEditor) moveWordRight() {
	txt.cursor.x = len(txt.rows[txt.cursor.y]) - len(regexp.MustCompile(`^\s*\S+\s*`).ReplaceAllString(txt.rows[txt.cursor.y][txt.cursor.x:], ""))
}

func (txt *TextEditor) saveFile() string {

	text := ""

	for LOOP := 1; LOOP <= len(txt.rows); LOOP++ {
		element := txt.rows[LOOP]

		if LOOP == len(txt.rows) {
			text += element
		} else {text += element+"\n"}

	}

	return text
}

func removeValue(mt map[int]string, iter int) map[int]string {
	dv := false
	rmap := make(map[int]string)

	for LOOP := 1; LOOP <= len(mt); LOOP++ {
		index := LOOP
		element := mt[LOOP]

		if index == iter {dv = true} else if dv {
			rmap[LOOP-1] = element
		}else{
			rmap[LOOP] = element
		}
		
	}

	return rmap
}
