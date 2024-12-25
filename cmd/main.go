package main

import (
	"log"

	"github.com/dwiprst13/restapi-go/cmd/api"
	"github.com/dwiprst13/restapi-go/config"
	"github.com/dwiprst13/restapi-go/db"
	"github.com/go-sql-driver/mysql"
	// "github.com/go-oauth2/oauth2/v4/server"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})
	if err != nil{
		log.Fatal(err)
	}


	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
