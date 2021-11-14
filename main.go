package main

import "fmt"

const LOAD = 2

func main() {

	var floor4, floor7, floor10, temp int

	fmt.Println("Введите количество человек на 4, 7 и 10 этаже:")

	fmt.Scan(&floor4)
	fmt.Scan(&floor7)
	fmt.Scan(&floor10)

	for {	
		fmt.Println("на 10 этаже осталось:", floor10)
		fmt.Println("на 7 этаже осталось:", floor7)
		fmt.Println("на 4 этаже осталось:", floor4)
    fmt.Println()

		temp = LOAD //количество свободных мест в лифте
		if floor10 > 0 && temp > 0 {
			if floor10 >= temp {
				floor10 -= temp
				temp = 0

				continue
			} else {
				temp -= floor10
				floor10 = 0
			}

		}

		if floor7 > 0 && temp > 0 {
			if floor7 >= temp {
				floor7 -= temp
				temp = 0

				continue
			} else {
				temp -= floor7
				floor7 = 0
			}
		}

		if floor4 > 0 && temp > 0 {
			if floor4 > temp {
				floor4 -= temp
				temp = 0

				continue
			} else {
				temp -= floor4
				floor4 = 0
			}
		}
		if floor10 == 0 && floor7 == 0 && floor4 == 0 {
			fmt.Println("все спустились")
			break
		}

	}
}
