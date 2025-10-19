package funcForBoard

func ReverseString(s string) string { //функция отзеркаливает строку
	board := []rune(s)             // Преобразовываем строку в руны
	left, right := 0, len(board)-1 // Определяем границы слева и справа
	for left < right {             // Пока левая граница меньше правой
		board[left], board[right] = board[right], board[left] // Меняем местами элементы
		left++
		right--
	}
	return string(board)
}

func OddNamber(namber int, c, w, b, f string) string {
	for i := 0; i < namber/2; i++ {
		c = c + w + b
	}
	d := " " + c + w + "\n"               //лишний пробел вначале для более красивого визуального выравнивания при выводе
	c = ReverseString(c) + " " + b + "\n" //лишний пробел вначале для более красивого визуального выравнивания при выводе
	c = c + d
	for i := 0; i < (namber-1)/2; i++ {
		f += c
	}
	f = d + f
	return f
}

func EvenNamber(namber int, c, w, b, f string) string {
	for i := 0; i < namber/2; i++ {
		c = c + w + b
	}
	d := ReverseString(c)
	c = " " + c + "\n" + d + "\n" //лишний пробел вначале для более красивого визуального выравнивания при выводе
	for i := 0; i < namber/2; i++ {
		f += c
	}
	return f
}
