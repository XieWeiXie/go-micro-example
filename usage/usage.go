package usage

import (
	"fmt"

	"gopkg.in/wuxiaoxiaoshen/go-version.v1.v1/common"
)

func NewHello() {
	var hello common.Hello
	fmt.Println(hello.Print())
	fmt.Println(hello.Println("Python"))
}
