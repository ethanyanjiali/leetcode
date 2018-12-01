package main

import "fmt"


func search(board [][]byte, row, col, pos int, word string) bool {
	if pos >= len(word) {
		return true
	}
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) {
		return false
	}
	if board[row][col] != word[pos] {
		return false
	}
	board[row][col] ^= 255
	res := search(board, row - 1, col, pos + 1, word) || 
		search(board, row + 1, col, pos + 1, word) ||
		search(board, row, col - 1, pos + 1, word) ||
		search(board, row, col + 1, pos + 1, word)
	board[row][col] ^= 255
	if res == true {
		return true
	}
	return res
}

func exist(board [][]byte, word string) bool {
	res := false
    for i := range board {
    	for j := range board[0] {
    		res = search(board, i, j, 0, word)
    		if res == true {
    			return true
    		}
    	}
    }
    return false
}

func main() {
	board := [][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}
	word := "ABCCED"
	fmt.Println(exist(board, word))
	board = [][]byte{{'b','b'},{'a','b'},{'b','a'}}
	word = "a"
	fmt.Println(exist(board, word))
	board = [][]byte{{'a', 'a'}}
	word = "aaa"
	fmt.Println(exist(board, word))
	board = [][]byte{{'a'}}
	word = "a"
	fmt.Println(exist(board, word))
}