package main 

import(
    "fmt"
)

func main(){
    defer fmt.Println("end differ")
    fmt.Println("start")
}
