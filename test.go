package main

import (
	"fmt"

	"github.com/Dharshan2208/code-compiler/internal/sandbox"
)

func main() {
	sb := sandbox.Sandbox{}

	result := sb.Run(
		"compiler-python",
		"/home/chifuyu/Coding/Self/Projects/code-compiler/test",
		[]string{
			"python3",
			"main.py",
		},
	)

	fmt.Println("STDOUT:")
	fmt.Println(result.Stdout)

	fmt.Println("STDERR:")
	fmt.Println(result.Stderr)

	fmt.Println("ERROR:")
	fmt.Println(result.Error)
}
