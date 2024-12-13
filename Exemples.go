package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"runtime"
	"time"
)

func SeePlatform() {
	fmt.Println("GOOS: ", runtime.GOOS, " | GOARCH: ", runtime.GOARCH)

}

func Loop(quantity int) {
	var quantityExecutions int64
	quantityExecutions = 2
	var t float64
	t = 0

	iterations := 0
	for iterations < quantity {

		fmt.Println("Quantidade de exibições", quantityExecutions)
		fmt.Println()
		t++
		iterations++
		quantityExecutions = int64(math.Pow(2, t))
		time.Sleep(200 * time.Millisecond)
	}
}

func TesteCaracteres() {
	s := "Hello playground"

	for _, v := range s {
		fmt.Printf("%v - %T - %#U - %#X \n", v, v, v, v)
	}

	fmt.Println("")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - %T - %#U - %#X \n", s[i], s[i], s[i], s[i])
	}
}

func MyFirstRequestInGo() {
	resp, err := http.Get("https://api.adviceslip.com/advice")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf(resp.Status)

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}
