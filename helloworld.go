package main

import (
	"context"
	"fmt"
	"hello/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	// "reflect"
	// "strconv"
)

var printFunc = fmt.Println

/*
func basics() {
	// cli inputs
	fmt.Println("Hello, What is your name?")
	// reader := bufio.NewReader(os.Stdin)
	// name, err := reader.ReadString('\n')
	// if err == nil {
	// 	fmt.Println("Hello", name)
	// } else {
	// 	log.Fatal(err)
	// }

	// varibale declaration
	var firstName string = "John"
	var n1, n2 = 1.1, 1.3
	var lastName = "Doe"
	someVar := 0
	someVar = 1

	// types of varibales
	printFunc(reflect.TypeOf(firstName))
	printFunc(reflect.TypeOf(n1))
	printFunc(reflect.TypeOf(n2))
	printFunc(reflect.TypeOf(lastName))
	printFunc(reflect.TypeOf(someVar))

	//typecasting
	printFunc(int(n1))

	num1 := "5000000"
	convertedNum1, err := strconv.Atoi(num1)
	printFunc(convertedNum1, err, reflect.TypeOf(convertedNum1))
}
*/

func main() {

	// basics()

	//HTTP Routes
	// http.HandleFunc("/health", func(http.ResponseWriter, *http.Request) {
	// 	log.Println("Healthy and Running!")
	// })

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	hl := handlers.NewHealth(l)
	gb := handlers.NewGoodbye(l)
	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/hello", hh)
	sm.Handle("/health", hl)
	sm.Handle("/bye", gb)
	sm.Handle("/products", ph)

	//HTTP Server Simple
	// http.ListenAndServe(":8000", sm)

	s := &http.Server{
		Addr:         ":8000",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful termination!", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
