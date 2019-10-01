package main

import (
	"fmt"
	"math/big"
)

//Fib структура для числа Фибоначчи
type Fib struct {
	Rchan chan int
}

//GenFib Функция вычисления чисел Фибоначчи меньше 2147483647
func GenFib() *Fib {
	c := make(chan int)

	go func() {
		for i, j := 0, 1; ; i, j = i+j, i {
			c <- i
		}
	}()

	return &Fib{Rchan: c}
}

//PrintFib Функция вывода чисел Фибоначчи меньше 2147483647
func PrintFib(count int) {
	F := GenFib()
	for n := 0; n < count; n++ {
		fmt.Println(<-F.Rchan)
	}
}

//GenBigFib Функция вычисления чисел Фибоначчи больше 2147483647
func GenBigFib(count int) {
	fibf := big.NewInt(0)
	fibs := big.NewInt(1)
	for i := 0; i < count; i++ {
		fibf.Add(fibf, fibs)
		fibf, fibs = fibs, fibf

		fmt.Printf("%v\n", fibf)
	}
}

func main() {
	PrintFib(25) //Вызов функции вывода чисел Фибоначчи меньше 2147483647
	//GenBigFib(1000) //Вызов функции вывода чисел Фибоначчи больше 2147483647
}
