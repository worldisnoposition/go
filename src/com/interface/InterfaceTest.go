package interfaceTest

import "fmt"

type ISayHello interface {
   sayHello()
}

//美国人
type AmericalPerson struct {}
func (person AmericalPerson)sayHello(){
   fmt.Println("Hello！")
}
//中国人
type ChinesePerson struct {}
func (person ChinesePerson)sayHello(){
   fmt.Println("你好！")
}

func greet(i ISayHello){
   i.sayHello()
}

func Test() {
    ameriacal := AmericalPerson{}
    chinese := ChinesePerson{}
    var i ISayHello
    i = ameriacal
    i.sayHello()
    i = chinese
    i.sayHello()
}
