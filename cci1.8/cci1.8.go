/*
          ~Caleb McCarthy~  Jun/10/2018
Cracking the code interview problem 1.8
Zero Matrix: Write an algorithm such that if an element in an MxN matrix is 0, its entire row and
column are set to O.


       \Test DataSet/  \Transformed DataSet/
        _____________     ____________
	|1 ,2 ,5 ,12|     |1 ,2 ,0 ,12|
	-------------     -------------
        |4 ,5 ,0 ,13|     |0 ,0 ,0 ,0 |
	------------- ->  -------------
        |7 ,8 ,9 ,14|     |7 ,8 ,0 ,14|
	-------------     -------------
        |10,11,16,15|     |10,11,0 ,15|
	~~~~~~~~~~~~~     ~~~~~~~~~~~~~

	Runtime Speed complexity:O(N)
	Space Complexity        :O(N)

*/

package main

import "fmt"

func main() {
	var matrix = [][]int{}
	var horizontal = make(map[int]int)
	var vertical = make(map[int]int)

	matrix = append(matrix, []int{1, 2, 5, 12})
	matrix = append(matrix, []int{4, 5, 0, 13})
	matrix = append(matrix, []int{7, 8, 9, 14})
	matrix = append(matrix, []int{10, 11, 16, 15})

	fmt.Println(matrix)

	for i := 0; i < len(matrix); i++ {

		fmt.Println("")
		for j := 0; j < len(matrix[i]); j++ {
			//fmt.Printf("%d,", matrix[i][j])
			if matrix[i][j] == 0 {
				horizontal[j] = j
				vertical[i] = i
			}

		}

	}
	for k := range horizontal {
		zeroOutVertical(matrix, k)
	}

	for k := range vertical {
		zeroOutHorizontal(matrix, k)

	}

	fmt.Println(matrix)

}

func zeroOutVertical(matrix [][]int, vertical int) {
	var length = len(matrix)
	for i := 0; i < length; i++ {
		matrix[i][vertical] = 0

	}

}

func zeroOutHorizontal(matrix [][]int, horizontal int) {
	var width = len(matrix[horizontal])
	for i := 0; i < width; i++ {
		matrix[horizontal][i] = 0

	}

}
