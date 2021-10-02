package main

import (
	"fmt"
	"log"
    "os"
	"encoding/json"
	"io"
	// "sort"
)

type offer struct {
	Offer_id	string `json:"offer_id"`
	Market_sku	int `json:"market_sku"`
	Price		int `json:"price"`
}

type OffersList struct {
	Offers	[]offer `json:"offers"`
}

func main() {
	var n int
	var m int
	i := 0

	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &m)

    // err := json.NewDecoder(os.Stdin).Decode(&list[0])
    // if err != nil {
    //     log.Fatal(err)
    // }
	list := make([]OffersList, n)
	dec := json.NewDecoder(os.Stdin)
    for i < n {
        err := dec.Decode(&list[i])
        if err == io.EOF {
            return
        }
        if err != nil {
            log.Fatal(err)
        }
		i++
    }

	newList := OffersList{}
	i = 0
	j := 0
	for i < n {
		for _, val := range list[i].Offers {
			newList.Offers = append(newList.Offers, val)
			j++
			if j == m {
				break
			}
		}
		i++
		if j == m {
			break
		}
	}
	// fmt.Printf("%d\n", j)
	data, err := json.Marshal(newList)
	if err != nil {
		log.Fatal("marshalling problem: %s", err)
	}
	fmt.Printf("%s\n", data)
}