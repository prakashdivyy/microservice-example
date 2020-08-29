package main

import (
	"fmt"
	"log"
	"net"

	"microservice/product-service/configs"
	"microservice/product-service/products"
	"microservice/product-service/rpc"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Server model
type Server struct {
	DB *gorm.DB
}

// CreateProduct function implementation
func (s *Server) CreateProduct(context context.Context, in *rpc.CreateProductRequest) (*rpc.ProductResponse, error) {
	product := products.Product{Name: in.Name, Stock: in.Stock}
	log.Printf("Creating product with name %s and stock %d", product.Name, product.Stock)
	err := s.DB.Save(&product).Error
	if err != nil {
		return &rpc.ProductResponse{}, err
	}
	log.Printf("Success create product with name %s and id %d", product.Name, product.ID)
	return &rpc.ProductResponse{Id: int32(product.ID), Name: product.Name, Stock: product.Stock}, nil
}

// GetProduct function implementation
func (s *Server) GetProduct(context context.Context, in *rpc.GetProductRequest) (*rpc.ProductResponse, error) {
	log.Printf("Get product with id %d", in.Id)
	var product products.Product
	s.DB.Where("id = ?", in.Id).Find(&product).Limit(1)
	if product.ID == 0 {
		log.Printf("Product with id %d not found", in.Id)
	} else {
		log.Printf("Product with id %d found", in.Id)
	}
	return &rpc.ProductResponse{Id: int32(product.ID), Name: product.Name, Stock: product.Stock}, nil
}

func main() {
	db := configs.InitDB()

	log.Println("Starting product-service on port 50050...")

	// create a listener on TCP port 50050
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50050))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a gRPC server object
	grpcServer := grpc.NewServer()
	// attach the Ping service to the server
	rpc.RegisterProductServiceServer(grpcServer, &Server{DB: db})

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
