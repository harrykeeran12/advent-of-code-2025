package day3

import (
	"testing"
)

// func TestInput1(t *testing.T) {
//     ans := largestJolts("987654321111111")
//     if ans != 98 {
//         t.Errorf("TestPart1 wrong, wanted 98, got %d", ans)
//     }
// }

// func TestInput2(t *testing.T) {
//     ans := largestJolts("811111111111119")
//     if ans != 89 {
//         t.Errorf("TestPart1 wrong, wanted 89, got %d", ans)
//     }
// }

// func TestInput3(t *testing.T) {
//     ans := largestJolts("818181911112111")
//     if ans != 92 {
//         t.Errorf("TestPart1 wrong, wanted 92, got %d", ans)
//     }
// }

// func TestInput4(t *testing.T) {
//     ans := largestJolts("234234234234278")
//     if ans != 78 {
//         t.Errorf("TestPart1 wrong, wanted 78, got %d", ans)
//     }
// }

// func TestInput5(t *testing.T){
// 	ans := part1("input.txt")
// 	if ans != 357{
// 		t.Errorf("TestPart5 wrong, wanted 357, got %d", ans)
// 	}
// }

// func TestInput6(t *testing.T){
// 	ans := part1("input2.txt")
// 	if ans != 17109{
// 		t.Errorf("TestPart5 wrong, wanted 357, got %d", ans)
// 	}
// }

func Test2Input1(t *testing.T) {
	ans := largestTwelveJolts("987654321111111")
	if ans != 987654321111 {
		t.Errorf("Test2Input1 wrong, wanted 987654321111, got %d", ans)
	}
}

func Test2Input2(t *testing.T) {
	ans := largestTwelveJolts("811111111111119")
	if ans != 811111111119 {
		t.Errorf("Test2Input1 wrong, wanted 811111111119, got %d", ans)
	}
}

func Test2Input3(t *testing.T) {
	ans := largestTwelveJolts("234234234234278")
	if ans != 434234234278 {
		t.Errorf("Test2Input1 wrong, wanted 434234234278, got %d", ans)
	}
}

func Test2Input4(t *testing.T) {
	ans := largestTwelveJolts("818181911112111")
	if ans != 888911112111 {
		t.Errorf("Test2Input1 wrong, wanted 888911112111, got %d", ans)
	}
}
