package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Igor-Kreshchenko/go-calculator/info"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMetrics() (weight float64, height float64) {
	weight = getUserInput(info.WeightPrompt)
	height = getUserInput(info.HeightPrompt)

	return
}

func getUserInput(promptText string) float64 {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -2)
	enteredValue, _ := strconv.ParseFloat(userInput, 64)

	return enteredValue
}
