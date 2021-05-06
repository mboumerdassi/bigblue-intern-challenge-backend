package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var url = "http://localhost:8080/products"

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

func main() {
	client := http.DefaultClient
	//// TODO use the client to call the local server using the method GET
	//// and the url
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		log.Fatal(err)
	}

	///// TODO get the products inside a local variable products
	//// and display one of them
	products := ProductsResponse{}
	err = json.Unmarshal(body, &products)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(*products.Products[0])
}
