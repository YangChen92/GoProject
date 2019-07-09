package main
import (
	"fmt"
)





func main(){
	var x interface{}

	var b2 float32 = 2.1
	x = b2

	if y,ok := x.(float32); ok{
		fmt.Println("convert success")
		fmt.Println("y 的类型是 %T 值是=%v",y,y)
	} else{
		fmt.Println("conver fail")
	}

	fmt.Println("继续执行...")
}

