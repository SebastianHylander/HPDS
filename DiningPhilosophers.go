package main

import (
	"fmt"
	"time"
)

const nPhil = 5

func main() {
	//Udlever forks til philo (async)
	//Hvis man har 1 fork, print "thinking", og læg den ned igen.
	//Udlever hele tiden indtil en får 2 forks.
	//Hvis man har 2 forks, print "eating", og læg begge ned igen.

	leftFork := make(chan bool)
	rightFork := make(chan bool)

	for i := 1; i <= nPhil; i++ {
		go Philosophers(i, leftFork, rightFork)
		//recieve message:
		/* msg := <-leftFork
		msg1 := <-rightFork
		fmt.Println(i, msg, msg1) */
	}
}

func Philosophers(id int, leftFork chan bool, rightFork chan bool) {
	for i := 0; i < 3; i++ { //All should eat at least 3 times
		fmt.Printf("Philosopher %d is thinking\n", id)
		time.Sleep(time.Millisecond * 500)

		//request forks, send message
		fmt.Printf("Philosopher %d request forks\n", id)
		<-leftFork
		<-rightFork

		//release forks
		//OBS: Der skal være en betingelse, at når en af dem er true, og den anden er false, så realeaser man (deadlock)
		fmt.Printf("Philosopher %d has eaten; releasing forks\n", id)
		leftFork <- true
		rightFork <- true
	}
}
