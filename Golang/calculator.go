package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)
func Prompt(label string) string{
    reader := bufio.NewReader(os.Stdin)
	fmt.Fprint(os.Stderr, label+" ")
	userInput, _ := reader.ReadString('\n')
	return strings.TrimSpace(userInput)
}
func DoTheMath(f string, o string, s string) (string, error) {
	fNum, err := strconv.ParseFloat(f, 64)
	if err != nil {
		return "", err
	}
	sNum, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return "", err
	}
	var output float64
	switch o {
	case "*":
		output = fNum * sNum
	case "/":
		output = fNum / sNum
	case "+":
		output = fNum + sNum
	case "-":
		output = fNum - sNum
	default:
		return "Invalid operation", nil
	}
	return fmt.Sprintf("%f", output), nil
}
func main(){
	first_number := Prompt("Enter first number:")
	operation := Prompt("Enter operation:(+, -, *, /):")
	second_number := Prompt("Enter second number:")
	output, err := DoTheMath(first_number, operation, second_number)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Result: %s\n", output)
}