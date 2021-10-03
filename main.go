package main

import "fmt"

var mul = 25214903917
var mask = (1 << 48) - 1

func mulandmask(a int) int {
	return (a*mul + 11) & mask
}

func clac(mapseed, blockx, blockz int) [2]int {
	temp := mapseed ^ mul&mask

	first := mulandmask(temp)
	second := mulandmask(first)
	i := (first >> 16 << 32) + (second << 16 >> 32) | 1

	third := mulandmask(second)
	fourth := mulandmask(third)
	j := (third >> 16 << 32) + (fourth << 16 >> 32) | 1

	temp = ((((16*blockx*i + 16*blockz*j) ^ mapseed) + 60009) ^ mul) & mask
	relativex := mulandmask(temp) >> 44
	relativez := mulandmask(mulandmask(temp)) >> 44

	diamondx := relativex + 16*blockx
	diamondz := relativez + 16*blockz

	relative := [2]int{diamondx, diamondz}

	return relative
}

func main() {
	var seed, x, z int
	fmt.Print("请输入种子,x坐标,z坐标:")
	fmt.Scanf("%d %d %d", &seed, &x, &z)

	fmt.Println(clac(seed, x, z))
}
