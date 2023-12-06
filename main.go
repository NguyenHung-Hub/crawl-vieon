package main

import (
	"crawl/src/db"
	"crawl/src/handler"
	"crawl/src/models"
	"crawl/src/util"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config: ", err)
	}
	dsn := "root:123456@tcp(127.0.0.1:3306)/vieon?charset=utf8mb4&parseTime=True&loc=Local"
	connect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("can not connect to database")
	} else {
		log.Println(connect)
	}

	store := db.NewStore(connect)

	err = connect.AutoMigrate(&models.Ribbon{})
	if err != nil {
		log.Fatal("can not migration Ribbon: ", err)
	}
	err = connect.AutoMigrate(&models.Content{})
	if err != nil {
		log.Fatal("can not migration Content: ", err)
	}
	err = connect.AutoMigrate(&models.UrlCrawl{}, &models.RibbonIDCrawl{}, &models.ContentIDCrawl{})
	if err != nil {
		log.Fatal("can not migration Info: ", err)
	}

	server, err := handler.NewServer(config, store)
	if err != nil {
		log.Println("can not create server")
	}

	server.Start()
}
