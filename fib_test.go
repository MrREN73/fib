package main

import "testing"

func BenchmarkRun(b *testing.B) {

	PrintFib(25)    //Вызов функции вывода чисел Фибоначчи меньше 2147483647
	GenBigFib(1000) //Вызов функции вывода чисел Фибоначчи больше 2147483647

}
