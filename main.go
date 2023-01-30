package main

import "fmt"

func main() {

	//1. Написать функцию, принимающую аргумент типа int. Выводить на экран "Да", если аргумент больше 10 и меньше 20, "Нет" если аргумент меньше 10, "Наверное" если аргумент больше 20.
	var input float64
	fmt.Println("Введите число")
	fmt.Scanf("%f", &input)
	if input > 10 && input < 20 {
		fmt.Println("ДА")
	} else if input < 10 {
		fmt.Println("НЕТ")
	} else if input > 20 {
		fmt.Println("Наверное")
	}

}
