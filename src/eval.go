package main

import(
	"os"
	"strings"
	"time"
    "strconv"
    "github.com/rivo/tview"
)

type variable struct{
	name string
	vtype string
	value string
}

type function struct{
	name string
	kwargs map[int]string
}

type list struct{
	name string
	args map[int]string
}

type loop struct{
	args map[int]Token
}

func (lex *Lexer) eval(){

	//Debug := debug{alias: "EVAL: ", color: "red"}

//	var transpiled string = ""
	var variables = make(map[string]variable)
	var lists = make(map[string]list)
	var looping bool = false

	var comment bool = false

	var tempLines = make(map[int]Token)

	for i := 0; i <= len(lex.tokens); i++{
		if comment ||
		   lex.tokens[i].value == "#" {
			if lex.tokens[i].value == ";" {comment = !comment}
		}else if len(lex.tokens[i].token) > 0 && len(lex.tokens[i].value) > 0{
			addToken(i, lex, tempLines)
		}
	} 

	for i := 0; i <= len(lex.tokens); i++ {
		if _, ok := lex.tokens[i];!ok {continue}

		if lex.tokens[i].token == "loop" &&
			lex.tokens[Math.min(i+1, len(lex.tokens))].value == "(" &&
			lex.tokens[Math.min(i+3, len(lex.tokens))].value == ")" &&
			lex.tokens[Math.min(i+4, len(lex.tokens))].value == "{" || looping {
				loopNum := 0
				endofLoop := i+4
				var err error
				eol := make(map[int]Token)

				tDebug := debug{alias: "ERROR: ", color: "red"}

				if lex.tokens[Math.min(i+2, len(lex.tokens))].token == "variable"{
					if _, ok := lists[lex.tokens[Math.min(i+2, len(lex.tokens))].value]; ok {
				        loopNum = len(lists[lex.tokens[Math.min(i+2, len(lex.tokens))].value].args)
					}else{
						loopNum, err =  strconv.Atoi(variables[lex.tokens[Math.min(i+2, len(lex.tokens))].value].value)
						loopNum++
						tDebug.NewError(err, "Cannot instantiate list index with string: %s ", lex.tokens[Math.min(i+2, len(lex.tokens))].value)
					}
				}else if lex.tokens[Math.min(i+2, len(lex.tokens))].token == "number"{
					loopNum, err = strconv.Atoi(lex.tokens[Math.min(i+2, len(lex.tokens))].value)
					loopNum++
					tDebug.NewError(err, "Cannot instantiate list index with string: %s ", lex.tokens[Math.min(i+2, len(lex.tokens))].value)
				}

				for j := endofLoop; j <= len(lex.tokens); j++ {
					if lex.tokens[j].value == "}"{
						endofLoop = j
						j = len(lex.tokens) + 1
						continue
					}else{
						eol[len(eol)] = lex.tokens[j]
					}
				}

				LOOP := loop{args: eol}
				VARIABLE := variable{name: "$loop", vtype: "loop", value: "1"}

				for i := 1; i <= loopNum-1; i++{
					VARIABLE.value = strconv.Itoa(i)
					variables[VARIABLE.name] = VARIABLE
					LOOP.run(variables, lists)
				}

				i = endofLoop
				
		}else if lex.tokens[i].token == "type" &&
			 lex.tokens[Math.min(i+1, len(lex.tokens))].token == "variable" &&
			 lex.tokens[Math.min(i+2, len(lex.tokens))].value == "="{
				if lex.tokens[i].value == "list"{
					if lex.tokens[Math.min(i+3, len(lex.tokens))].value  == "["{
						callEnd := i+1
			 			for f := callEnd; f <= len(lex.tokens); f++{
			 				if lex.tokens[Math.min(f, len(lex.tokens))].value == "]"{
			 					callEnd = f
			 					f = len(lex.tokens) + 1
			 				}
			 			}

			 		tlist := make(map[int]string)

			 		for j := i+1; j <= callEnd; j++{
			 			if lex.tokens[Math.min(j, len(lex.tokens))].value != "[" && 
			 				lex.tokens[Math.min(j, len(lex.tokens))].value != "]" &&
			 				lex.tokens[Math.min(j, len(lex.tokens))].value != "," &&
			 				lex.tokens[Math.min(j, len(lex.tokens))].value != "=" {
			 				if lex.tokens[Math.min(j, len(lex.tokens))].token == "variable"{
			 					tlist[len(tlist)] = variables[lex.tokens[Math.min(j, len(lex.tokens))].value].value
			 				}else{
			 		//			fmt.Println("tlist: " + lines[i][Math.min(j, len(lines[i]))].value, len(tlist))
			 					tlist[len(tlist)] = lex.tokens[Math.min(j, len(lex.tokens))].value
			 				}
			 		//	Debug.out("VALUE: %s TYPE: %s", lines[i][Math.min(j, len(lines[i]))].value, lines[i][Math.min(j, len(lines[i]))].token)	
			 			}
			 		}


					listme :=  list{name: lex.tokens[Math.min(i+1, len(lex.tokens))].value, args: tlist}

					lists[listme.name] = listme

					//for i := 1; i <= len(listme.args)-1; i++{
					//	Debug.out("%s : %s", listme.name, listme.args[i])
					//}

			 			
					}
					
			 	}
			 	

				
			
				tempVar := variable{name: lex.tokens[Math.min(i+1, len(lex.tokens))].value, vtype: lex.tokens[i].value, value: lex.tokens[Math.min(i+3, len(lex.tokens))].value}
				variables[tempVar.name] = tempVar
				//Debug.out("VARIABLE %s TYPE %s VALUE %s", tempVar.name, tempVar.vtype, tempVar.value)
			}else if lex.tokens[i].token == "identifier" &&
			 	lex.tokens[Math.min(i+1, len(lex.tokens))].value == "("{
			 	callEnd := i+1
			 	for f := callEnd; f <= len(lex.tokens); f++{
			 		if lex.tokens[Math.min(f, len(lex.tokens))].value == ")"{
			 			callEnd = f
			 			f = len(lex.tokens) + 1
			 		}
			 	}

			 	tlist := make(map[int]string)

			 	tDebug := debug{alias: "ERROR: ", color: "red"}


			 	for j := i+1; j <= callEnd; j++{
			 		if lex.tokens[Math.min(j, len(lex.tokens))].value != "(" && 
			 		lex.tokens[Math.min(j, len(lex.tokens))].value != ")" &&
			 		lex.tokens[Math.min(j, len(lex.tokens))].value != ","{
			 			if lex.tokens[Math.min(j, len(lex.tokens))].token == "variable"{
							if lex.tokens[Math.min(j+1, len(lex.tokens))].value == "[" &&
							   lex.tokens[Math.min(j+3, len(lex.tokens))].value == "]" &&
							   lex.tokens[Math.min(j+2, len(lex.tokens))].token == "variable"  {

								slist := lists[lex.tokens[Math.min(j, len(lex.tokens))].value].args
							   
							   	intVar, err := strconv.Atoi(variables[lex.tokens[Math.min(j+2, len(lex.tokens))].value].value)
					
								tDebug.NewError(err, "Cannot call list index with string: %s ", variables[lex.tokens[Math.min(j+2, len(lex.tokens))].value].value)

								value := slist[intVar]

								tlist[len(tlist)] = value
								
							   }else if lex.tokens[Math.min(j+1, len(lex.tokens))].value == "[" &&
							   		 lex.tokens[Math.min(j+3, len(lex.tokens))].value == "]"{

									slist := lists[lex.tokens[Math.min(j, len(lex.tokens))].value].args
							   
							   		intVar, err := strconv.Atoi(lex.tokens[Math.min(j+2, len(lex.tokens))].value)
					
									tDebug.NewError(err, "Cannot Call list index with string: %s ", lex.tokens[Math.min(j+2, len(lex.tokens))].value)

									tlist[len(tlist)] = slist[intVar]
							   }else{
							   		if lex.tokens[Math.min(j, len(lex.tokens))].value != "[" &&
							   			lex.tokens[Math.min(j+2, len(lex.tokens))].value != "]" &&
							   			lex.tokens[Math.min(j-1, len(lex.tokens))].value != "[" &&
							   			lex.tokens[Math.min(j+1, len(lex.tokens))].value != "]"{
							   				tlist[len(tlist)] = variables[lex.tokens[Math.min(j, len(lex.tokens))].value].value	
							   			}
							   }		
			 			}else{
			 				if lex.tokens[Math.min(j, len(lex.tokens))].token != "operator" &&
			 					lex.tokens[Math.min(j-1, len(lex.tokens))].value != "[" &&
			 					lex.tokens[Math.min(j+1, len(lex.tokens))].value != "]" {
			 					tlist[len(tlist)] = lex.tokens[Math.min(j, len(lex.tokens))].value
			 				}
			 			}
			 		}
			 	}


				fstdlib := function{name: lex.tokens[i].value, kwargs: tlist}

				fstdlib.stdlib()
			}
	}	

	//saveFile(transpiled)
	
}

