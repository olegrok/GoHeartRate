package main

import (
	"fmt"
	//"github.com/olegrok/GoHeartRate/webcam"
	"github.com/olegrok/GoHeartRate/pull"
	"time"
)

type WorkerWithData struct {
	data []int
	ch chan<-pull.Worker
}

type WorkerWithHttp struct {
	message []string
}

/*
	todo
	WorkerWithDatabase
	...
	...
 */

func (w WorkerWithData) Work() {
	fmt.Println(w.data)
	w.ch<- WorkerWithHttp{[]string{"Processing finished"}}
}

func (w WorkerWithHttp) Work() {
	fmt.Println(w.message)
}


func main()  {
	const n = 3
	//webcam.Start()
	ch := pull.CreatePull(n)

	tw := WorkerWithData{[]int{1, 2, 3}, ch}

	ch<- tw
	ch<- WorkerWithHttp{[]string{"First", "Second"}}

	/* todo correct completion of the program */
	time.Sleep(3000)
	close(ch)
}
