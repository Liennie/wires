package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func (a Point) Add(b Point) Point {
	return Point{
		x: a.x + b.x,
		y: a.y + b.y,
	}
}

func abs(i int) uint {
	if i < 0 {
		return uint(-i)
	}
	return uint(i)
}

func (a Point) Dist(b Point) uint {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func decode(path string) (dir Point, len int) {
	switch path[0] {
	case 'U':
		dir = Point{y: 1}
	case 'D':
		dir = Point{y: -1}
	case 'R':
		dir = Point{x: 1}
	case 'L':
		dir = Point{x: -1}
	}

	len, err := strconv.Atoi(path[1:])
	if err != nil {
		log.Panicf("Error: %v", err)
	}

	return dir, len
}

func part1(wires []string) uint {
	start := Point{0, 0}
	sb := uint8(255)

	board := map[Point]uint8{
		start: sb,
	}

	var minDist uint

	for wi, wire := range wires {
		wb := uint8(wi)
		position := start

		for _, path := range strings.Split(wire, ",") {
			dir, len := decode(path)

			for i := 0; i < len; i++ {
				position = position.Add(dir)
				if b, ok := board[position]; ok && b != wb {
					if b != sb {
						dist := position.Dist(start)
						if dist < minDist || minDist == 0 {
							minDist = dist
						}
					}
				} else {
					board[position] = wb
				}
			}
		}
	}

	return minDist
}

type Wire struct {
	dist uint
	have bool
}

type Board map[Point][2]Wire

func part2(wires []string) uint {
	start := Point{0, 0}
	board := Board{
		start: {Wire{}, Wire{}},
	}

	var minDist uint

	for wi, wire := range wires {
		position := start
		dist := uint(0)
		for _, path := range strings.Split(wire, ",") {
			dir, len := decode(path)

			for i := 0; i < len; i++ {
				position = position.Add(dir)
				dist += 1

				b := board[position]
				if !b[wi].have {
					b[wi].dist = dist
					b[wi].have = true
					board[position] = b
				}

				if b[1-wi].have {
					d := dist + b[1-wi].dist
					if d < minDist || minDist == 0 {
						minDist = d
					}
				}
			}
		}
	}

	return minDist
}

func main() {
	// wires := []string{
	// 	"R8,U5,L5,D3",
	// 	"U7,R6,D4,L4",
	// }
	// wires := []string{
	// 	"R75,D30,R83,U83,L12,D49,R71,U7,L72",
	// 	"U62,R66,U55,R34,D71,R55,D58,R83",
	// }
	// wires := []string{
	// 	"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
	// 	"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
	// }
	wires := []string{
		"R997,D443,L406,D393,L66,D223,R135,U452,L918,U354,L985,D402,R257,U225,R298,U369,L762,D373,R781,D935,R363,U952,L174,D529,L127,D549,R874,D993,L890,U881,R549,U537,L174,U766,R244,U131,R861,D487,R849,U304,L653,D497,L711,D916,R12,D753,R19,D528,L944,D155,L507,U552,R844,D822,R341,U948,L922,U866,R387,U973,R534,U127,R48,U744,R950,U522,R930,U320,R254,D577,L142,D29,L24,D118,L583,D683,L643,U974,L683,U985,R692,D271,L279,U62,R157,D932,L556,U574,R615,D428,R296,U551,L452,U533,R475,D302,R39,U846,R527,D433,L453,D567,R614,U807,R463,U712,L247,D436,R141,U180,R783,D65,L379,D935,R989,U945,L901,D160,R356,D828,R45,D619,R655,U104,R37,U793,L360,D242,L137,D45,L671,D844,R112,U627,R976,U10,R942,U26,L470,D284,R832,D59,R97,D9,L320,D38,R326,U317,L752,U213,R840,U789,L152,D64,L628,U326,L640,D610,L769,U183,R844,U834,R342,U630,L945,D807,L270,D472,R369,D920,R283,U440,L597,U137,L133,U458,R266,U91,R137,U536,R861,D325,R902,D971,R891,U648,L573,U139,R951,D671,R996,U864,L749,D681,R255,U306,R154,U706,L817,D798,R109,D594,R496,D867,L217,D572,L166,U723,R66,D210,R732,D741,L21,D574,L523,D646,R313,D961,L474,U990,R125,U928,L58,U726,R200,D364,R244,U622,R823,U39,R918,U549,R667,U935,R372,U241,L56,D713,L735,U735,L812,U700,L408,U980,L242,D697,L580,D34,L266,U190,R876,U857,L967,U493,R871,U563,L241,D636,L467,D793,R304,U103,L950,D503,R487,D868,L358,D747,L338,D273,L485,D686,L974,D724,L534,U561,R729,D162,R731,D17,R305,U712,R472,D158,R921,U827,L944,D303,L526,D782,R575,U948,L401,D142,L48,U766,R799,D242,R821,U673,L120",
		"L991,D492,L167,D678,L228,U504,R972,U506,R900,U349,R329,D802,R616,U321,R252,U615,R494,U577,R322,D593,R348,U140,L676,U908,L528,D247,L498,D79,L247,D432,L569,U206,L668,D269,L25,U180,R181,D268,R655,D346,R716,U240,L227,D239,L223,U760,L10,D92,L633,D425,R198,U222,L542,D790,L596,U667,L87,D324,R456,U366,R888,U319,R784,D948,R641,D433,L519,U950,L689,D601,L860,U233,R21,D214,L89,U756,L361,U258,L950,D483,R252,U206,L184,U574,L540,U926,R374,U315,R357,U512,R503,U917,R745,D809,L94,D209,R616,U47,R61,D993,L589,D1,R387,D731,R782,U771,L344,U21,L88,U614,R678,U259,L311,D503,L477,U829,R861,D46,R738,D138,L564,D279,L669,U328,L664,U720,L746,U638,R790,D242,R504,D404,R409,D753,L289,U128,L603,D696,L201,D638,L902,D279,L170,D336,L311,U683,L272,U396,R180,D8,R816,D904,L129,D809,R168,D655,L459,D545,L839,U910,L642,U704,R301,D235,R469,D556,L624,D669,L174,D272,R515,D60,L668,U550,L903,D881,L600,D734,R815,U585,R39,D87,R198,D418,L150,D964,L818,D250,L198,D127,R521,U478,L489,D676,L84,U973,R384,D167,R372,D981,L733,D682,R746,D803,L834,D421,R153,U752,L381,D990,R216,U469,L446,D763,R332,D813,L701,U872,L39,D524,L469,U508,L700,D382,L598,U563,R652,D901,R638,D358,L486,D735,L232,U345,R746,U818,L13,U618,R881,D647,R191,U652,R358,U423,L137,D224,R415,U82,R778,D403,R661,D157,R393,D954,L308,D986,L293,U870,R13,U666,L232,U144,R887,U364,L507,U520,R823,D11,L927,D904,R618,U875,R143,D457,R459,D755,R677,D561,L499,U267,L721,U274,L700,D870,L612,D673,L811,D695,R929,D84,L578,U201,L745,U963,L185,D687,L662,U313,L853,U314,R336",
	}

	// fmt.Println("Part 1: ", part1(wires))
	fmt.Println("Part 2: ", part2(wires))

	part1(wires)
	// part2(wires)
}
