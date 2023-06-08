package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var nonStop int64

func main() {
	go func() {
		last := time.Now()
		for range time.NewTicker(time.Millisecond).C {
			now := time.Now()
			// t2. Println will write to os.Stdout.
			// t7. When fluentd is down, this goroutine will be stuck here.
			fmt.Println("Hello:", now)
			if now.Sub(last) >= time.Second {
				nonStop++
				last = now
			}
		}
	}()

	http.HandleFunc("/stdout", stdout)
	http.HandleFunc("/stderr", stderr)
	panic(http.ListenAndServe(":8090", nil))
}

func stdout(w http.ResponseWriter, req *http.Request) {
	// t3. Calling /stdout API is OK.
	// t8. Calling /stdout API is *not* OK after t4.
	_, _ = fmt.Fprintf(os.Stdout, "request writing stdout")
	_, _ = fmt.Fprintf(w, "nonStop: %v\n", nonStop)
	w.WriteHeader(http.StatusOK)
}

func stderr(w http.ResponseWriter, req *http.Request) {
	// t4. Calling /stderr API is OK.
	// t9. Calling /stderr API is *still* OK after t4.
	_, _ = fmt.Fprintf(os.Stderr, "request writing stderr")
	_, _ = fmt.Fprintf(w, "nonStop: %v\n", nonStop)
	w.WriteHeader(http.StatusOK)
}
