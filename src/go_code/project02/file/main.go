package main
import(
	"fmt"
	"os"
	"bufio"
	"io"
	// "io/ioutil"
)

// func main(){
// 	file , err := os.Open("d:/test.txt")
// 	if err != nil {
// 		fmt.Println("err=",err)
// 	}

// 	fmt.Printf("file=%v",file)

// 	//关闭
// 	defer file.Close()

// 	//创建一个 *reader ,是带缓冲的
// 	/*
// 	*
// 	const(
// 		defaultBufSize = 4096 //默认的缓冲区大小为4096
// 	)
// 	*/
// 	reader := bufio.NewReader(file)

// 	for{
// 		str, err := reader.ReadString('\n')//读到一个换行就结束
// 		//输出内容
// 		fmt.Print(str)
// 		if err == io.EOF{//io.EOF表示文件的末尾
// 			break
// 		}
// 	}

// 	fmt.Println("文件读取结束")

// }

// func main(){
// 	//使用ioutil.ReadFile一次性将文件读取到位
// 	file := "d:/test.txt"
// 	content, err := ioutil.ReadFile(file)//文件的open和close被封装在了ReadFile
// 	if err != nil{
// 		fmt.Printf("read file err=%v",err)
// 	}
	
// 	fmt.Printf("%v",content) //[]byte
// 	fmt.Printf("%v",string(content))
// }

// func main1(){
// 	filepath := "d:/write.txt"
// 	content := "hello yang"
	
// 	file, err := os.OpenFile(filepath, os.O_WRONLY | os.O_CREATE, 0666)
// 	if err != nil{
// 		fmt.Printf("open file err=%v\n",err)
// 	}

// 	defer file.Close()

// 	// 写入时，使用带缓存的 *write
// 	write := bufio.NewWriter(file)
// 	for i:=0;i<5;i++{
// 		write.WriteString(content)
// 	}

// 	//因为write是带缓存的，因此在调用WriteString方法时，其实内容是先写入到缓存的
// 	//所以需要调用flush方法麻将缓冲的数据真正写入到文件
// 	write.Flush()


// }

// func main(){
// 	filePath := "d:/write1.txt"
// 	content := "hello qqq"
// 	file,err := os.OpenFile(filePath,os.O_WRONLY | os.O_CREATE,0666)

// 	if err != nil{
// 		fmt.Printf("open file err=%v\n",err)
// 	}

// 	defer file.Close()

// 	write := bufio.NewWriter(file)

// 	write.WriteString(content)

// 	write.Flush()

// 	file1 , err := os.OpenFile(filePath,os.O_CREATE | os.O_WRONLY,0666)
	
// 	if err != nil{

// 	}

// 	write1 := bufio.NewWriter(file1)

// 	write1.WriteString(content)

	
// }

// func main(){

// 	readFile := "d:/read.txt"
// 	writeFile := "d:/write.txt"

// 	data, err := ioutil.ReadFile(readFile)

// 	if err != nil{
// 		//说明读取文件有错误
// 		fmt.Printf("read file err=%v\n",err)
// 		return
// 	}

// 	err = ioutil.WriteFile(writeFile,data,0666)

// 	if err != nil{
// 		fmt.Printf("write file error=%v\n",err)
// 	}


// }
type count struct{
	ChCount int
	NumCount int
	SpaceCount int
	OtherCount int

}
func main(){

	filePath := "d:/read.txt"

	file,err := os.Open(filePath)


	if err != nil{
		fmt.Printf("file open err=%v",err)
	}

	defer file.Close()

	read := bufio.NewReader(file)

	var count count

	for{
		str,err := read.ReadString('\n')
		if err == io.EOF{
			break
		}

		// str = []rune(str)

		for _,v := range str{

			switch  {
			case v>='a' && v <='z':
				fallthrough
			case v>='A' && v<='Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v>='0' && v<= '9':
				count.NumCount++
			default:
				count.OtherCount++
				
			}

		}
	}

	fmt.Printf("Char=%v num=%v space=%v other=%v",count.ChCount,count.NumCount,count.SpaceCount,count.OtherCount)
	

}