package main 

import "fmt"

func main(){
    var Array[5] int = [5]int {1,2,3,4}
    fmt.Printf("type of Array = %T\n", Array)
    for index,value := range(Array){
        fmt.Println("index = ",index,"value = ",value)
    }
}
