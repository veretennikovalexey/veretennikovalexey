// создание переменной, ввод данных с клавиатуры, условный оператор if

package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a > 0 {
		fmt.Print("Число положительное")
	} else if a < 0 {
		fmt.Print("Число отрицательно")
	} else {
		fmt.Print("Ноль")
	}
}
