package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	var places = make([]int, n)
	var occupiedSpace = m/2 + m%2
	var directionShift int

	if m%2 == 0 {
		directionShift = 1
	} else {
		directionShift = -1
	}

	for personNumber, shiftCounter := 0, 1; personNumber < n; personNumber, shiftCounter = personNumber+1, shiftCounter+1 {

		places[personNumber] = occupiedSpace
		occupiedSpace += directionShift * shiftCounter
		directionShift *= -1

		if shiftCounter == m {
			shiftCounter = 1
			occupiedSpace = m/2 + m%2
		}
	}
	fmt.Println(places)
}
