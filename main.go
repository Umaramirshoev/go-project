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
