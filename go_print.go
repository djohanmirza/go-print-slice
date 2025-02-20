package print_slice

import "fmt"

func Print(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}

func Print2D(slice [][]int) {
	for _, row := range slice {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}
}
