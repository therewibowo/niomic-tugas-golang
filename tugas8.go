package main

import "fmt"
import "runtime"
import "math/rand"
import "time"

func main()  {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)

	var message = make(chan int)
	
	go kirimdata(message)
	terimdata(message)
}

func kirimdata(ch chan <- int)  {
	for i := 0; true; i++ {
		ch <-i
		time.Sleep(time.Duration(rand.Int()%10+1)* time.Second)

	}
}

func terimdata(ch <- chan int)  {
	loop:
	for{
		select{
		case data:= <-ch:
			fmt.Println("haloo teman-teman, apa kabar ", data, "\n")
		case <-time.After(time.Second*5):
			fmt.Println("timeout , tidak ada aktivitas dalam  detik")
			break loop	
		}
	}
}
