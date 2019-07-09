package main
import (
	"fmt"
)

func main(){
	intchan := make(chan int,50)
	exitchan := make(chan bool,1)

	go writeData(intchan)
	go readData(intchan,exitchan)

	for{
		_,ok:= <- exitchan
		if !ok{
			break
		}
	}

}

func writeData(intchan chan int){
	for i := 1;i<=50;i++{
		intchan<- i
	}
	close(intchan)
}
func readData(intchan chan int,exitchan chan bool){
	for{
		v,ok :=<-intchan
		if !ok{
			break
		}
		fmt.Printf("readData读到数据=%v\n",v)
	}

	exitchan<- true
	close(exitchan)
}