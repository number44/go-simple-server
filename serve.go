package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)
func main() {
	
	if len(os.Args) < 2 {
        log.Fatal("Add directory to comand f.e. server ./static")
    }
	fmt.Printf("Directory served %v \n" , os.Args[1])
    directory := os.Args[1]
    fs := http.FileServer(http.Dir(directory))
    http.Handle("/", fs)
	println("Static serve on http://localhost:8080/")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
