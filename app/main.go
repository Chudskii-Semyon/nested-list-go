package main

import (
	"awesomeProject/domain"
	_listDelivery "awesomeProject/list/http/delivery"
	_listRepository "awesomeProject/list/repository/postgres"
	_listUsecase "awesomeProject/list/usecase"

	_itemDelivery "awesomeProject/item/http/delivery"
	_itemRepository "awesomeProject/item/repository/postgres"
	_itemUsecase "awesomeProject/item/usecase"

	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool("debug") {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbType := viper.GetString("database.type")
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbPass := viper.GetString("database.password")
	dbUser := viper.GetString("database.user")
	dbName := viper.GetString("database.dbName")
	sslmode := viper.GetString("database.sslmode")

	dbConn, err := dbFactory(dbType, dbHost, dbPort, dbPass, dbUser, dbName, sslmode)

	err = dbConn.DB().Ping()

	if err != nil {
		log.Fatal("ping =>", err)
	}

	dbConn.AutoMigrate(&domain.List{}, domain.Item{})

	defer func() {
		err := dbConn.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	listRepository := _listRepository.NewPostgresListRepository(dbConn)
	listUsecase := _listUsecase.NewListUsecase(listRepository)

	itemRepository := _itemRepository.NewPostgresItemRepository(dbConn)
	itemUsecase := _itemUsecase.NewItemUsecase(itemRepository)

	r := httprouter.New()

	_listDelivery.NewListHandler(r, listUsecase)
	_itemDelivery.NewItemHandler(r, itemUsecase)

	log.Fatal(http.ListenAndServe(viper.GetString("server.address"), r))
}

func dbFactory(dbType, host, port, pass, user, name, sslmode string) (dbConn *gorm.DB, err error) {
	connection := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s", dbType, user, pass, host, port, name, sslmode)

	return gorm.Open(dbType, connection)
}
