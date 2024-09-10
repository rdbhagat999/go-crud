package dsa

import (
	"fmt"
	"sync"
	"time"
)

func subTask(wg *sync.WaitGroup, index int, msg string) {
	defer wg.Done()

	fmt.Printf("\n*** Subtask %d START\n", index)
	fmt.Printf("%s \n", msg)
	fmt.Printf("*** Subtask %d END\n", index)
	println("")
}

func Chimken_Biryani_With_Go_Routines() {

	var wg sync.WaitGroup

	println("")
	println("CHIMKEN BIRYANI PROCESS STARTED")
	println("")

	tasks := []string{
		"SON cleaning dinner plates",
		"MOM choping vegetables",
		"DAD frying chimken",
		"SISTER preparing marination masala",
	}

	for i := 0; i < len(tasks); i++ {
		wg.Add(1)
		go subTask(&wg, i+1, tasks[i])
	}

	wg.Wait()

	println("Mix all ingredients")
	println("Let it cook for 3 hours")
	time.Sleep(3 * time.Second)
	println("Chimken Biryani is ready!")
	println("")
	println("CHIMKEN BIRYANI PROCESS END")
	println("")
}
