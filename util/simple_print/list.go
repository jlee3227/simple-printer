package simple_print

import (
	"bufio"
	"fmt"
	"os"
)

func ReadListFromFile(path string) ([]string, error) {
	var list []string
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		list = append(list, line)
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func ReadListFromIO() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	fmt.Println("Enter list title:")
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	lines = append(lines, scanner.Text())

	fmt.Println("Input list items. Inputting an empty line will print the list.")
	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		lines = append(lines, "- "+line)
	}

	err := scanner.Err()
	if err != nil {
		return nil, err
	}

	return lines, nil
}
