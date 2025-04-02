package main

import (
	"context"
	"log"
	"net"
	"net/http"

	// "github.com/blainsmith/grpc-gateway-openapi-example/gen/protos/go/protos"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	helloworldpb "github.com/myuser/myrepo/gen/go"
	"github.com/myuser/myrepo/service" // import the hello service package
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// start the gRPC server
	lis, err := net.Listen("tcp", "localhost:5566")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	helloworldpb.RegisterGreeterServer(grpcServer, service.NewServer())
	reflection.Register(grpcServer)
	log.Println("gRPC server ready on localhost:5566...")
	go grpcServer.Serve(lis)

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.NewClient(
		"0.0.0.0:5566",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	defer conn.Close()

	// create an HTTP router using the client connection above
	// and register it with the service client
	rmux := runtime.NewServeMux()
	err = helloworldpb.RegisterGreeterHandler(ctx, rmux, conn)
	if err != nil {
		log.Fatal(err)
	}

	// create a standard HTTP router
	mux := http.NewServeMux()

	// mount the gRPC HTTP gateway to the root
	mux.Handle("/", rmux)

	// mount a path to expose the generated OpenAPI specification on disk
	mux.HandleFunc("/swagger-ui/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./gen/go/hello_world.swagger.json")
	})

	// mount the Swagger UI that uses the OpenAPI specification path above
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./third_party/swagger-ui"))))

	log.Println("HTTP server ready on localhost:8080...")
	err = http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
