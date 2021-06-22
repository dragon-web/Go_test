package main 

import "fmt"

func  printArr(arr1 [4]int){
    arr1[0] =  123
    for index,value := range(arr1){
        fmt.Println("index = ", index,"value = ",value)
    }
}
func main(){
    array := [4]int{1,2,3}
    printArr(&array)
    fmt.Println(array[0])
}
