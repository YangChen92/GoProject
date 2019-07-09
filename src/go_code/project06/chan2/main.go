package main
import(
	"fmt"
)

func main(){

	intchan := make(chan int,10)

	for i:=1;i<=10;i++{
		intchan <- i
	}
	// close(intchan)

	strchan := make(chan string,5)
	for i :=1;i<=5;i++{
		strchan <- "hello" + fmt.Sprintf("%d",i)
	}
	// close(strchan)
	for{
		select{
		case v := <-intchan :
			fmt.Printf("从intchan读取的数据%d\n",v)
		case v := <-strchan :
			fmt.Printf("从strchan读取的数据%s\n",v)
		default :
			fmt.Printf("都取不到数据\n")
			return
		}
	}

}