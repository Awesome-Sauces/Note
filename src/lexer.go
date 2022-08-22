package main

import(
	"fmt"
	"strings"
)

type Lexer struct{
	text string
	position int
	char string
	tokens map[int]Token
	word string
	whitespace bool
	quotes bool
	variable bool
}

type mathlib struct{
	instance string
	digits map[int]string
}

var Math = mathlib{instance: "LEXER", digits: map[int]string{0: "0", 1: "1", 2: "2",
															 3: "3", 4: "4", 5: "5",
															 6: "6", 7: "7", 8: "8",
															 9:"9"}}


type Token struct{
	token string
	value string
}

func (math mathlib) isDigit(str string) bool {
	for LOOP := 0; LOOP <= len(math.digits); LOOP++ {
		if str == math.digits[LOOP] {return true}
	}

	return false
}

func (math mathlib) max(a int, b int) int {
	if b > a || b == a {return b} else {return a}
}

func (math mathlib) min(a int, b int) int {
	if b < a || b == a {return b} else {return a}
}

func (math mathlib) containsList(a string, b map[int]string) bool {
	for LOOP := 0; LOOP <= len(b); LOOP++ {
		if strings.Contains(a, b[LOOP]) {return true}
	}

	return false
}

func (math mathlib) isOperator(a string) bool{
	if(a == "%" ||
	   a == "=" ||
	   a == "/" ||
	   a == "(" ||
	   a == ")" ||
	   a == "*" || 
	   a == "+" ||
	   a == "-" ||
	   a == "!" ||
	   a == "&" ||
	   a == ";" ||
	   a == "{" ||
	   a == "}" ||
	   a == ":" ||
	   a == "," ||
	   a == "[" ||
	   a == "]" ||
	   a == "|"){return true}
	return false
}

func (lex *Lexer) lex(){

	fmt.Println("Lexer working")

	lex.position = 0

	lex.tokens = make(map[int]Token)

	lex.word = ""

	lex.nextChar()

	lex.whitespace = false

	lex.quotes = false

	lex.variable = false

	for lex.position <= len(lex.text){
		lex.iter()
	}
}

func (lex *Lexer) iter(){
	lex.nextChar()
	
	// Check for keywords in here

	if lex.char == "\"" && !lex.variable{
		if lex.quotes{
			lex.tokens[len(lex.tokens)] = Token{token: "string", value:lex.word}
			lex.newToken("\"")
			lex.quotes = false
		}else {lex.quotes = true}
	}else if lex.char == "$" || lex.variable{
			if lex.char == " " || Math.isOperator(lex.char){
				lex.word = lex.word[:Math.max(len(lex.word)-1, 0)]
				lex.tokens[len(lex.tokens)] = Token{token: "variable", value:rWhite(lex.word)}
				if Math.isOperator(lex.char){lex.tokens[len(lex.tokens)] = Token{token: "operator", value:lex.char}}
				lex.variable = false
				lex.newToken("variable")
			}else {lex.variable = true}
	}else if lex.char == "." && !lex.quotes || Math.isDigit(lex.char) && !lex.quotes {
		if !Math.isDigit(nextChar(lex.text, lex.position)) && nextChar(lex.text, lex.position) != "."{
			lex.tokens[len(lex.tokens)] = Token{token: "number", value:rWhite(rWhite(lex.word))}
			lex.newToken("number")
		}
	}else if rWhite(lex.word) == "func"{
		lex.tokens[len(lex.tokens)] = Token{token: "func", value:rWhite(lex.word)}
		lex.newToken("func")
	}else if rWhite(lex.word) == "loop"{
		lex.tokens[len(lex.tokens)] = Token{token: "loop", value:rWhite(lex.word)}
		lex.newToken("loop")
	}else if rWhite(lex.word) == "int"{
		lex.tokens[len(lex.tokens)] = Token{token: "type", value:rWhite(lex.word)}
		lex.newToken("type")
	}else if rWhite(lex.word) == "string"{
		lex.tokens[len(lex.tokens)] = Token{token: "type", value:rWhite(lex.word)}
		lex.newToken("type")
	}else if rWhite(lex.word) == "list"{
		lex.tokens[len(lex.tokens)] = Token{token: "type", value:rWhite(lex.word)}
		lex.newToken("type")
	}else if rWhite(lex.word) == "END_PROGRAM"{
		lex.tokens[len(lex.tokens)] = Token{token: "endprogram", value:rWhite(lex.word)}
		lex.newToken("type")
	}else if Math.isOperator(lex.char){
		lex.tokens[len(lex.tokens)] = Token{token: "operator", value:rWhite(lex.word)}
		lex.newToken("operator")
	}else if lex.char == " " || lex.whitespace{
		if !lex.quotes && !lex.variable{
			if lex.word == " " {
				lex.newToken("whitespace")
			}else{
				lex.tokens[len(lex.tokens)] = Token{token: "identifier", value:rWhite(lex.word)}
				lex.newToken("identifier")
			}	
		}
	}
}

func rWhite(e string) string {
	a := strings.ReplaceAll(e, " ", "")
	a = strings.ReplaceAll(a, "\n", "")
	a = strings.ReplaceAll(a, "\t", "")
	a = strings.ReplaceAll(a, "\"", "")

	return a
}

func (lex *Lexer) newToken(stype string){
	switch stype{
		case "\"":
			if(!lex.variable){lex.word=""}
		case "variable":
			if(!lex.variable){lex.word=""}
		case "func":
			if(!lex.variable){lex.word=""}
		case "loop":
			if(!lex.variable){lex.word=""}
		case "type":
			if(!lex.variable){lex.word=""}
		case "operator":
			if(!lex.variable){
				lex.word = lex.word[:Math.max(len(lex.word)-1, 0)]
				if len(rWhite(lex.word)) > 0 {

					lex.tokens[Math.max(len(lex.tokens)-1, 0)] = Token{token: "identifier", value:rWhite(lex.word)}
				
					lex.tokens[len(lex.tokens)] = Token{token: "operator", value:lex.char}
				}
				//if(lex.word != "" || lex.word != " "){lex.tokens[len(lex.tokens)] = Token{token: "identifier", value:lex.word}}
				//fmt.Println(lex.word + ":VALUE")
				lex.word = ""
			}
		case "number":
			if(!lex.variable){lex.word=""}
		case "whitespace":
			if(!lex.variable){lex.word=""}
		case "identifier":
			if(!lex.variable){lex.word=""}	
	}
}

func (lex *Lexer) nextChar(){
	lex.char = lex.text[Math.max(lex.position-1, 0):lex.position]
	lex.word += lex.char
	lex.position++
}

func nextChar(text string, position int) string{
	text = text[Math.max(position-1, 0):position]
	return  text
}
