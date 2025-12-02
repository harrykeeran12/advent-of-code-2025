package day2

import (
	"fmt"
	"testing"
)

func TestInput1(t *testing.T) {
    ans := part_1("input.txt")
    if ans != 1227775554 {
        t.Errorf("TestPart1 wrong, wanted 1227775554, got %d", ans)
    }
}
// func TestInput2(t *testing.T) {
//     // ans := part_1("input2.txt")
//     // fmt.Println(ans)
// }

func Test2Input1(t *testing.T) {
    ans := part_2("input.txt")
    if ans != 4174379265 {
        t.Errorf("TestPart2 wrong, wanted 4174379265, got %d", ans)
    }
}
func Test2Input2(t *testing.T) {
    ans := part_2("input2.txt")
    fmt.Println(ans)
}

func Test2Invalid(t *testing.T){
	ans := isNewInvalid(2121212121)
	if ans != true{
		t.Errorf("Test2Input2 wrong, wanted true, got %v", ans)
	}
}
