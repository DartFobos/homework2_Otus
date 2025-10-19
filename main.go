package main

import (
	funcForBoard "Chessboard/func"
	"fmt"
)

func main() {
	w := "  "
	b := "# "
	c := ""
	f := ""
	var namber int
	fmt.Print("Введите желаемый размер доски (число от 2 до 88): ")
	fmt.Scan(&namber)
	if namber > 1 && namber < 89 {
		if namber%2 != 0 {
			f = funcForBoard.OddNamber(namber, c, w, b, f)
		} else {

			f = funcForBoard.EvenNamber(namber, c, w, b, f)
		}
		fmt.Print(f)
	} else {
		fmt.Println("неверно введено число")
	}
}
