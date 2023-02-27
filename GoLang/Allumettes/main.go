package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	if input == 7 {
		fmt.Println("sept")
	} else {
		fmt.Println("choisis un autre nombre fdp")
	}
	return
}
