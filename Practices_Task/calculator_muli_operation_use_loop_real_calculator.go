package main

import (
	"errors"
	"fmt"
	"strings"
	//"strconv"
	// "strings"
	"unicode"
)

func calculatorMultipleOperationsLoop(){
	for{
		var expression string
		fmt.Print("Enter Exprssion : ")
		fmt.Scan(&expression)
		if(expression == "exit" || expression == "Exit" || expression == "EXIT"){
			fmt.Print("Close Calculator!")
			break
		}
	}
}




func tokenize(expr string)([]string, error){
	var tokens []string
	tempStoreDigit := ""


	for _,ch := range expr{
		if(unicode.IsDigit(ch) || ch == '.'){
			tempStoreDigit += string(ch)
		}else if(strings.ContainsRune("+-*/()",ch)){
			if(tempStoreDigit != ""){
				tokens = append(tokens, tempStoreDigit)
				tempStoreDigit = ""
			}
			tokens = append(tokens, string(ch))
		}else{
			return nil, errors.New("invalid expression!")
		}
	}

	if(tempStoreDigit != ""){
		tokens = append(tokens, tempStoreDigit)
	}

	return tokens, nil
}

