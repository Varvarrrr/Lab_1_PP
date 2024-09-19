# Lab_PP
Laboratornaya_1

1. Написать программу, которая выводит текущее время и дату.
func timeNow() {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println("Current time and date: ", formattedTime)
}

2. Создать переменные различных типов (int, float64, string, bool) и вывести их на экран.
func variableOutput() {
	var age int = 20
	var happinessLevel float64 = 9.99
	var name string = "Petryakova Varvara"
	var education bool = true

	fmt.Println(age, "|", happinessLevel, "|", name, "|", education)
}

3. Использовать краткую форму объявления переменных для создания и вывода переменных.
func shortFormVariableOutput() {
	num := 5
	num_float := 10.765
	line := "Hello, World!!!"
	boolean := false

	fmt.Println(num, "|", num_float, "|", line, "|", boolean)
}

4. Написать программу для выполнения арифметических операций с двумя целыми числами и выводом результатов.
func arithmeticOperations() {
	var a, b int

	fmt.Println("Введите число a: ")
	fmt.Scan(&a)
	fmt.Println("Введите число b: ")
	fmt.Scan(&b)

	fmt.Println("Результат умножения:", a*b, "|", "Результат деления:", a/b, "|", "Результат вычитания:", a-b, "|", "Результат сложения:", a+b)
}

5. Реализовать функцию для вычисления суммы и разности двух чисел с плавающей запятой.
func calculatingSumAndDifferenceFloatingNumbers() {
	var a, b float64

	fmt.Println("Введите число a: ")
	fmt.Scan(&a)
	fmt.Println("Введите число b: ")
	fmt.Scan(&b)

	fmt.Println("Результат вычитания:", a-b, "|", "Результат сложения:", a+b)
}

6. Написать программу, которая вычисляет среднее значение трех чисел.
func averageNumbers() {
	var a, b, c float64

	fmt.Println("Введите число a: ")
	fmt.Scan(&a)
	fmt.Println("Введите число b: ")
	fmt.Scan(&b)
	fmt.Println("Введите число c: ")
	fmt.Scan(&c)

	average := (a + b + c) / 3

	fmt.Printf("Среднее значение: %.2f", average)
}
