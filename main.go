package main

import (
	"time"
	"fmt"
	"log"
	"net/http"

	"./handler"
	"./db"
)

func main() {

	var postgres *db.Postgres
    var err error
    for i := 0; i < 10; i++ {
        time.Sleep(3 * time.Second)
        postgres, err = db.ConnectPostgres()
    }
    if err != nil {
        panic(err)
    } else if postgres == nil {
        panic("postgres is nil")
    }

	mux := handler.SetUpRouting(postgres)
	
    fmt.Println("Connection to postgres successful")

    // fmt.Println("http://localhost:8080")
	// log.Fatal(http.ListenAndServe(":8080", mux))
	


	// mux := handler.SetUpRouting()
	fmt.Println("Server running at: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
