package main

import (
	"flag"
	"fmt"
)

type terminalValue struct {
	seed int // 地图种子
	x    int // 查询x座标
	z    int // 查询y座标
}

var (
	terVar = terminalValue{}
	mul    = 25214903917
	mask   = (1 << 48) - 1
)

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

func init() {
	flag.IntVar(&terVar.seed, "s", 0, "Map seed")
	flag.IntVar(&terVar.x, "x", 0, "x coordinate")
	flag.IntVar(&terVar.z, "z", 0, "z coordinate")
	flag.Parse()
}

func main() {
	if terVar.seed != 0 || terVar.x != 0 || terVar.z != 0 {
		fmt.Println(clac(terVar.seed, terVar.x, terVar.z))
	} else {
		fmt.Println("Don't have input")
		return
	}
}
