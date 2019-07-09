package service

import(
	// "fmt"
	"go_code/project01/customer/model"
)

//该CustomerService,完成对Customer的操作，包括增删改查
type CustomerService struct{
	customers []model.Customer
	customerNum int

}

// func (this Customer) GetInfo() string{
// 	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",this.Id,
// 		this.Name,this.Gender,this.Age,this.Phone,this.Email)
// 		return info
// }