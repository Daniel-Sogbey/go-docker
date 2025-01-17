package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/", handlePage)

	addr := os.Getenv("PORT")
	srv := http.Server{
		Addr:         addr,
		Handler:      m,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("server listening on ", addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func handlePage(w http.ResponseWriter, _ *http.Request) {
	body := `
	<html>
		<body>
			<p> Hello from Go Server </p>
		</body>
	</html>
	`
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	w.Write([]byte(body))
}
