package main
import(
	"fmt"
	"reflect"
)

func reflectTest(b interface{}){

	rType := reflect.TypeOf(b)
	fmt.Println("rType=",rType)
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=",rVal)

	n1 := 2+rVal.Int()

	fmt.Printf("rval=%v type=%T\n",rVal,rVal)
	fmt.Println("Value=",n1)

	iv := rVal.Interface() 
	num2 := iv.(int)
	fmt.Println("iv=",num2)

}

func reflectTest01(b interface{}){

	rType := reflect.TypeOf(b)
	fmt.Println("rType=",rType)
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=",rVal)

	iv := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T \n",iv,iv)

	stu,ok := iv.(Student)
	if ok{
		fmt.Printf("stu.Name=%v\n",stu.Name)
	}

	

}

func reflectTest02(b interface{}){
	rVal := reflect.ValueOf(b)
	rType := reflect.TypeOf(b)
	fmt.Printf("rValue=%v rType=%v \n",rVal,rType)

	iv := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T \n",iv,iv)

	num := iv.(float64)
	fmt.Println("num=",num)
}

type Student struct{
	Name string
	Age int
}

func main(){

	// var num int = 10
	// reflectTest(num)
	// stu := Student{
	// 	Name :"yang",
	// 	Age : 10,
	// }
	// reflectTest01(stu)
	// var num1 float64 = 12.12
	// reflectTest02(num1)

	var str string = "yang"
	fs := reflect.ValueOf(&str)
	fs.Elem().SetString("chen")
	fmt.Printf("%v \n",str)

}