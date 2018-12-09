package main

// import "fmt"
import (
	"fmt"
	"math/rand"
)

func main() {
   
   /* 这是我的第一个简单的程序 */
   fmt.Println("Hello, World!")
   fmt.Println("My favorite is",rand.Intn(100))
   refer()
}

func refer(){
   var a = 5
   fmt.Println(a)
   fmt.Println(&a)
}