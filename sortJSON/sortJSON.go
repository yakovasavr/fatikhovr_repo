package main

import (
	"fmt"
	// "log"
	"net/http"
	"os"
	"io/ioutil"
	"encoding/json"
	"sort"
)

type sl []int

func (arru sl) makeArr() []int{
	i := 0
	for _, value := range arru {
		if value > 0 {
			arru[i] = value
			i++
		}
	}
	return (arru[:i])
}

func main() {
	var addr string
	var port string
	var a string
	var b string
	var arru sl

	fmt.Scanf("%s\n", &addr)
	fmt.Scanf("%s\n", &port)
	fmt.Scanf("%s\n", &a)
	fmt.Scanf("%s\n", &b)

	addr = "http://" + addr[7:] + ":" + port + "?a=" + a + "&b=" + b
	// fmt.Println(addr)

	resp, err := http.Get(addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error", err)
	}
	// fmt.Printf("%s", body)

	err2 := json.Unmarshal(body, &arru)
	if err2 != nil {
		fmt.Println("error:", err2)
	}

	// fmt.Printf("%+v", arru)
	arru2 := arru.makeArr()
	sort.Sort(sort.Reverse(sort.IntSlice(arru2)))
	for i, value := range arru2 {
		if i != 0 {
			fmt.Printf("%c", '\n')
		}
		fmt.Printf("%d", value)
	}
}