package day13

import (
	"os"
	"strconv"
	"strings"
)

// Linear algebra
// Let's say we need 'a' steps of vector A and 'b' steps of vector B.
// This gives us the equation:
// a(10,5) + b(6,8) = (250,150)

// Breaking this into x and y components:
// 10a + 6b = 250  (equation 1)
// 5a + 8b = 150   (equation 2)

// We can solve this using substitution or matrix methods. Let's use matrices:
// [10  6][a] = [250]
// [5   8][b]   [150]

// To solve this, we need to find the inverse of the coefficient matrix:
// det = (10)(8) - (6)(5) = 80 - 30 = 50

// inverse = (1/50)[8    -6]
//                 [-5   10]

// Multiplying:
// [a] = (1/50)[8    -6][250]
// [b]        [-5   10][150]

// a = (8(250) + -6(150))/50
// b = (-5(250) + 10(150))/50

// Let's calculate:
// a = (2000 - 900)/50 = 22
// b = (-1250 + 1500)/50 = 5

// Therefore, you need:

// 22 steps of vector A (10,5)
// 5 steps of vector B (6,8)

// To verify:
// 22(10,5) + 5(6,8)
// = (220,110) + (30,40)
// = (250,150)
// The minimum time would be 27 steps total (22 + 5).
func Part2() int {
	dat, err := os.ReadFile("./day13/input")
	if err != nil {
		panic(err)
	}

	tenTrillion := 10_000_000_000_000
	var machines []machine
	values := strings.Split(string(dat), "\n")
	for i := 0; i < len(values); i += 4 {
		// fmt.Println("asd1")
		a := strings.Split(values[i], " ")
		xA := strings.Split(a[2], "+")
		yA := strings.Split(a[3], "+")
		XAInt, _ := strconv.Atoi(xA[1][:len(xA)])
		YAInt, _ := strconv.Atoi(yA[1])

		// fmt.Println("asd2")
		b := strings.Split(values[i+1], " ")
		xB := strings.Split(b[2], "+")
		yB := strings.Split(b[3], "+")
		XBInt, _ := strconv.Atoi(xB[1][:len(xB)])
		YBInt, _ := strconv.Atoi(yB[1])

		// fmt.Println("asd3")
		prize := strings.Split(values[i+2], " ")
		xPrize, _ := strconv.Atoi(prize[1][2 : len(prize[1])-1])
		yPrize, _ := strconv.Atoi(prize[2][2:])

		machines = append(machines, machine{
			a:     point{XAInt, YAInt},
			b:     point{XBInt, YBInt},
			prize: point{xPrize + tenTrillion, yPrize + tenTrillion},
		})
	}

	var total int
	for _, mach := range machines {
		total += countLA2(mach)
	}

	return total
}

func countLA2(mach machine) int {
	// divisor = (aX * bY) - (aY * bX)
	div := (mach.a.x * mach.b.y) - (mach.a.y * mach.b.x)
	// inverse the matrix then multiply it
	// a = ((bY * PrizeX) - (-bX * PrizeY)) / divisor
	a := ((mach.b.y * mach.prize.x) + ((-mach.b.x) * mach.prize.y)) / div
	// b = (((-aY) * PrizeX) + (aX * PrizeY)) / divisor
	b := (((-mach.a.y) * mach.prize.x) + (mach.a.x * mach.prize.y)) / div

	resultX := (mach.a.x * a) + (mach.b.x * b)
	resultY := (mach.a.y * a) + (mach.b.y * b)
	if resultX != mach.prize.x && resultY != mach.prize.y {
		// fmt.Println("not same, a:", a, "b:", b, "resultX:", resultX, "resultY:", resultY, "prizeX:", mach.prize.x, "prizeY:", mach.prize.y)
		return 0
	}

	// fmt.Println("same, a:", a, "b:", b, "resultX:", resultX, "resultY:", resultY, "prizeX:", mach.prize.x, "prizeY:", mach.prize.y)
	if a < 0 || b < 0 {
		return 0
	}

	return (a * 3) + b
}
