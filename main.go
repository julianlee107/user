package main

import (
	"fmt"

	"gorm.io/gorm/schema"

	"github.com/julianlee107/user/domain/service"

	"github.com/julianlee107/user/domain/repository"

	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/julianlee107/user/handler"
	pb "github.com/julianlee107/user/proto/user/pb/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)
	// 初始化服务
	srv.Init()
	// 创建数据库
	dsn := "root:root@tcp(127.0.0.1:3306)/micro?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		fmt.Println(err)
	}

	// 创建服务实例
	userDataService := service.NewUserDataService(repository.NewUserRepository(db))
	// Register handler
	err = pb.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: userDataService})
	if err != nil {
		return
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
