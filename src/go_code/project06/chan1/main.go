package main

import(
	"fmt"
)

func main(){

	putchan := make(chan int,1000)
	primechan := make(chan int,2000)
	exitchan := make(chan bool,4)

	go putNum(putchan)

	for i:=0;i<4;i++{
		go primeNum(putchan,primechan,exitchan)
	}

	go func(){
		for i:=0;i<4;i++{
			<-exitchan
		}

		close(primechan)
	}()

	for{
		res ,ok := <- primechan
		if !ok{
			break
		}

		fmt.Printf("素数=%d\n",res)
	}
	fmt.Println("main线程退出")



}

func putNum(putchan chan int){
	for i := 1;i<8000;i++{
		putchan <- i
	}

	close(putchan)
}

func primeNum(putchan chan int,primechan chan int,exitchan chan bool){

	var flag bool
	for{

		num ,ok := <- putchan
		if !ok{
			break
		}
		flag = true
		for i:=2;i<num;i++{
			if num%i == 0{
				flag = false
				break
			}
		}

		if flag{
			primechan<- num
		}
	}

	fmt.Println("一个primeNum协程因取不到数据而退出")
	exitchan<-true
}
