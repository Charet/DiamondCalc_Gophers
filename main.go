package main

import (
	"flag"
	"fmt"
)

type terminalValue struct {
	seed    int // 地图种子
	x       int // 查询x座标
	z       int // 查询y座标
	version string
}

var (
	terVar = terminalValue{}
	mul    = 25214903917
	mask   = (1 << 48) - 1
)

func mulandmask(a int) int {
	return (a*mul + 11) & mask
}

func diamond(mapseed, blockx, blockz, ver int) [2]int {
	temp := mapseed ^ mul&mask

	first := mulandmask(temp)
	second := mulandmask(first)
	i := (first >> 16 << 32) + (second << 16 >> 32) | 1

	third := mulandmask(second)
	fourth := mulandmask(third)
	j := (third >> 16 << 32) + (fourth << 16 >> 32) | 1

	temp = ((((16*blockx*i + 16*blockz*j) ^ mapseed) + ver) ^ mul) & mask
	relativex := mulandmask(temp) >> 44
	relativez := mulandmask(mulandmask(temp)) >> 44

	diamondx := relativex + 16*blockx
	diamondz := relativez + 16*blockz

	relative := [2]int{diamondx, diamondz}

	return relative
}

func lazuli(seed, blockx, blockz, ver int) [2]int {
	relative := diamond(seed, blockx, blockz, ver)
	lazulix := relative[0]
	diamondz := relative[1]
	var lazuliz int

	if diamondz%16 < 4 {
		lazuliz = diamondz + 16 - 4 + (diamondz % 16)
	} else {
		lazuliz = diamondz - 4
	}
	return [2]int{lazulix, lazuliz}
}

func init() {
	flag.IntVar(&terVar.seed, "s", 0, "Map seed")
	flag.IntVar(&terVar.x, "x", 0, "x coordinate")
	flag.IntVar(&terVar.z, "z", 0, "z coordinate")
	flag.StringVar(&terVar.version, "v", "1.16", "Game version")
	flag.Parse()
}

func main() {
	if terVar.seed != 0 || terVar.x != 0 || terVar.z != 0 {
		if terVar.version == "1.16" {
			fmt.Println("Diamond -> ", diamond(terVar.seed, terVar.x, terVar.z, 60009))
			fmt.Println("Lazuli -> ", lazuli(terVar.seed, terVar.x, terVar.z, 60009))
		} else if terVar.version == "1.17" {
			fmt.Println(diamond(terVar.seed, terVar.x, terVar.z, 60011))
		}
	} else {
		fmt.Println("Don't have input")
		return
	}
}
