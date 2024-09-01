package dsa

import (
	"errors"
	"fmt"
	"math"
	"time"
)

type Stack struct {
	// Write a program to implement a stack using an array.
	data []int
}

func (s *Stack) Push(item int) {
	s.data = append(s.data, item)
}

func (s *Stack) Pop() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("Stack is empty")
	}

	item := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]

	return item, nil
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func FindMinMaxInArray(a []int) (int, int) {

	if len(a) == 0 {
		return 0, 0
	}

	min := a[0]
	max := a[0]

	for _, v := range a {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return min, max
}

func FibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}

	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

func FactorialRecursive(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * FactorialRecursive(n-1)
}

func sendMessageToChannel(ch chan<- string, msg string) {
	time.Sleep(5 * time.Second)
	ch <- msg
}

func CallGoRoutine(sendTextMsg string) {
	ch := make(chan string)

	go sendMessageToChannel(ch, sendTextMsg)

	msg := <-ch

	fmt.Println(msg)
}
