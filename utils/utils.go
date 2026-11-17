package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"uuid"
)

var reader = bufio.NewReader(os.Stdin) // 1 reader dùng chung toàn chương trình

func GenerateId() string {
	return uuid.New().String()
}

func ReadInput(promt string) string {
	fmt.Print(promt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func GetPositiveInt(promt string) int {
	for {
		input := ReadInput(promt)
		value, err := strconv.Atoi(input)
		if err == nil && value > 0 {
			return value
		}

		fmt.Println("XXX Gia tri khong hop le ")
	}
}

func GetPositiveFloat(promt string) float64 {
	for {
		input := ReadInput(promt)
		value, err := strconv.ParseFloat(input, 64)
		if err == nil && value > 0 {
			return value
		}

		fmt.Println("XXX Gia tri khong hop le ")
	}
}

func GetOptionalFloat(promt string, currentValue float64) float64 {
	fmt.Printf("%s (Hien tai: %.2f): ", promt, currentValue)
	input := ReadInput(promt)
	if input == "" {
		return currentValue
	}

	value, err := strconv.ParseFloat(input, 64)
	if err == nil && value < 0 {
		return currentValue
	}

	return value

}

func GetNoneEmptyString(promt string) string {
	for {
		input := ReadInput(promt)
		if input != "" {
			return input
		}

		fmt.Println("Hay nhap it nhat 1 ki tu ")
	}
}

func GetOptionalString(prompt string, currentValue string) string {
	fmt.Printf("%s (Hien tai: %s): ", prompt, currentValue)
	input := ReadInput(prompt)
	if input == "" {
		return currentValue
	}
	return input

}

func ClearScreen() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("XXX error clearing screen: ", err)
	}
}

type HasId interface {
	GetId() int
}

func IsDuplicateId[T HasId](id int, list []T) bool {
	for _, item := range list {
		if item.GetId() == id {
			return false
		}
	}

	return true
}




