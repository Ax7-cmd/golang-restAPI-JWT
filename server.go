package main
import (
	"fmt"
	"net"

	mgrpc "golang-restAPI-JWT/Core/Grpc"
	log "github.com/Sirupsen/logrus"
	pb "golang-restAPI-JWT/Core/Grpc/Services"
	
	"golang-restAPI-JWT/Config"
	"golang-restAPI-JWT/Database"
	"golang-restAPI-JWT/Database/Seed"
	"golang-restAPI-JWT/Core/Router"
	"golang-restAPI-JWT/Core/Models"
	"google.golang.org/grpc"
)

// Api server start from here. router is define your api router and public it.
func main() {
	// GORM DATABASE
	Database.Mysql, Database.Err = Database.ConnectToDB("main")
	if Database.Err != nil {
		fmt.Println("status error : ", Database.Err)
	} else {
		fmt.Println("database connected")
	}
	defer Database.Mysql.Close()
	// auto migrate
	Database.Mysql.AutoMigrate(&Models.User{})
	Database.Mysql.AutoMigrate(&Models.Category{})
	Database.Mysql.AutoMigrate(&Models.Product{})
	Database.Mysql.AutoMigrate(&Models.Cart{})
	var category Models.Category
	var product Models.Product
	err := Models.FirstCategory(&category)
	err = Models.FirstProduct(&product)
	if err != nil && category.Name == "" && product.Name == "" {
		err = Seed.CategoryProductSeed()
		if err != nil {
			fmt.Println("seeder error : ", err)
		}
	}


	// // Redis DB
	// Redis.Client = Redis.NewClient()

	// GRPC
	// Here will enable grpc server, if you don`t want it, you can disable it
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 10000))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		var opts []grpc.ServerOption
		grpcServer := grpc.NewServer(opts...)
		pb.RegisterRouteGuideServer(grpcServer, mgrpc.NewServer())
		grpcServer.Serve(lis)
	}()
	app_env := Config.GoDotEnvVariable("APP_ENV")

	Router.Start(app_env)
}
