package main

import "fmt"

var (
 nome string
 n1 int
 n2 int32
)

func main(){
 var mensagem string = "Aula 03 - Go 101"
 fmt.Println(mensagem)

 var b, f, s = true, 2.3, "Ölá"
 fmt.Println(b, f, s)

 var total int
 total++
 fmt.Println("total:", total)

 nome = "Luiz Gustavo"
 fmt.Println("Hello", nome, "!")

 var x, y = 10, 20
 fmt.Println(x, y)
 x, y = y, x
 fmt.Println(x, y)
}
