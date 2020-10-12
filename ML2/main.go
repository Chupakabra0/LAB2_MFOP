package main

import (
	"fmt"
	"math"
)

type ThreeElementInt64Set struct {
	a, b, c int64
}

func CountDifference(set ThreeElementInt64Set, number int64) ThreeElementInt64Set {
	return ThreeElementInt64Set{ int64(math.Abs(float64(set.a - number))), int64(math.Abs(float64(set.b - number))), int64(math.Abs(float64(set.c - number))) }
}

func FindNum(set ThreeElementInt64Set, dif ThreeElementInt64Set, pred func(int64, int64) bool) int64 {
	if pred(dif.a, dif.b) {
		if pred(dif.a, dif.c) {
			return dif.a
		} else {
			return dif.c
		}
	} else if pred(dif.b, dif.c) {
		return dif.b
	} else {
		return dif.c
	}
}

func CountNum(first int64, number int64) int64 {
	if first == number {
		return 1
	}
	return 0
}

func CountNumbers(set ThreeElementInt64Set, distance int64, number int64) (counter int64) {
	counter += CountNum(set.a, number + distance)
	counter += CountNum(set.b, number + distance)
	counter += CountNum(set.c, number + distance)
	counter += CountNum(set.a, number - distance)
	counter += CountNum(set.b, number - distance)
	counter += CountNum(set.c, number - distance)
	return
}

func ElByDistance(first int64, distance int64, number int64) int64 {
	if first - distance == number || first + distance == number {
		return first
	}
	return 0
}

func SumElementsByDistance(set ThreeElementInt64Set, distance int64, number int64) (sum int64) {
	sum += ElByDistance(set.a, distance, number)
	sum += ElByDistance(set.b, distance, number)
	sum += ElByDistance(set.c, distance, number)
	return
}

func main() {
	var (
		abc ThreeElementInt64Set
		x   int64
	)

	_, _ = fmt.Print("Введите множество ABC: ")
	_, _ = fmt.Scan(&abc.a, &abc.b, &abc.c)

	_, _ = fmt.Print("Введите x: ")
	_, _ = fmt.Scan(&x)

	var (
		dif      = CountDifference(abc, x)
		minDIf   = FindNum(abc, dif, func(first int64, second int64) bool { return first <= second })
		maxDif   = FindNum(abc, dif, func(first int64, second int64) bool { return first >= second })
		countMin = CountNumbers(abc, minDIf, x)
		countMax = CountNumbers(abc, maxDif, x)
		result   int64
	)

	if countMin > countMax {
		result = SumElementsByDistance(abc, minDIf, x)
	} else {
		result = SumElementsByDistance(abc, maxDif, x)
	}

	fmt.Println(result)
}