package main 

import "fmt" 

func return_sigle(r1 int,r2 int ) (int , int){
    fmt.Println("r1 = ", r1,"r2 = ",r2)
    return 10,20 
}
func main(){
    value1,value2 := return_sigle(10,100)
    fmt.Println("value1 = " , value1 ,"value2 = " , value2);
}
