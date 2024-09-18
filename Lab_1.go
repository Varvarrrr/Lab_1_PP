package main

import (
	"fmt"
	"time"
)

func main() {
	var choice int
	fmt.Println("Выберите задание (от 1 до 6):") 
	fmt.Scan(&choice)

	switch choice {
	case 1:
		timeNow() // Написать программу, которая выводит текущее время и дату.
	case 2:
		fmt.Println("Задание ещё не готово!!!")
	case 3:
		fmt.Println("Задание ещё не готово!!!")
	case 4:
		fmt.Println("Задание ещё не готово!!!")
	case 5:
		fmt.Println("Задание ещё не готово!!!")
	case 6:
		fmt.Println("Задание ещё не готово!!!")
	default:
		fmt.Println("Выберете число от 1 до 6!!!")
	}

}

func timeNow() {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println("Current time and date: ", formattedTime)
}
