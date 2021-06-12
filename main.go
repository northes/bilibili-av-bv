package main

import (
	"fmt"
	"math"
	"strconv"
)

var table = "fZodR9XQDSUm21yCkr6zBqiveYah8bt4xsWpHnJE7jL5VG3guMTKNPAwcF"
var tr = map[string]int{}
var s = []int{11, 10, 3, 8, 4, 6}
var xor = 177451812
var add = 8728348608

func init() {
	tableByte := []byte(table)
	for i := 0; i < 58; i++ {
		tr[string(tableByte[i])] = i
	}
}

func main() {
	//fmt.Println(tr)
	fmt.Println(bv2av("BV1px411A7ir"))
	fmt.Println(av2bv(1989613))
}

// bv转av
func bv2av(bv string) string {
	var r int
	arr := []rune(bv)

	for i := 0; i < 6; i++ {
		r += tr[string(arr[s[i]])] * int(math.Pow(float64(58), float64(i)))
	}
	return strconv.Itoa((r - add) ^ xor)
}

// av转bv
func av2bv(av int) string {
	x := (av ^ xor) + add
	r := []string{"B", "V", "1", " ", " ", "4", " ", "1", " ", "7", " ", " "}
	for i := 0; i < 6; i++ {
		r[s[i]] = string(table[int(math.Floor(float64(x/int(math.Pow(float64(58), float64(i))))))%58])
	}
	var result string
	for i := 0; i < 12; i++ {
		result += r[i]
	}
	return result
}
