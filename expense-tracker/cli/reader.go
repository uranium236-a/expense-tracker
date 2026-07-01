package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadInput(prompt string) (string, error) {
	fmt.Print(prompt)

	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)

	return input, nil
}

func ReadInt(prompt string) (int, error) {
	fmt.Print(prompt)

	input, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)

	integer, err := strconv.Atoi(input)

	if err != nil {
		return 0, err
	}

	return integer, nil
}

func ConvertInt(input string) (int, error) {
	input = strings.TrimSpace(input)

	integer, err := strconv.Atoi(input)

	if err != nil {
		return 0, err
	}

	return integer, nil
}
