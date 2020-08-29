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
	s.DB.Save(&product)
	return &rpc.ProductResponse{Id: int32(product.ID), Name: product.Name, Stock: product.Stock}, nil
}

// GetProduct function implementation
func (s *Server) GetProduct(context context.Context, in *rpc.GetProductRequest) (*rpc.ProductResponse, error) {
	var product products.Product
	s.DB.Where("id = ?", in.Id).Find(&product).Limit(1)
	return &rpc.ProductResponse{Id: int32(product.ID), Name: product.Name, Stock: product.Stock}, nil
}

func main() {
	db := configs.InitDB()

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
