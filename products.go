package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	InitialStock int    `json:"initial_stock"`
	ImageURL     string `json:"image_url"`
}

type ProductsResponse struct {
	Products []*Product `json:"products"`
}

var products = []*Product{
	&Product{
		ID:           "SHOP-0011",
		Name:         "product 1",
		Price:        10,
		InitialStock: 20,
		ImageURL:     "https://img-0.journaldunet.com/pb5J23IF99HMY_P9HtO64G38epM=/1240x/smart/32d90de13a5f411c86709152f70fc67c/ccmcms-jdn/10861192.jpg",
	},
	&Product{
		ID:           "SHOP-0012",
		Name:         "product 1 in blue",
		Price:        11,
		InitialStock: 5,
		ImageURL:     "https://img-0.journaldunet.com/pb5J23IF99HMY_P9HtO64G38epM=/1240x/smart/32d90de13a5f411c86709152f70fc67c/ccmcms-jdn/10861192.jpg",
	},
	&Product{
		ID:           "SHOP-0021",
		Name:         "product 2",
		Price:        15,
		InitialStock: 10,
		ImageURL:     "https://img-0.journaldunet.com/pb5J23IF99HMY_P9HtO64G38epM=/1240x/smart/32d90de13a5f411c86709152f70fc67c/ccmcms-jdn/10861192.jpg",
	},
	&Product{
		ID:           "SHOP-0022",
		Name:         "product 2 big",
		Price:        30,
		InitialStock: 10,
		ImageURL:     "https://img-0.journaldunet.com/pb5J23IF99HMY_P9HtO64G38epM=/1240x/smart/32d90de13a5f411c86709152f70fc67c/ccmcms-jdn/10861192.jpg",
	},
	&Product{
		ID:           "SHOP-0031",
		Name:         "product 3",
		Price:        8,
		InitialStock: 20,
		ImageURL:     "https://img-0.journaldunet.com/pb5J23IF99HMY_P9HtO64G38epM=/1240x/smart/32d90de13a5f411c86709152f70fc67c/ccmcms-jdn/10861192.jpg",
	},
	&Product{
		ID:           "SHOP-0041",
		Name:         "product 4",
		Price:        100,
		InitialStock: 2,
		ImageURL:     "https://img-0.journaldunet.com/pb5J23IF99HMY_P9HtO64G38epM=/1240x/smart/32d90de13a5f411c86709152f70fc67c/ccmcms-jdn/10861192.jpg",
	},
}

// Handler replies with the products response to any request
func Handler(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(&ProductsResponse{
		Products: products,
	})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {

	fmt.Println("Starting server at port 8080. URL: http://localhost:8080/")
	http.HandleFunc("/products", Handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
