package main

import (
	"context"
	"log"

	"microservice/product-service/rpc"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50050", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := rpc.NewProductServiceClient(conn)

	product, err := c.GetProduct(context.Background(), &rpc.GetProductRequest{Id: 1})

	if err != nil {
		log.Fatalf("Error when calling GetProduct: %s", err)
	}

	log.Printf("Response from server: %v", product)

	product2, err2 := c.CreateProduct(context.Background(), &rpc.CreateProductRequest{Name: "Product A", Stock: 10})

	if err2 != nil {
		log.Fatalf("Error when calling CreateProduct: %s", err)
	}

	log.Printf("Response from server: %v", product2)
}
