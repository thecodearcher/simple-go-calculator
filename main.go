package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var answer int = 0

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nSimple Calculator")
	fmt.Println("Give space between each input e.g \"2 + 4 - 1\"")
	fmt.Println("-----------------------------------------------")

	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	handler(strings.Split(text, " "))
	fmt.Println("-----------------------------------------------")
	fmt.Println()
	fmt.Println(text, " = ", answer)
	fmt.Println()

}

func handler(values []string) {
	numb := 0
	for index, val := range values {
		prevNumb := 0
		if index == 0 {
			prevNumb, _ = strconv.Atoi(values[index])
		} else {
			prevNumb, _ = strconv.Atoi(values[index-1])
		}

		if index < len(values)-1 {
			numb, _ = strconv.Atoi(values[index+1])
		} else {
			numb, _ = strconv.Atoi(values[index])
		}

		if index > 0 {
			cal(prevNumb, numb, val)
		}

	}

}

func cal(prevNumb int, nextNumb int, operator string) {
	if answer != 0 {
		prevNumb = answer
	}
	switch operator {
	case "+":
		answer = prevNumb + nextNumb
	case "-":
		answer = prevNumb - nextNumb
	case "*":
		answer = prevNumb * nextNumb
	case "/":
		answer = prevNumb / nextNumb
	}
}
