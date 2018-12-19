package channelTest

import (
    "fmt"
    "strconv"
)
var i = 0
func sum(a []int, c chan int) {//注意这里的c chan int
    sum := 0
    var t = i
    i=i+1
    for _, v := range a {
        sum += v
        fmt.Println("v="+strconv.Itoa(v)+",t="+strconv.Itoa(t))
    }
    fmt.Println("sum="+strconv.Itoa(sum))
    c <- sum  // send sum to c
}
func Test() {
    a := []int{7, 2, 8, -9, 4, 0,1,3}
    c := make(chan int)
    for _, v := range a[:len(a)/2] {
        fmt.Println("数组里v="+strconv.Itoa(v))
    }
    go sum(a[0:2], c)
    go sum(a[2:4], c)
    go sum(a[4:6], c)
    go sum(a[6:8], c)
    w,x,y,z:= <-c, <-c, <-c, <-c  // receive from c
    fmt.Println(w,x,y,z)//因为是多线程，顺序是不一定的
}