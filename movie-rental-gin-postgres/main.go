package main

import (
	"database/sql"
	"log"
	"movieRentals/api"
	db "movieRentals/db/sqlc"
	"movieRentals/model"
	"movieRentals/reader"
	"movieRentals/service"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	config := initConfig()
	reader, err := reader.NewFileReader(config.InitFile.Path)
	if err != nil {
		log.Fatalf("error reading movies json %v", err)
	}

	store := initdb(config.DbConfig.Driver, config.DbConfig.ConnStr)
	cartService := service.NewCartService(*model.NewCart(), store)
	priceService := service.NewPriceService(store)
	httpAddress := config.ServerConfig.Address

	runGinServer(store, reader, cartService, priceService, httpAddress)
}

func initConfig() *model.Config {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var config model.Config
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode config, %v", err)
	}

	return &config
}

func initdb(driver string, connStr string) db.Store {
	conn, err := sql.Open(driver, connStr)

	if err != nil {
		log.Fatalf("error connecting to database %v", err)
	}

	err = conn.Ping()

	if err != nil {
		log.Fatalf("error connecting to database %v", err)
	}

	return db.NewStore(conn)
}

func runGinServer(db db.Store, reader api.MoviesReader, cartService service.ICartService, priceService service.IPriceService, address string) {
	server, err := api.NewServer(db, reader, cartService, priceService)
	if err != nil {
		log.Fatal("cannot create server")
	}

	err = server.Start(address)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
