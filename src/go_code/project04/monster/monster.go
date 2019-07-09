package monster
import(
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Monster struct{
	Name string
	Age int
	Skill string
}

func (this *Monster) Store() bool{

	code,err := json.Marshal(this)

	if err != nil{
		fmt.Println("marshal err=",err)
		return false
	}

	//保存文件
	filePath := "d:/monster.txt"
	err = ioutil.WriteFile(filePath,code,0666)
	if err != nil{
		fmt.Println("write file err=",err)
		return false
	}

	return true

}

func (this *Monster) ReStore() bool{

	filePath := "d:/monster.txt"
	data ,err := ioutil.ReadFile(filePath)
	if err != nil{
		fmt.Println("ReadFile err=",err)
		return false
	}

	err = json.Unmarshal(data,this)
	if err != nil{
		fmt.Println("UnMarshal err=",err)
		return false
	}

	return true
}