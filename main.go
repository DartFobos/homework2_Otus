package main

import "fmt"

func reverseString(s string) string { //функция отзеркаливает строку
	chars := []rune(s)             // Преобразовываем строку в руны (работает даже с юникодом)
	left, right := 0, len(chars)-1 // Определяем границы слева и справа
	for left < right {             // Пока левая граница меньше правой
		chars[left], chars[right] = chars[right], chars[left] // Меняем местами элементы
		left++
		right--
	}
	return string(chars) // Возвращаем преобразованную строку
}
func main() {
	w := "  "
	b := "# "
	c := ""
	d := ""
	f := ""
	var namber int
	fmt.Print("Введите желаемый размер доски (число): ")
	fmt.Scan(&namber)
	if namber > 1 {
		if namber%2 != 0 {
			for i := 0; i < namber/2; i++ {
				c = c + w + b
			}
			d = " " + c + w + "\n"                //лишний пробел вначале для более красивого визуального выравнивания при выводе
			c = reverseString(c) + " " + b + "\n" //лишний пробел вначале для более красивого визуального выравнивания при выводе
			c = c + d
			for i := 0; i < (namber-1)/2; i++ {
				f += c
			}
			f = d + f
		} else {
			for i := 0; i < namber/2; i++ {
				c = c + w + b
			}
			d = reverseString(c)
			c = " " + c + "\n" + d + "\n" //лишний пробел вначале для более красивого визуального выравнивания при выводе
			for i := 0; i < namber/2; i++ {
				f += c
			}
		}
		fmt.Print(f)
	} else {
		fmt.Println("неверно введено число")
	}
}
