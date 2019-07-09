package main
import "fmt"
// import "errors"
// import "strings"
// import "sort"
import "encoding/json"
// func main(){
// 	// fmt.Println("hello world")
// 	fmt.Println("姓名\t性别\t籍贯\t住址\t\n杨\t男\t连云港\t北京")
// }
// func main(){
// 	var age int
// 	fmt.Println("请输入年龄")
// 	fmt.Scanln(&age)

// 	if age >=18 {
// 		fmt.Println("大于等于18")
// 	}
// }

// func cal(n1 float64,n2 float64, operator byte) float64 {
// 	var res float64
// 	switch operator {
// 	case '+':
// 		res = n1 + n2
// 	case '-':
// 		res = n1 - n2
// 	case '*':
// 		res = n1 * n2
// 	case '/':
// 		res = n1 / n2
// 	default:
// 		fmt.Println("操作符有误")
// 	}
// 	return res
// }

// func test(n int){
// 	if n > 2{
// 		n--
// 		test(n)
// 	}

// 	fmt.Println("n=",n)
// }

// func fbn(n int) int {
// 	if n == 1 || n == 2{
// 		return 1
// 	}else{
// 		return fbn(n-1) + fbn(n - 2)
// 	}
// }
// func getSum(n1 int,n2 int) int {
// 	return n1 + n2
// }

// func myFun(funvar func(int,int) int,num1 int, num2 int) int{
// 	return funvar(num1,num2)
// }

// func main(){
	// a := getSum
	// fmt.Printf("a的类型%T,getSum类型是%T\n",a,getSum)
	// // res := a(10,20)
	// res := myFun(getSum,11,22)
	// fmt.Println("res=",res)

	// res1 :=func(n1 int,n2 int) int{
	// 	return n1+n2
	// }(10,11)
	// fmt.Println("res1=",res1)

	// a := func(n1 int,n2 int)int{
	// 	return n1+n2
	// }

	// res1 := a(11,22)
	// res2 := a(22,33)

	// fmt.Printf("res1=%d,res2=%d",res1,res2)
// }

// func makeSuffix(suffix string) func(string) string{
// 	return func (name string) string{
// 		if(!strings.HasSuffix(name,suffix)){
// 			return name + suffix
// 		}
// 		return name
// 	}
// }

// func main(){
// 	f := makeSuffix(".jpg")
// 	name_1 := f("qwer")
// 	name_2 := f("asd.jpg")
// 	fmt.Println(name_1,name_2)
// }


// func test(){
// 	defer func(){
// 		err := recover()
// 		if err !=nil{
// 			fmt.Println("err=",err)
// 			fmt.Println("send email")
// 		}
// 	}()

// 	num1 := 10
// 	num2 := 0
// 	res := num1/num2
// 	fmt.Println("res=",res)
// }

// func readConf (name string) (err error){
// 	if name == "conf.ini"{
// 		return nil
// 	}else{
// 		return errors.New("file name error")
// 	}
// }

// func test(){
// 	var score [5]float64

// 	for i:=0;i<len(score);i++{
// 		fmt.Printf("输入第%d个数：",i+1)
// 		fmt.Scanln(&score[i])
// 	}

// 	for i:=0;i<len(score);i++{
// 		fmt.Printf("score[%d]=%v\n",i,score[i])
// 	}
// }

func test1(){
	var num1 [3]int = [3]int{1,2,3}
	fmt.Println("num1=",num1)
	var num2 = [3]int{4,5,6}
	fmt.Println("num2=",num2)
	var num3 = [...]int{7,8,9}
	fmt.Println("num3=",num3)
	var num4 = [...]int{1:10,0:11,2:12}
	fmt.Println("num4=",num4)
	str05 := [...]string{1:"tom",0:"jack",2:"rose",4:"mary",3:"andy"}
	fmt.Println("str05=",str05)
}

func test2(){
	heroes := [...]string{"关","张","赵"}

	for index,value := range heroes{
		fmt.Printf("index=%v value=%v",index,value)
		fmt.Printf("heroes[%d]=%v\n",index,value)
	}

	for _,v := range heroes{
		fmt.Printf("元素的值=%v\n",v)
	}
}


func test3(){
	// var slice3 []int = []int{11,22,33}
	// slice3 = append(slice3,44,55)
	// fmt.Println("slice3",slice3)
	var slice4 []int = []int{1,2,3,4,5}
	var slice5 = make([]int,10)
	copy(slice5,slice4)
	fmt.Println("slice4",slice4)
	fmt.Println("slice5",slice5)
}

