package main 

import "fmt"

func printArray(Array []int){
    for _,value := range(Array){
        fmt.Println("value = ",value)
        Array[0] = 222
    }
}
func main(){
    myArray := []int{1,2,34,4}
    printArray(myArray)
    fmt.Println("After")
    fmt.Println("myArray[0] = " , myArray[0])
}
