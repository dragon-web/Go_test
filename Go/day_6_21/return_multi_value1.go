package main 

import "fmt" 

func return_sigle(r1 int,r2 int ) (va1 int ,va2 string){
    fmt.Println("r1 = ", r1,"r2 = ",r2)
    va1 = 12;
    va2 = "hello world"
    return va1,va2 
}
func main(){
    value1,value2 := return_sigle(10,100)
    fmt.Println("value1 = " , value1 ,"value2 = " , value2);
}
