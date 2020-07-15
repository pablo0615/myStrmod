// myStrmod.go - myStrmod module
package myStrmod

// Variadic functions
func Concat(s1 string, args ...string) string {
    result := s1
    for _, arg := range args {
        result += arg
    }
    return result
}

func Concatb(s1 string, args ...string) string {
    result := s1
    for _, arg := range args {
        result += " " + arg
    }
    return result
}

func Rep(ch string, num int) string {
    if num <= 1 { return ch }     // patch
    var result string
    for i:= 0; i < num; i++ {
        result += ch
    }
    return result
}

// New functions
func Update(s1 *string, args ...string) {
    for _, arg := range args {
        *s1 += arg
    }
}

func Updateb(s1 *string, args ...string) {
    for _, arg := range args {
        *s1 += " " + arg
    }
}
