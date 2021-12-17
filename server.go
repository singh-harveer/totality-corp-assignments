package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"totality/users/apis"
	"totality/users/client"
	"totality/users/dao/localdb"
	"totality/users/service"
	"totality/users/totality"
	"totality/users/userpb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	portGRPC = 9001
	apiPort  = 8000
)

func main() {
	var addr = fmt.Sprintf(":%d", portGRPC)
	log.Println("starting grpc service")

	var listen, err = net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on %s:\n%v", addr, err)
	}

	var userManager totality.UserManager
	userManager, err = localdb.NewLocalDB()
	if err != nil {
		log.Fatalf("failed to initialize new user manager: %v", err)
	}

	var server = grpc.NewServer()
	userpb.RegisterUserServiceServer(server, &service.UserService{UserDB: userManager})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := server.Serve(listen); err != nil {
			log.Fatalf("failed to serve grpc server over port %d : %v", portGRPC, err)
		}
	}()

	var userManagerClient totality.UserManager
	userManagerClient, err = client.NewUserClient(context.Background(), addr)
	if err != nil {
		log.Fatalf("failed to create grpc client: %v", err)
	}

	var handler = apis.NewHandler(userManagerClient)

	var router = gin.Default()
	router.GET("/users/:id", handler.GetUserByID)

	router.POST("/users", handler.GetUsers)

	var apiAddr = fmt.Sprintf(":%d", apiPort)

	var srv = &http.Server{
		Addr:    apiAddr,
		Handler: router,
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("failed to listen on: %s\n", err)
		}
	}()

	log.Printf("user api server started successfully on %s", apiAddr)

	// handle shutdown
	var ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		s := <-sigCh
		log.Printf("got signal %v, attempting graceful shutdown", s)
		cancel()
		//Stop grpc service
		server.GracefulStop()
		//stop api server
		srv.Shutdown(ctx)
	}()

	wg.Wait()
}
