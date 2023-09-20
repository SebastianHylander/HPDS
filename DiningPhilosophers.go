/*
To prevent deadlocking we made use of circlular deadlock prevention, which allows the philosophers to wait
their forks without causing a circular wait condition aka deadlock
We found this method in the following article (section 10.2.2):
https://lass.cs.umass.edu/~shenoy/courses/fall13/lectures/Lec10_notes.pdf
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

const nPhil = 5
const nForks = 5

var wg sync.WaitGroup

func main() {
	forks := make([]chan bool, nForks)

	for i := 0; i < nForks; i++ {
		forks[i] = make(chan bool)
		go func(i int) {
			forks[i] <- true
		}(i)
	}

	for i := 1; i <= nPhil; i++ {
		wg.Add(1) //increments the amount of nonfinished go-routines
		go func(i int) {
			Fork(forks[i%nForks])
		}(i)
		go func(i int) {
			Philosophers(i, forks[(i)%nForks], forks[(i+1)%nForks])
		}(i)
	}
	wg.Wait()
}

func Philosophers(id int, leftFork chan bool, rightFork chan bool) {
	var timesEaten = 0

	for { //All should eat at least 3 times
		fmt.Printf("Philosopher %d is thinking\n", id)
		time.Sleep(time.Millisecond * 500) //Timeout to give the philosopher a chance to think

		//request forks, send message
		//fmt.Printf("Philosopher %d request forks\n", id)
		<-leftFork
		<-rightFork

		//release forks
		timesEaten += 1
		fmt.Printf("Philosopher %d is eating for the %d. time; releasing forks\n", id, timesEaten)
		time.Sleep(time.Millisecond * 500) //Timeout to give the philosopher a chance to eat
		if timesEaten == 3 {
			fmt.Printf("Philosopher %d is Done\n", id)
			wg.Done() //decrements the amount of non-finished go-routines

		}
		if id%2 == 0 {
			leftFork <- true
			rightFork <- true
		} else {
			rightFork <- true
			leftFork <- true
		}

	}
}

func Fork(place chan bool) {

	for {
		var fork = <-place
		place <- fork
	}

}
