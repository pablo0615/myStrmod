// main.go - Main program
package main
import (
    "fmt"
    v1 "github.com/pablo0615/myStrmod"
)

func main() {
    s1 := v1.Concat("alpha", "bet")
    fmt.Println(s1)
    s2 := v1.Concatb("fast", "car")
    fmt.Println(s2)
    s3 := v1.Rep("=", 0)    // test patch
    fmt.Println(s3)
}