func BubbleSort(arr *[5]int){
	fmt.Println("排序前arr=",(*arr))
	temp := 0
	for i := 0;i < len(*arr) - 1;i++{
		for j := 0;j < len(*arr)-1 -i;j++{
			if(*arr)[j] > (*arr)[j + 1]{
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
}

func DouArr(){
	// var arr [4][6]int
	// arr[1][2] = 1
	// arr[2][1] = 2
	// arr[2][3] = 3

	// for i:=0;i<4;i++{
	// 	for j:=0;j<6;j++{
	// 		fmt.Print(arr[i][j]," ")
	// 	}
	// 	fmt.Println()
	// }
	var arr = [2][3]int{{1,2,3},{4,5,6}}

	for i := 0;i<len(arr);i++{
		for j := 0; j<len(arr[i]);j++{
			fmt.Printf("%v\t",arr[i][j])
		}
		fmt.Println()
	}

	for i,v := range arr{
		for j,v2 := range v{
			fmt.Printf("arr[%v][%v]=%v",i,j,v2)
		}
		fmt.Println()
	}

}


func ClassRoom(){
	var arr = [3][5]int{{5,7,5,7,6},{2,4,2,4,3},{6,8,6,8,7}}

	for i,v1 := range arr{
		totle := 0
		for _,v2 := range v1{
			totle += v2
		}
		average := totle/len(v1)
		fmt.Printf("%d班总分=%d,平均分=%d",i+1,totle,average)
		fmt.Println()
	}
}

func TestMap(){
	var a map[string]string
	a = make(map[string]string,10)
	a["no1"] = "关羽"
	a["no2"] = "张飞"
	a["no3"] = "赵云"
	fmt.Println(a)

	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "广州"
	fmt.Println(cities)

	heroes := map[string]string{
		"hero1" : "毛",
		"hero2" : "朱",
		"hero3" : "周", 
	}
	heroes["no4"] = "彭"
	fmt.Println("heroes=",heroes)


}

func StudentMap(){
	student := make(map[string]map[string]string)
	student["stu01"] = make(map[string]string,3)
	student["stu01"]["name"] = "yang"
	student["stu01"]["sex"] = "男"
	student["stu01"]["arrdess"] = "北京"

	student["stu02"] = make(map[string]string,3)
	student["stu02"]["name"] = "chen"
	student["stu02"]["sex"] = "男"
	student["stu02"]["address"] = "shanghai"

	// fmt.Println(student)
	// fmt.Println(student["stu02"])
	// fmt.Println(student["stu02"]["name"])
	fmt.Println(len(student["stu01"]))

	// for key1,value1 := range student{
	// 	fmt.Printf("k1=%v  v1=%v",key1,value1)
	// 	for key2,value2 := range value1{
	// 		fmt.Printf("\t k2=%v  v2=%v",key2,value2)
	// 	}
	// }
}


type Stu struct{
	Name string
	Age int
	Address string
}

type cat struct{
	name string
	color string
	age int
}

type Person struct{
	name string
	age int
	score [5]float64
	ptr *int
	slice []int
	map1 map[string]string
}

type Monster struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}

func main(){
	//1。创建一个Monster变量
	monster := Monster{"牛魔王",500,"芭蕉扇"}
	//2.将monster变量序列号为json格式字符串
	// json.Marshal 函数中的反射
	jsonStr, err := json.Marshal(monster)
	if err!= nil{
		fmt.Println("json 处理错误",err)
	}
	fmt.Println("jsonStr",string(jsonStr))
}


// func main(){
// 	// StudentMap()
// var p1 Person
// fmt.Println(p1)
// 	if p1.ptr == nil{
// 		fmt.Println("ok1")
// 	}
// 	if p1.slice == nil{
// 		fmt.Println("ok2")
// 	}
// 	if p1.map1 == nil{
// 		fmt.Println("ok3")
// 	}

// 	p1.slice = make([]int,10)
// 	p1.slice[0] = 100
	
// 	p1.map1 = make(map[string]string)
// 	p1.map1["key1"] = "yang"

// 	fmt.Println(p1)

// 	cat1 := cat{"xiaohuang","yellow",5}

// 	fmt.Println(cat1)

// 	var p3 *Person = new(Person)

// 	(*p3).name = "yc"
// 	p3.age = 28

// 	fmt.Println(*p3)


// }