func (lp *loop) run(variables map[string]variable, lists map[string]list){
	for i := 0; i <= len(lp.args); i++ {
		if _, ok := lp.args[i];!ok {continue}
		if lp.args[i].token == "identifier" &&
			 	lp.args[Math.min(i+1, len(lp.args))].value == "("{
			 	callEnd := i+1
			 	for f := callEnd; f <= len(lp.args); f++{
			 		if lp.args[Math.min(f, len(lp.args))].value == ")"{
			 			callEnd = f
			 			f = len(lp.args) + 1
			 		}
			 	}

			 	tlist := make(map[int]string)

			 	tDebug := debug{alias: "ERROR: ", color: "red"}


			 	for j := i+1; j <= callEnd; j++{
			 		if lp.args[Math.min(j, len(lp.args))].value != "(" && 
			 		lp.args[Math.min(j, len(lp.args))].value != ")" &&
			 		lp.args[Math.min(j, len(lp.args))].value != ","{
			 			if lp.args[Math.min(j, len(lp.args))].token == "variable"{
							if lp.args[Math.min(j+1, len(lp.args))].value == "[" &&
							   lp.args[Math.min(j+3, len(lp.args))].value == "]" &&
							   lp.args[Math.min(j+2, len(lp.args))].token == "variable"  {

								slist := lists[lp.args[Math.min(j, len(lp.args))].value].args
							   
							   	intVar, err := strconv.Atoi(variables[lp.args[Math.min(j+2, len(lp.args))].value].value)
					
								tDebug.NewError(err, "Cannot call list index with string: %s ", variables[lp.args[Math.min(j+2, len(lp.args))].value].value)

								value := slist[intVar]

								tlist[len(tlist)] = value
								
							   }else if lp.args[Math.min(j+1, len(lp.args))].value == "[" &&
							   		 lp.args[Math.min(j+3, len(lp.args))].value == "]"{

									slist := lists[lp.args[Math.min(j, len(lp.args))].value].args
							   
							   		intVar, err := strconv.Atoi(lp.args[Math.min(j+2, len(lp.args))].value)
					
									tDebug.NewError(err, "Cannot Call list index with string: %s ", lp.args[Math.min(j+2, len(lp.args))].value)

									tlist[len(tlist)] = slist[intVar]
							   }else{
							   		if lp.args[Math.min(j, len(lp.args))].value != "[" &&
							   			lp.args[Math.min(j+2, len(lp.args))].value != "]" &&
							   			lp.args[Math.min(j-1, len(lp.args))].value != "[" &&
							   			lp.args[Math.min(j+1, len(lp.args))].value != "]"{
							   				tlist[len(tlist)] = variables[lp.args[Math.min(j, len(lp.args))].value].value	
							   			}
							   }		
			 			}else{
			 				if lp.args[Math.min(j, len(lp.args))].token != "operator" &&
			 					lp.args[Math.min(j-1, len(lp.args))].value != "[" &&
			 					lp.args[Math.min(j+1, len(lp.args))].value != "]"{
			 					tlist[len(tlist)] = lp.args[Math.min(j, len(lp.args))].value
			 				}
			 			}
			 		}
			 	}


				fstdlib := function{name: lp.args[i].value, kwargs: tlist}

				fstdlib.stdlib()
			}
	}
}

