package main

import(
	"fmt"
	"os"
)

func main(){
	fmt.Println("参数个数=",len(os.Args))

	for i ,v := range os.Args{
		fmt.Printf("args[%v]=%v\n",i,v)
	}
}