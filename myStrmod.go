// myStrmod.go - myStrmod module
package myStrmod

func Concat(s1, s2 string) string {
    return s1 + s2
}

func Concatb(s1, s2 string) string {
    return s1 + " " + s2
}

func Rep(ch string, num int) string {
    if num <= 1 { return ch }     // patch
    var result string
    for i:= 0; i < num; i++ {
        result += ch
    }
    return result
}
