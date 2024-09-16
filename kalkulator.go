package main 

import "fmt"

func main(){

    var number1,number2 float64
    var operator string

    fmt.Print("Enter the first number : ")
    fmt.Scanln(&number1)

    fmt.Print("Enter the second number : ")
    fmt.Scanln(&number2)

    fmt.Print("Enter the Operator ( + - * / ) : ")
    fmt.Scanln(&operator)

    switch operator {

    case "+" : 
        fmt.Printf("%f %s %f = %f", number1, operator, number2, number1 + number2)

    case "-" : 
        fmt.Printf("%f %s %f = %f", number1, operator, number2, number1 - number2)

    case "*" : 
        fmt.Printf("%f %s %f = %f", number1, operator, number2, number1 * number2) 

    case "/": 
        if number2 == 0.0{
            fmt.Println("Divide by Zero Situation")    
        }else{
            fmt.Printf("%f %s %f = %f", number1, operator, number2, number1 / number2) 
        }

default: 
    fmt.Println("Invalid Operator")
    }
}