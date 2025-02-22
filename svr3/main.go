package svr2

import "fmt"

func main() {
	m := make(map[string]string)
	for k, _ := range m {
		fmt.Println(k)
	}
	fmt.Println("svr2, hello world")
}
