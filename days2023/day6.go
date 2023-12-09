package days2023

import (
	"fmt"
	"math"
)

func RunDaySix() {

	// RunDaySixPartOne()
	RunDaySixPartTwo()
}

func RunDaySixPartTwo() {

	b := []int64{58819676, 434104122191218}
	// b := []int{60947882, 475213810151650}
	// b := []int{71530, 940200}
	T := b[0]
	D := b[1]
	discriminant := T*T + 2*T + 1 - 4*D

	sqrtDiscriminant := math.Sqrt(float64(discriminant))
	fmt.Println(sqrtDiscriminant)
	root1 := (float64(T) - sqrtDiscriminant + 1) / 2
	root2 := (float64(T) + sqrtDiscriminant + 1) / 2
	fmt.Printf("root 1: %f\n", root1)
	fmt.Printf("root 2: %f\n", root2)

	r1Int := int64(math.Ceil(root1))
	rint2 := int64(math.Floor(root2))

	result := rint2 - r1Int
	fmt.Printf("Number of solutions: %d\n", result)
}

func RunDaySixPartOne() {

	a := [][]int{{58, 434}, {81, 1041}, {96, 2219}, {76, 1218}}
	// b := [][]int{{60, 434}, {94, 1041}, {78, 2219}, {82, 1218}}
	result := 1
	for _, race := range a {
		raceLength := race[0]
		distanceToBeat := race[1]
		speed := 0
		count := 0
		for i := 0; i <= raceLength; i++ {
			speed = i
			distanceTraveled := (raceLength - i) * speed
			if distanceTraveled > distanceToBeat {
				count++
			}
		}
		result *= count
	}
	fmt.Println(result)
}
