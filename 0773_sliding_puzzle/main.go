package main

import (
	"fmt"
	"strconv"
	"strings"
)

type State struct {
	Row   int
	Col   int
	Board *[][]int
}

func serialize(board [][]int) string {
	sb := strings.Builder{}
	for _, row := range board {
		for _, val := range row {
			sb.WriteString(strconv.Itoa(val))
		}
	}
	return sb.String()
}

func createNewBoard(row, col, tagetRow, targetCol int, board *[][]int) *[][]int {
	// go built-in copy is just shallow copy
	newBoard := deepCopy(*board)
	newBoard[tagetRow][targetCol], newBoard[row][col] = newBoard[row][col], newBoard[tagetRow][targetCol]
	return &newBoard
}

// deep copy
func deepCopy(src [][]int) [][]int {
	// initalize a slice of slice with same dimension of input
	dst := make([][]int, len(src))
	for i := range dst {
		dst[i] = make([]int, len(src[0]))
	}
	for i, row := range src {
		for j, val := range row {
			dst[i][j] = val
		}
	}
	return dst
}

func slidingPuzzle(board [][]int) int {
	newBoard := deepCopy(board)
	init := State{0, 0, &newBoard}
	for rowIdx, row := range board {
		for colIdx, val := range row {
			if val == 0 {
				init.Row = rowIdx
				init.Col = colIdx
			}
		}
	}
	queue := []State{init}
	steps := 0
	visited := map[string]bool{}
	for len(queue) != 0 {
		newQueue := []State{}
		for _, curr := range queue {
			serialized := serialize(*curr.Board)
			// if hit the taget, return
			if serialized == "123450" {
				return steps
			}
			// if visited, skip
			if _, ok := visited[serialized]; ok {
				continue
			}
			visited[serialized] = true
			var next *[][]int
			// seach each direction
			if curr.Row > 0 {
				next = createNewBoard(curr.Row, curr.Col, curr.Row-1, curr.Col, curr.Board)
				newQueue = append(newQueue, State{curr.Row - 1, curr.Col, next})
			}
			if curr.Row < len(*curr.Board)-1 {
				next = createNewBoard(curr.Row, curr.Col, curr.Row+1, curr.Col, curr.Board)
				newQueue = append(newQueue, State{curr.Row + 1, curr.Col, next})
			}
			if curr.Col > 0 {
				next = createNewBoard(curr.Row, curr.Col, curr.Row, curr.Col-1, curr.Board)
				newQueue = append(newQueue, State{curr.Row, curr.Col - 1, next})
			}
			if curr.Col < len((*curr.Board)[0])-1 {
				next = createNewBoard(curr.Row, curr.Col, curr.Row, curr.Col+1, curr.Board)
				newQueue = append(newQueue, State{curr.Row, curr.Col + 1, next})
			}
		}
		queue = newQueue
		steps++
	}
	if len(queue) == 0 {
		return -1
	}
	return steps
}

func main() {
	board := [][]int{
		{1, 2, 0},
		{4, 5, 3},
	}
	fmt.Println(slidingPuzzle(board))
}
