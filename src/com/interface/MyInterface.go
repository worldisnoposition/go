package interfaceTest

import "fmt"

type IMy interface{
	my()
}

type Xiaohebo struct{
	name string
}

func(xiaohebo Xiaohebo) my(){
	fmt.Printf(xiaohebo.name)
}

func Info() {
	xiaohebo := Xiaohebo{}
	xiaohebo.name = "肖贺博"
	var m IMy
	m = xiaohebo
	m.my()
	j := JapanPerson{}
	j.sayHello();
	chinese := ChinesePerson{}
	chinese.sayHello()
}

//日本人
type JapanPerson struct {}
func (person JapanPerson)sayHello(){
   fmt.Println("japan Hello！")
}