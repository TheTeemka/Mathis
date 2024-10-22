package main

import (
	"fmt"
	"os"

	"github.com/TheTeemka/SiMath/logic"
)

func main() {
	str := os.Args[1]
	parser := logic.NewParser(str)
	root := parser.Parse()
	fmt.Println("Answer:", root.Solve())
}
