package main

import (
    "fmt"
    "log"
    "net/http"
)

func sampleApp(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "It's a sample App!")
    fmt.Println("Endpoint Hit: sampleApp")
}

func handleRequests() {
    http.HandleFunc("/sampleapp", sampleApp)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
    fmt.Println("Started")
    handleRequests()

}
