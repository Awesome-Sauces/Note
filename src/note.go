package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"unicode/utf8"
)

var clear map[string]func() //create a map for storing clear funcs
var rmvStr string
var position int
var lstr string
var rstr string
var str string

// COLORS map of colors for terminal
var COLORS map[string]string

func init() {
	COLORS = make(map[string]string)
	COLORS["colorReset"] = "\033[0m"
	COLORS["colorRed"] = "\033[31m"
	COLORS["colorGreen"] = "\033[32m"
	COLORS["colorYellow"] = "\033[33m"
	COLORS["colorBlue"] = "\033[34m"
	COLORS["colorPurple"] = "\033[35m"
	COLORS["colorCyan"] = "\033[36m"
	COLORS["colorWhite"] = "\033[37m"
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
	clear["mac"] = func() {
        cmd := exec.Command("clear") //Windows example, its tested 
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func main() {
	CheckArgs()
	// Creating ASCII key
	CallClear()
	ASCII := Symbols()
	dat, err := os.ReadFile(os.Args[1])
	str = string(dat)
	position = len(str)
	check(err)
	// Setting Cursor position to the end of file text
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// Hide Text
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	var b []byte = make([]byte, 3)
	for {
		b[1] = 0
		b[2] = 0
		LiveUpdate("NULL", "Update")
		os.Stdin.Read(b)
		switch ASCII[int(b[0])] {
		case "SPACE":
			LiveUpdate(" ", "AddChar")
		case "ESC":
			X := ASCII[int(b[0])]
			Y := ASCII[int(b[1])]
			Z := ASCII[int(b[2])]
			if X == "ESC" {
				if Y == "[" {
					if Z == "A" {
						ArrowUp()
						LiveUpdate("NULL", "Update")
						//go MousePointerLocationCalc("UP_ARROW", str)
					} else if Z == "B" {
						ArrowDown()
						LiveUpdate("NULL", "Update")
						//go MousePointerLocationCalc("DOWN_ARROW", str)
					} else if Z == "D" {
						position--
						if position == -1 {
							position++
						}
						process()
						LiveUpdate("NULL", "Update")

						//go MousePointerLocationCalc("LEFT_ARROW", str)
					} else {
						position++
						if position == len(str)+1 {
							position--
						}
						process()
						LiveUpdate("NULL", "Update")
						//go MousePointerLocationCalc("RIGHT_ARROW", str)
					}
				} else {
					process()
					// Un-hide Text
					exec.Command("stty", "-F", "/dev/tty", "echo").Run()
					// Save changes to file
					overWrite()
					CallClear()
					os.Exit(3)
				}

			}

		case "TAB":
			LiveUpdate("	", "AddChar")
		case "LF":
			LiveUpdate("\n", "AddChar")
		case "DEL":
			process()
			LiveUpdate("NULL", "DelChar")
		default:
			LiveUpdate(ASCII[int(b[0])], "AddChar")
		}

	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func LiveUpdate(still string, UpdateType string) int {
	// I have no idea how to make live text without it being buggy/using 3rd party libraries
	switch UpdateType {
	// This case adds a character to the file!
	case "AddChar":
		position++
		str = lstr + still + rstr
		process()
		CallClear()
		fmt.Fprintf(os.Stderr, "\r%s%s%s<|>%s%s", COLORS["colorWhite"], lstr, COLORS["colorPurple"], COLORS["colorWhite"], rstr)
	// This case deletes a character from the file. (This was way harder than supposed too)
	case "DelChar":
		process()
		str = lstr + trimFirstRune(rstr)
		position--
		process()
		CallClear()
		if position == -1 {
			position++
		}
		LiveUpdate("NULL", "Update")
		fmt.Fprintf(os.Stderr, "\r%s%s%s<|>%s%s", COLORS["colorWhite"], lstr, COLORS["colorPurple"], COLORS["colorWhite"], rstr)
	// This case just updates the text in terminal
	case "Update":
		process()
		CallClear()
		fmt.Fprintf(os.Stderr, "\r%s%s%s<|>%s%s", COLORS["colorWhite"], lstr, COLORS["colorPurple"], COLORS["colorWhite"], rstr)
		return 0
	}
	return 0
}

func overWrite() {
	d1 := []byte(str)
	err := os.WriteFile(os.Args[1], d1, 0644)
	check(err)
}

func process() int {
	if position > len(str) {
		return 0
	} else if position < 0 {
		return 0
	} else {
		lstr = str[0:position]
		rstr = str[position:len(str)]
		fmt.Println(len(rstr))
		if len(rstr) == 0 {
			return 0
		} else {
			rmvStr = rstr[0:1]
		}
		return 0
	}

}

func CheckArgs() {
	if len(os.Args) < 1 {
		os.Exit(3)
	}
}

func ArrowUp() int {
	// Making map for storing where all the NewLine characters are located.
	NewLinePosition := make(map[int]int)
	// Finding all NewLine Characters start from position
	for loopy := 0; loopy <= 0; loopy++ {
		for iter := position; iter >= 1; iter-- {
			if strings.Contains(str[iter-1:iter], "\n") == true {
				NewLinePosition[len(NewLinePosition)] = iter
			}
		}
		// Setting the New position
		if NewLinePosition[0]-1 > len(str) {
			return 0
		} else if NewLinePosition[0]-1 < 0 {
			return 0
		} else {
			position = NewLinePosition[0] - 1
			for k := range NewLinePosition {
				delete(NewLinePosition, k)
			}
		}

	}
	return 0
}

func ArrowDown() int {
	for location := position; location <= len(str); location++ {
		if strings.Contains(str[location-1:location], "\n") == true {
			if location+1 > len(str) {
				return 0
			} else {
				position = location + 1
			}
			return 0
		}
	}
	return 0
}

func Symbols() map[int]string {
	symbols := make(map[int]string)
	symbols[0] = "NUL"
	symbols[1] = "SOH"
	symbols[2] = "STX"
	symbols[3] = "ETX"
	symbols[4] = "EOT"
	symbols[5] = "ENQ"
	symbols[6] = "ACK"
	symbols[7] = "BEL"
	symbols[8] = "BS"
	symbols[9] = "TAB"
	symbols[10] = "LF"
	symbols[11] = "VT"
	symbols[12] = "FF"
	symbols[13] = "CR"
	symbols[14] = "SO"
	symbols[15] = "SI"
	symbols[16] = "DLE"
	symbols[17] = "DC1"
	symbols[18] = "DC2"
	symbols[19] = "DC3"
	symbols[20] = "DC4"
	symbols[21] = "NAK"
	symbols[22] = "SYN"
	symbols[23] = "ETB"
	symbols[24] = "CAN"
	symbols[25] = "EM"
	symbols[26] = "SUB"
	symbols[27] = "ESC"
	symbols[28] = "FS"
	symbols[29] = "GS"
	symbols[30] = "RS"
	symbols[31] = "US"
	symbols[32] = "SPACE"
	symbols[33] = "!"
	symbols[34] = "\""
	symbols[35] = "#"
	symbols[36] = "$"
	symbols[37] = "%"
	symbols[38] = "&"
	symbols[39] = "'"
	symbols[40] = "("
	symbols[41] = ")"
	symbols[42] = "*"
	symbols[43] = "+"
	symbols[44] = ","
	symbols[45] = "-"
	symbols[46] = "."
	symbols[47] = "/"
	symbols[48] = "0"
	symbols[49] = "1"
	symbols[50] = "2"
	symbols[51] = "3"
	symbols[52] = "4"
	symbols[53] = "5"
	symbols[54] = "6"
	symbols[55] = "7"
	symbols[56] = "8"
	symbols[57] = "9"
	symbols[58] = ":"
	symbols[59] = ";"
	symbols[60] = "<"
	symbols[61] = "="
	symbols[62] = ">"
	symbols[63] = "?"
	symbols[64] = "@"
	symbols[65] = "A"
	symbols[66] = "B"
	symbols[67] = "C"
	symbols[68] = "D"
	symbols[69] = "E"
	symbols[70] = "F"
	symbols[71] = "G"
	symbols[72] = "H"
	symbols[73] = "I"
	symbols[74] = "J"
	symbols[75] = "K"
	symbols[76] = "L"
	symbols[77] = "M"
	symbols[78] = "N"
	symbols[79] = "O"
	symbols[80] = "P"
	symbols[81] = "Q"
	symbols[82] = "R"
	symbols[83] = "S"
	symbols[84] = "T"
	symbols[85] = "U"
	symbols[86] = "V"
	symbols[87] = "W"
	symbols[88] = "X"
	symbols[89] = "Y"
	symbols[90] = "Z"
	symbols[91] = "["
	symbols[92] = "\\"
	symbols[93] = "]"
	symbols[94] = "^"
	symbols[95] = "_"
	symbols[96] = "`"
	symbols[97] = "a"
	symbols[98] = "b"
	symbols[99] = "c"
	symbols[100] = "d"
	symbols[101] = "e"
	symbols[102] = "f"
	symbols[103] = "g"
	symbols[104] = "h"
	symbols[105] = "i"
	symbols[106] = "j"
	symbols[107] = "k"
	symbols[108] = "l"
	symbols[109] = "m"
	symbols[110] = "n"
	symbols[111] = "o"
	symbols[112] = "p"
	symbols[113] = "q"
	symbols[114] = "r"
	symbols[115] = "s"
	symbols[116] = "t"
	symbols[117] = "u"
	symbols[118] = "v"
	symbols[119] = "w"
	symbols[120] = "x"
	symbols[121] = "y"
	symbols[122] = "z"
	symbols[123] = "{"
	symbols[124] = "|"
	symbols[125] = "}"
	symbols[126] = "~"
	symbols[127] = "DEL"

	return symbols
}
