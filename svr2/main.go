package main

import (
	"fmt"

	"github.com/ethanvc/testgolanglint/svr2/util"
)

func main() {
	m := make(map[string]string)
	for k, _ := range m {
		fmt.Println(k)
	}
	fmt.Println("svr2, hello world")
	util.Func()
}
