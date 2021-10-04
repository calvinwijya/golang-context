package belajarcontext_test

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	fmt.Println(contextF.Value("f")) //dapat
	fmt.Println(contextF.Value("c")) //milik parent
	fmt.Println(contextF.Value("b")) //tidak bisa, different parent
	fmt.Println(contextA.Value("b")) //tidak bisa mengabil data child

	fmt.Println(contextA.Value("b"))

	//context parent tidak bisa mengakses value ke child ny
}

func CreatCounter(ctx context.Context) chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination

}

func TestContextWithCancel(t *testing.T) {

	fmt.Println("total goroutine: ", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancle := context.WithTimeout(parent, 4*time.Second)

	destination := CreatCounter(ctx)
	for n := range destination {
		fmt.Println("counter: ", n)
		if n == 10 {
			break
		}
	}

	cancle()
	time.Sleep(1 * time.Second)
	fmt.Println("total goroutine: ", runtime.NumGoroutine())
}

func TestContextWithTimeOut(t *testing.T) {

	fmt.Println("total goroutine: ", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancle := context.WithTimeout(parent, 5*time.Second)
	defer cancle()

	destination := CreatCounter(ctx)
	for n := range destination {
		fmt.Println("counter: ", n)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("total goroutine: ", runtime.NumGoroutine())
}
