package main
import (
	"fmt"
)

type Goods struct{
	Name string
	price float64
}

type Brand struct{
	Name string
	Address string
}

type TV struct{
	Goods
	Brand
}

type TV2 struct{
	*Goods
	*Brand
}


func main(){
	key := ""
	fmt.Scanln(&key)
	tv := TV{Goods{"电视机001",3000},Brand{"Haier","beijing"},}
	tv2 :=TV2{&Goods{"电视机002",5000},&Brand{"长虹","shanghai"},}

	fmt.Println("tv=",tv.price)
	fmt.Println("tv2=",*tv2.Goods)
	fmt.Println("key=",key)

}