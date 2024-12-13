package main

import (
	"flag"
	"log"

	_ "github.com/lib/pq"
	"web-11/internal/hello/api"
	"web-11/internal/hello/config"
	"web-11/internal/hello/provider"
	"web-11/internal/hello/usecase"
)

func main() {
	configPath := flag.String("config-path", "../../configs/hello_example.yaml", "путь к файлу конфигурации")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	use := usecase.NewUsecase(cfg.Usecase.DefaultMessage, prv)
	srv := api.NewServer(cfg.IP, cfg.Port, cfg.API.MaxMessageSize, use)

	log.Printf("Сервер запущен на %s\n", srv.Address)
	srv.Run()
}