package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringInput() string {
	reader := bufio.NewReader(os.Stdin)
	inp, _ := reader.ReadString('\n')

	return strings.TrimSpace(inp)
}

func getIntInput() int {
	input := stringInput()
	num, err := strconv.Atoi(input)

	for err != nil {
		fmt.Println(err)
		fmt.Println("Please try again.")
		fmt.Print("Insert value: ")
		input = stringInput()
		num, err = strconv.Atoi(input)
	}

	return num
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
