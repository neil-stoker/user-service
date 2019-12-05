package main

import (
	"fmt"
	"github.com/micro/go-micro"
	pb "github.com/neil-stoker/user-service/proto/user"
	"log"
)

func main() {
	db, err := CreateConnection()
	defer db.Close()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService})
	//cannot use &(service literal) (value of type *service) as user.UserServiceHandler value in argument to
	//  pb.RegisterUserServiceHandler: wrong type for method ValidateToken

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
