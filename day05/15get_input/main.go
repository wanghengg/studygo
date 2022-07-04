package main

import (
	"bufio"
	"fmt"
	"os"
)

func useBufIO() {

	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	fmt.Printf("你输入的内容是: %s", s)
}

func main() {
	useBufIO()
}