func (funct *function) stdlib(){
	Debug := debug{alias: "Stdlib ", color: "blue"}
	

	switch funct.name{
	case "Escape":
		Debug.out("%s", tview.Escape("Hello"))
	case "LineColors":
		pColorTheme.setlnColor(strings.ReplaceAll(funct.kwargs[0], "\"", ""))
		pColorTheme.setlnbgColor(strings.ReplaceAll(funct.kwargs[2], "\"", ""))
		pColorTheme.setStyleColor(strings.ReplaceAll(funct.kwargs[3], "\"", ""))
		
	case "Background":
		pColorTheme.changeBGcolor(strings.ReplaceAll(funct.kwargs[0], "\"", ""))
	case "NewKeyword":
		Debug.out("%s:%s:%s", funct.kwargs[0], funct.kwargs[1], funct.kwargs[2])
		pColorTheme.NewKeyword(cKeyWord{
									extension: strings.ReplaceAll(funct.kwargs[0], "\"", ""),
									name: strings.ReplaceAll(funct.kwargs[1], "\"", ""),
									color: strings.ReplaceAll(funct.kwargs[2], "\"", "")})			
	case "printf":

		output := ""

		for i := range funct.kwargs {output += funct.kwargs[i]}
		Debug.out(strings.ReplaceAll(output, "\"", ""))
	case "sleep":
		strVar := funct.kwargs[0]
		intVar, err := strconv.Atoi(strVar)
		check(err)
		time.Sleep(time.Duration(intVar) * time.Second)
	case "memory":
		    // convert to uintptr
    //	var adr uint64
    //	adr, err := strconv.ParseUint(&funct.kwargs[0], 0, 64)
   	//	check(err)
    //	var ptr uintptr = uintptr(adr)

    	str := funct.kwargs[0]

    	Debug.outM("%s", &str)
				
	}
}

/*
func ptrToString(ptr uintptr) string {
    p := unsafe.Pointer(ptr)
    return *(*string)(p)
}
*/

func addToken(i int, lex *Lexer, tempLines map[int]Token){
	if lex.tokens[i].token != "" && lex.tokens[i].value != "" ||
	   lex.tokens[i].token != "\t" && lex.tokens[i].value != "\t" ||
	   lex.tokens[i].token != " " && lex.tokens[i].value != " " ||
	   lex.tokens[i].token != "\n" && lex.tokens[i].value != "\n"   {
		tempLines[len(tempLines)] = lex.tokens[i]	
	}
}
 
func saveFile(text string){
	d1 := []byte(text)
	myfile, err := os.Create("ROCKY-OUTPUT.go")
	check(err)
	myfile.Close()
	err = os.WriteFile("ROCKY-OUTPUT.go", d1, 0644)
	check(err)
}
