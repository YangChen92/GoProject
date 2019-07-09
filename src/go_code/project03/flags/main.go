package main

import(
	"fmt"
	"flag"
)

func main(){
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user,"u","","user")
	flag.StringVar(&pwd,"pwd","","password")
	flag.StringVar(&host,"h","localhost","host")
	flag.IntVar(&port,"port",3306,"port")
	flag.Parse()

	fmt.Printf("user=%v pwd=%v host=%v port=%v",user,pwd,host,port)
}

