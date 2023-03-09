package main

import "fmt"

func main() {

	var (
		l int
		p int
	)

	var (
		nextL int
		nextP int
	)
	fmt.Print("which line = ")
	fmt.Scanln(&l)
	fmt.Print("which pillar = ")
	fmt.Scanln(&p)

	fmt.Print("which line = ")
	fmt.Scanln(&nextL)
	fmt.Print("which pillar = ")
	fmt.Scanln(&nextP)

	nextArr := []int{ nextL, nextP}
	fmt.Println(nextArr)



	a := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 2, 3, 4, 5, 6, 7, 8}, // 1
		{0, 1, 2, 3, 4, 5, 6, 7, 8},  // 2
		{0, 1, 2, 3, 4, 5, 6, 7, 8},  // 3
		{0, 1, 2, 3, 4, 5, 6, 7, 8},  // 4 ..
		{0, 1, 2, 3, 4, 5, 6, 7, 8},  // 5
		{0, 1, 2, 3, 4, 5, 6, 7, 8},  // 6
		{0, 1, 2, 3, 4, 5, 6, 7, 8},  // 7
		{0, 1, 2, 3, 4, 5, 6, 7, 8},  // 8
	}

	
	

	for i := 1; i < len(a); i++ {
		for j := 1; j < len(a[i]); j++ {
			if i == l && j == p {

				// fmt.Printf("line: %v   pillarLeft: %v pillarRight: %v",i -2 ,j - 1,j + 1) // tepa chap on
				if i - 2  > 0 && i - 2 < 9 {
					fmt.Printf("line: %v", i-2)
					if j-1 > 0 && j-1 < 9 {
						fmt.Printf(" pillarLeft: %v", j - 1 )

					}
					if j+1 > 0 && j+1 < 9 {
						fmt.Printf(" pillarLeft: %v", j+1)

					}

				}

				println()
				// fmt.Printf("line: %v   pillarLeft: %v pillarRight: %v",i +2 ,j - 1,j + 1) // pas chap on

				if i+2 > 0 && i+2 < 9 {
					fmt.Printf("line: %v", i+2)
					if j-1 > 0 && j-1 < 9 {
						fmt.Printf(" pillarLeft: %v", j-1)

					}
					if j+1 > 0 && j+1 < 9 {
						fmt.Printf(" pillarLeft: %v", j+1)

					}

				}

				println()

				// fmt.Printf("pillarRight: %v lineTop: %v lineBotom: %v", j +2 , i -1 , i +1) // on tepa pas
				if j+2 > 0 && j+2 < 9 {
					fmt.Printf("pillarRight: %v", j+2)
					if i-1 > 0 && i-1 < 9 {
						fmt.Printf(" lineTop: %v", i-1)

					}
					if i+1 > 0 && i+1 < 9 {
						fmt.Printf(" lineBotom: %v", i+1)

					}

				}

				println( ) 
				// fmt.Printf("pillarLeft: %v lineTop: %v lineBotom: %v", j -2 , i -1 , i +1) // on tepa pas
				if j-2 > 0 && j-2 < 9 {
					fmt.Printf("pillarLeft:  %v", j-2)
					if i-1 > 0 && i-1 < 9 {
						fmt.Printf(" lineTop: %v", i-1)

					}
					if i+1 > 0 && i+1 < 9 {
						fmt.Printf(" lineBotom: %v", i+1)

					}

				}

			}

		}
		// fmt.Println()

	}

	fmt.Println()

}
