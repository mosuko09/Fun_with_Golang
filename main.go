package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var name string

	for {
		fmt.Print("Введите ваше имя: ")
		_, err := fmt.Scanln(&name)
		if err != nil {
			fmt.Println("Ошибка ввода. Попробуйте ещё раз.")
			continue
		}

		name = strings.TrimSpace(name)
		if name == "" {
			fmt.Println("Имя не может быть пустым. Пожалуйста, введите имя.")
			continue
		}

		hasCyr := false
		hasLat := false
		for _, r := range name {
			if unicode.Is(unicode.Cyrillic, r) {
				hasCyr = true
			}
			if unicode.Is(unicode.Latin, r) {
				hasLat = true
			}
		}

		if hasCyr && hasLat {
			fmt.Println("Похоже, вы могли перепутать раскладку клавиатуры: в имени смешаны кириллица и латиница. Хотите ввести имя заново? (y/n): ")
			var choice string
			fmt.Scan(&choice)
			if strings.ToLower(choice) == "y" {
				continue
			}
		}
		break
	}

	fmt.Printf("Привет, %s! Давай посчитаем.\n", name)

	var num1, num2 float64
	var op string

	for {
		fmt.Print("Введите первое число: ")
		if _, err := fmt.Scan(&num1); err == nil {
			break
		}
		fmt.Println("Некорректный ввод числа. Попробуйте ещё раз.")
	}

	for {
		fmt.Print("Введите операцию (+, -, *, /): ")
		if _, err := fmt.Scan(&op); err == nil && (op == "+" || op == "-" || op == "*" || op == "/") {
			break
		}
		fmt.Println("Некорректная операция. Допустимы только: +, -, *, /")
	}

	for {
		fmt.Print("Введите второе число: ")
		if _, err := fmt.Scan(&num2); err == nil {
			break
		}
		fmt.Println("Некорректный ввод числа. Попробуйте ещё раз.")
	}

	var result float64
	var hasError bool

	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Ошибка: деление на ноль!")
			hasError = true
		} else {
			result = num1 / num2
		}
	default:
		fmt.Println("Ошибка: неизвестная операция!")
		hasError = true
	}

	if !hasError {
		fmt.Printf("Результат: %.2f %s %.2f = %.2f\n", num1, op, num2, result)
	}
}
