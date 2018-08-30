package composite

import "fmt"

func AppendRune() {
	str := "Shanaka Rusith"
	var runeSlice []rune
	for _, cha := range str {
		runeSlice = append(runeSlice, cha)
	}
	fmt.Printf("%q\n", runeSlice)
}


func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = z[:zlen]
	} else {
		zcap := zlen
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z;
}

func AppendInt() {
	var x, y []int
	for i := 0; i < 10 ; i ++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}