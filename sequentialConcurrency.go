package main

import (
    "fmt"
    "sync"
    "time"
)

var wg sync.WaitGroup

func workerOne(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    defer time.Sleep(time.Second)
    //defer wg.Done()
    ch <- "Worker 1 done"
}

func workerTwo(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)

    //defer wg.Done()
    defer time.Sleep(time.Second)
    ch <- "Worker 2 done"
}

func workerThree(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    //defer wg.Done()
    defer time.Sleep(time.Second)
    ch <- "Worker 3 done"
}

func workerFour(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)

    //defer wg.Done()
    defer time.Sleep(time.Second)
    ch <- "Worker 4 done"
}

func workerFive(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)

    //defer wg.Done()
    defer time.Sleep(time.Second)
    ch <- "Worker 5 done"
}

func workerSix(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)

    //defer wg.Done()
    defer time.Sleep(time.Second)
     ch <- "Worker 6 done"
}

func workerSeven(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)

    //defer wg.Done()
    defer time.Sleep(time.Second)
    ch <- "Worker 7 done"
}

func workerEight(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)

    //defer wg.Done()
    defer time.Sleep(time.Second)
    ch <- "Worker 8 done"
}

func workerNine(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)

    //defer wg.Done()
    defer time.Sleep(time.Second)
    ch <- "Worker 9 done"
}

func workerTen(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    //defer wg.Done()
    defer time.Sleep(time.Second)
    ch <- "Worker 10 done"
}

func main() {

    //I would like to trial running 10 routines at a time

    firstchan := make(chan string)
    secondchan := make(chan string)
    thirdchan := make(chan string)
    fourthchan := make(chan string)
    fifthchan := make(chan string)
    sixthchan := make(chan string)
    seventhchan := make(chan string)
    eightchan := make(chan string)
    ninthchan := make(chan string)
    tenthchan := make(chan string)

    //so we are going to wait for 10 to finish then we'll start the next heat

    for i := 1; i <= 100; i++ {
	wg.Add(1)
	j:=i%10
	
	switch j {
	case 1:
		go workerOne(i,firstchan)
		fmt.Println(<-firstchan)
	case 2:
		go workerTwo(i,secondchan)
		fmt.Println(<-secondchan)
	case 3:
		go workerThree(i,thirdchan)
		fmt.Println(<-thirdchan)
	case 4:
		go workerFour(i,fourthchan)
		fmt.Println(<-fourthchan)
	case 5:
		go workerFive(i,fifthchan)
		fmt.Println(<-fifthchan)
	case 6:
		go workerSix(i,sixthchan)
		fmt.Println(<-sixthchan)
	case 7:
		go workerSeven(i,seventhchan)
		fmt.Println(<-seventhchan)
	case 8:
		go workerEight(i,eightchan)
		fmt.Println(<-eightchan)
	case 9:
		go workerNine(i, ninthchan)
		fmt.Println(<-ninthchan)
	case 0:
		go workerTen(i, tenthchan)
		fmt.Println(<-tenthchan)
	}
	wg.Done()
    }
    wg.Wait()
    fmt.Println("Main stopped")
}
