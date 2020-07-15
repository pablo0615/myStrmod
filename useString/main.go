// main.go - Main program
package main
import (
    "fmt"
    v1 "github.com/pablo0615/myStrmod"
    v2 "github.com/pablo0615/myStrmod/v2"
)

func main() {
    s1 := v2.Concat("thirty", "some", "thing")
    fmt.Println(s1)
    s2 := v2.Concatb("big", "fence", "post")
    fmt.Println(s2)
    s3 := "alpha"
    v2.Update(&s3, "bet", "ical")
    fmt.Println(s3)
    s4 := v1.Rep("=", 10)
    fmt.Println(s4)
}
