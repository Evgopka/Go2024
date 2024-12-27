package main

import (
	"fmt"
	"strconv"
)

func main() {
	var IPlist1 [4]byte
	IPlist1[0] = 127
	IPlist1[1] = 0
	IPlist1[2] = 0
	IPlist1[3] = 1
	IPlist2 := [4]byte{127, 0, 0, 1}
	printList(IPlist1)
	printList(IPlist2)
	fmt.Println(listEven(0, 11))
	fmt.Println(listEven(1, 0))
	fmt.Println(listEven(5, 100))
	countSymbols("Boing, boing, boom!") // Отсылка на Кли из Genshin Impact)
	result := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(result); i++ {
		fmt.Println("Число до:", strconv.Itoa(result[i]))
	}
	result = Map(result, DobleOrNothing)
	for i := 0; i < len(result); i++ {
		fmt.Println("Число после:", strconv.Itoa(result[i]))
	}

}

func printList(IPlist [4]byte) { //1.1
	var IPString string
	var temp string
	for i := 0; i < 4; i++ {
		temp = strconv.Itoa(int(IPlist[i]))
		if i < 3 {
			IPString = IPString + temp + "."
		} else {
			IPString = IPString + temp
		}
	}
	fmt.Println(IPString)
}

func countSymbols(text string) { //2.1
	symbolMap := map[byte]int{}
	glued := text
	for i := 0; i < (len(glued)); i++ {
		val, ok := symbolMap[glued[i]]
		if ok {
			symbolMap[glued[i]] = val + 1
		} else {
			symbolMap[glued[i]] = 1
		}
	}
	for key, value := range symbolMap {
		fmt.Println(string(key) + " : " + strconv.Itoa(value))
	}
}

func listEven(A int, B int) ([]int, error) { // 1.2
	if A%2 == 1 {
		A = A + 1
	}
	glue := []int{A}
	A = A + 2
	if B < A {
		return []int{}, fmt.Errorf("Первое число должно быть больше второго ;)")
	}
	i := 0
	for A <= B {
		glue = append(glue, A)
		A = A + 2
		i = i + 1
	}
	return glue, nil
}
func Map(n []int, f func(int) int) []int { //4.1
	result := make([]int, len(n))
	i := 0
	for i < len(n) {
		result[i] = f(n[i])
		i++
	}
	return result
}

func DobleOrNothing(x int) int { return x * x } //4.2
