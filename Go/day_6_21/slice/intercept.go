package main 

import "fmt"

func main(){
    slice := make([]int , 3 , 5)
    slice = append(slice,2,3,4)
    //slice[4] = 2147483647 
    fmt.Println("len = ",len(slice),"capacity = ",cap(slice))
    fmt.Printf("slice = turth %v",slice)
    s1 :=  slice[0:2]
    fmt.Printf("s1 = turth %v",s1)
}
