package mdl_two

import "fmt"

func opMin(values ...uint32) uint32 {
	min := values[0]
	for i, value := range values {
		if i == 0 || value < min {
			min = value
		}
	}
	return min
}

func opAvg(values ...uint32) uint32 {
	var sum uint32 = 0
	for _, value := range values {
		sum += value
	}
	return sum / uint32(len(values))
}

func op(opName string) func(values ...uint32) uint32 {
	switch opName {
	case "min":
		return opMin
	case "avg":
		return opAvg
	}
	return nil
}

func Exercise04() {
	fmt.Println("====== EXERCISE 04 ======")

	oper := op("min")
	fmt.Printf("minimum value from list %d\n", oper(1, 2, 4, 5))

	oper = op("avg")
	fmt.Printf("average value from list %d\n", oper(1, 2, 4, 5))
}
