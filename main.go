package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	NewHelloHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server on listening on port 8081")
	server.ListenAndServe()

}

// package main

// import "fmt"

// func sumPart(arr []int, ch chan int) {
// 	sum := 0
// 	for _, num := range arr {
// 		sum += num
// 	}
// 	ch <- sum
// }

// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
// 	numGoroutines := 3
// 	ch := make(chan int, numGoroutines)

// 	partize := len(arr) / numGoroutines

// 	for i := 0; i < numGoroutines; i++ {
// 		start := i * partize
// 		end := start + partize
// 		go sumPart(arr[start:end], ch)
// 	}
// 	totalsum := 0
// 	for i := 0; i < numGoroutines; i++ {
// 		totalsum += <-ch
// 	}
// 	fmt.Println("Total sum: ", totalsum)
// }

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"sync"
// )

// func main() {
// 	code := make(chan int)
// 	var wg sync.WaitGroup
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func() {
// 			getHttp(code)
// 			wg.Done()
// 		}()
// 	}
// 	go func() {
// 		wg.Wait()
// 		close(code)
// 	}()
// 	for res := range code {
// 		fmt.Printf("Status code:%d\n", res)

// 	}

// }

// func getHttp(codeCh chan int) {
// 	resp, err := http.Get("https://google.com")
// 	if err != nil {
// 		fmt.Println("Ошибка: ", err)
// 		return

// 	}

// 	codeCh <- resp.StatusCode

// }
