// main.go - Main program
package main
import (
    "fmt"
    v1 "github.com/pablo0615/myStrmod"
)

func main() {
    s1 := "alpha"
    v1.Update(&s1, "bet")
    fmt.Println(s1)
    s2 := "book"
    v1.Updateb(&s2, "deal")
    fmt.Println(s2)
}
