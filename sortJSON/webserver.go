package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	// "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	arr := []int{8, 6, -2,  2,  4,  17,  256, 1024, -17, -19}
	jdata, err := json.Marshal(arr)
	if err != nil {
		fmt.Println("error:", err)
	}
    fmt.Fprintf(w, "%s", jdata)
}


func main() {
	var addr string
	var port string
	var a string
	var b string

	fmt.Scanf("%s\n", &addr)
	fmt.Scanf("%s\n", &port)
	fmt.Scanf("%s\n", &a)
	fmt.Scanf("%s\n", &b)

	addr = addr[7:] + ":" + port
	fmt.Println(addr, port, a, b)

	fmt.Printf("Starting server at port %s\n", port)
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(addr, nil))
}