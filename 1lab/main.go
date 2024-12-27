package main

import (
	"fmt"
)

func main() {
	hello("Женя")
	err := printEven(1, 10)
	if err != nil {
		fmt.Println("error", err)
	}
	err = printEven(10, 1)
	if err != nil {
		fmt.Println("error", err)
	}
	c, err := apply(1, 1, "+")
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println(c)
	}
	d, err := apply(1, 0, "/")
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println(d)
	}
	_, err = apply(1, 1, "a")
	if err != nil {
		fmt.Println("error", err)
	}
}

func hello(name string) {
	fmt.Printf("Привет, %s!\n", name) // Используем Printf для %s, чтобы не было пробела в выводе после функции name
}

func printEven(x, y int) error {
	if x > y {
		return fmt.Errorf("x больше y")
	}
	for j := x; j < y+1; j++ {
		if j%2 == 0 {
			fmt.Print(j, " ")
		}
	}
	fmt.Println("")
	return nil // Если все успешно, возвращаем 1
}

func apply(a int, b int, o string) (int, error) {
	if o != "+" && o != "-" && o != "*" && o != "/" {
		return 0, fmt.Errorf("нет действия")
	} else if o == "+" {
		return a + b, nil
	} else if o == "-" {
		return a - b, nil
	} else if o == "*" {
		return a * b, nil
	} else if o == "/" && b != 0 {
		return a - b, nil
	} else {
		return 0, fmt.Errorf("Ошибка, нельзя делить на ноль")
	}
}
