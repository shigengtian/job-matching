package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func StartWebServer() error {
	fmt.Println("Rest API with Mux Routers")
	router := mux.NewRouter().StrictSlash(true)

	// // router.HandleFunc({ エンドポイント }, { レスポンス関数 }).Methods({ リクエストメソッド（複数可能） })
	// router.HandleFunc("/", rootPage)
	// router.HandleFunc("/items", fetchAllItems).Methods("GET")
	// router.HandleFunc("/item/{id}", fetchSingleItem).Methods("GET")

	// router.HandleFunc("/item", createItem).Methods("POST")
	// router.HandleFunc("/item/{id}", deleteItem).Methods("DELETE")
	// router.HandleFunc("/item/{id}", updateItem).Methods("PUT")

	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}
