package main

import (
"flag"
"fmt"
)

var (
	h string
	p int
	u string
	P string
)

func main() {
	flag.StringVar(&h, "h", "localhost", "端口号")
	flag.StringVar(&u, "u", "wendao", "用户")
	flag.StringVar(&P, "P", "wendao", "密码")
	flag.Parse()
	fmt.Println(h, u, p)
}
