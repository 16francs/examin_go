package main

import  (
	"github.com/16francs/examin_go/infrastructure/datastore"

	"github.com/16francs/examin_go/domain/model"
)

func main() {
	db := datastore.Connect()
  defer db.Close()

	// テーブルの作成
	// db.AutoMifrate(&model.User{}) <- User テーブルの作成

	db.AutoMigrate(&model.User{})
} 