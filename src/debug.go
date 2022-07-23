package main

import(
	"fmt"
	"strings"
	"time"
	"unicode"
	"os"
)

var colorReset = "\033[0m"

type debug struct {
	alias string
	color string
}

func (db *debug) sleep(e int){
	for i := 0; i <= e; i++{
		fmt.Println(db.color + db.alias + colorReset + " Sleeping...")
		time.Sleep(1 * time.Second)
	}
}

func (db *debug) NewError(err error, text string, a ...string){
	if(err != nil){
		for i := range a {text = strings.Replace(text, "%s", a[i], 1)}
		db.colorize()
		fmt.Println(db.color + db.alias + colorReset + text)
		os.Exit(0)
	}
}

func (db *debug) crash(){
	os.Exit(1)
}

func (db *debug) isSpace(e string) bool{
	for i := range e{
		if unicode.IsSpace(rune(e[i])){return true}
	}
	return false
}

func (db *debug) out(text string, a ...string){
	for i := range a {text = strings.Replace(text, "%s", a[i], 1)}
	db.colorize()
	fmt.Println(db.color + db.alias + colorReset + text)
}

func (db *debug) outM(text string, a ...*string){
	for i := range a {text = strings.Replace(text, "%s", *a[i], 1)}
	db.colorize()
	fmt.Printf("%s%s%s%v\n", db.color, db.alias, colorReset, &text)
}

func (db *debug) outf(text string, a ...string){
	for i := range a {text = strings.Replace(text, "%s", a[i], 1)}
	db.colorize()
	fmt.Printf(db.color + db.alias + colorReset + text)
}

func (db *debug) colorize(){
	switch db.color{
		case "yellow":
			db.color = "\033[33m"
		case "red":
			db.color = "\033[31m"
		case "green":
			db.color = "\033[32m"
		case "blue":
			db.color = "\033[34m"
		case "purple":
			db.color = "\033[35m"
		case "cyan":
			db.color = "\033[36m"
		case "white":
			db.color = "\033[37m"
	}
}

