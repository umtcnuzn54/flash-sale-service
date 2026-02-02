package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	totalRequests := 1000
	var wg sync.WaitGroup

	fmt.Println(" SALDIRI BAŞLIYOR! 1000 kişi aynı anda butona basıyor")
	start := time.Now()

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		go func() { 
			defer wg.Done()
			http.Get("http://localhost:3000/buy")
		}()
	}

	wg.Wait() 
	fmt.Printf("🏁 Saldırı bitti! Geçen süre: %v\n", time.Since(start))
	fmt.Println("Lütfen tarayıcıdan http://localhost:3000/status adresine girip stoğu kontrol et")
}
