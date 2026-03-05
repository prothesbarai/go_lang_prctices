package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func calculatorMultipleOperationsLoop() {
	for {
		var expression string
		fmt.Print("Enter Expression : ")
		fmt.Scan(&expression)
		if expression == "exit" || expression == "Exit" || expression == "EXIT" {
			fmt.Print("Calculator Close !")
			break
		}

		result , err := evaluate(expression)

		if(err != nil){
			fmt.Print("Error : ",err)
			continue
		}

		fmt.Print("\n",result,"\n")
	}
}



func evaluate(expr string)(float64,error){
	tokens , err := tokenized(expr)
	if(err != nil){
		return 0, err
	}
	postFixTokens := infixToPostfix(tokens)
	result := evaluatePostfixCalculation(postFixTokens)

	return result, nil
}


/// >>> Tokenized from given expression
func tokenized(expr string)([]string,error){
	var tokens []string
	tempStoreNum := ""

	for _, ch := range expr{
		if(unicode.IsDigit(ch) || ch == '.'){
			tempStoreNum += string(ch)
		}else if(strings.ContainsRune("+-*/()",ch)){
			if(tempStoreNum != ""){
				tokens = append(tokens, tempStoreNum)
				tempStoreNum = ""
			}
			tokens = append(tokens, string(ch))
		}else{
			return nil, errors.New("Invalid Character Detected!")
		}
	}

	if(tempStoreNum != ""){
		tokens = append(tokens, tempStoreNum)
	}

	return tokens,nil
}



// >>> Operator Priority
func precedence(op string) int{
	switch op{
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return 0
}


// >>> Expression to Convert PostFix  ["10","+","5","*","2"] → ["10","5","2","*","+"]
func infixToPostfix(tokens []string) []string{
	var stack []string
	var output []string

	for _, token := range tokens{

		if(unicode.IsDigit(rune(token[0]))){
			output = append(output, token)
		}else if(token == "("){
			stack = append(stack, token)
		}else if(token == ")"){
			for( len(stack) > 0 && stack[len(stack)-1] != "(" ) {
				output = append(output, stack[len(stack)-1]) // last Top element output e jasse  ex :  [*,(,+] 
				stack = stack[:len(stack) - 1] // Output e jaoya last top element delete korce stack theke  ex :  [*,(]
			}
			stack = stack[:len(stack)-1] //  ( remove cause ata output e jay na , now ex : [*]
		}else{
			for ( len(stack) > 0  && precedence(stack[len(stack)-1]) >= precedence(token) ){
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		}

	}

	for len(stack) > 0{
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return output
}


func evaluatePostfixCalculation(postFixTokens []string) float64{
	var stack []float64

	for _, tokens := range postFixTokens{
		if(unicode.IsDigit(rune(tokens[0]))){
			var num float64
			fmt.Sscanf(tokens, "%f" , &num)
			stack = append(stack, num)
		}else{
			var result float64

			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch tokens {
				case "+": 
					result = a + b 
				case "-": 
					result = a - b 
				case "*": 
					result = a * b 
				case "/": 
					result = a / b
			}

			stack =append(stack, result)
		}
	}

	return stack[0]
}