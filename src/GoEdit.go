package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var clear map[string]func() //create a map for storing clear funcs

func init() {
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
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func main() {
	colorReset := "\033[0m"

	argCheck()

	// Start Writing to File
	switch os.Args[1] {
	case "-a":
		dat, err := os.ReadFile(os.Args[2])
		check(err)
		input := startup(string(dat))
		overWrite(input)
		fmt.Println(colorReset)
		exec.Command("stty", "-T", "/dev/tty", "-echo").Run()
		os.Exit(3)
	}
}

func overWrite(input string) {
	d1 := []byte(input)
	err := os.WriteFile(os.Args[2], d1, 0644)
	check(err)
}
func argCheck() {
	if len(os.Args) > 3 {
		os.Exit(1)
	} else if len(os.Args) < 3 {
		os.Exit(1)
	} else {
		fmt.Println("\033[31m", "Checks Succeeded!\n Proceed with caution!")
	}
}

func helpQ() {
	colorBlue := "\033[34m"
	colorReset := "\033[0m"
	for help := 0; help <= 10; help++ {
		if help == 10 {
			fmt.Println(colorBlue, "Press ESC to exit and save!")
			fmt.Println(colorReset, "")
		} else {
			fmt.Println(colorReset, "")
		}
	}
}

func startup(file string) string {
	colorWhite := "\033[37m"
	m := make(map[int]string)
	var str string
	str = file
	fprint := 0
	for fileITER := 1; fileITER <= len(str); fileITER++ {
		m[len(m)] = str[fileITER-1 : fileITER]
	}
	ch := make(chan string)
	go func(ch chan string) {
		// disable input buffering
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
		// do not display entered characters on the screen
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
		var b = make([]byte, 1)
		for {
			os.Stdin.Read(b)
			ch <- string(b)
		}
	}(ch)
	for {
		if fprint <= 1 {
			CallClear()
			fmt.Println(colorWhite, str)
			helpQ()
			fprint++
		}
		stdin, _ := <-ch
		fmt.Println(stdin)
		time.Sleep(1000 * time.Millisecond)
		by := []uint8(stdin)
		if by[0] == 127 {
			time.Sleep(1 * time.Millisecond)
		} else if stdin == "`" {
			time.Sleep(1 * time.Millisecond)
		} else if by[0] == 27 {
			return str
		} else if by[0] == 65 {
			fmt.Println("Hello World you pressed up arrow")
			time.Sleep(100 * time.Millisecond)
		} else {
			m[len(m)] = stdin
			str += stdin
			CallClear()
			fmt.Println(colorWhite, str)
			go helpQ()
		}

		if by[0] == 127 {
			str = strings.TrimSuffix(str, m[len(m)-1])
			delete(m, len(m)-1)
			CallClear()
			fmt.Println(colorWhite, str)
			go helpQ()

		}

	}

}
