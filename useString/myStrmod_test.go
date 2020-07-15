// myStrmod_test.go - Test module 
package myStrings
import (
    "testing"
    v2 "github.com/pablo0615/myStrmod/v2"
)

func TestMyStrings( t *testing.T) {
    exp1 := "one"
    ret1 := v2.Concat("one")
    if ret1 != exp1 {
        t.Errorf("Concat() = %s, want %s", ret1, exp1)
    }

    exp2 := "two"
    ret2 := v2.Concatb("two")
    if ret2 != exp2 {
        t.Errorf("Concatb() = %s, want %s", ret2, exp2)
    }

    exp3 := "thirtysomething"
    ret3 := v2.Concat("thirty", "some", "thing")
    if ret3 != exp3 {
        t.Errorf("Concat() = %s, want %s", ret3, exp3)
    }

    exp4 := "big fence post"
    ret4 := v2.Concatb("big", "fence", "post")
    if ret4 != exp4 {
        t.Errorf("Concatb() = %s, want %s", ret4, exp4)
    }

    exp5 := "alphabetical"
    s1 := "alpha"
    v2.Update(&s1, "bet", "ical")
    if s1 != exp5 {
        t.Errorf("Update() = %s, want %s", s1, exp5)
    }

    exp6 := "alphabet soup"
    s2 := "alphabet"
    v2.Updateb(&s2, "soup")
    if s2 != exp6 {
        t.Errorf("Updateb() = %s, want %s", s2, exp6)
    }
}
