package main

import (
	"fmt"
	"log"
	"net/http"
	"rodrigo/go-REST/controllers"
	"rodrigo/go-REST/data"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func HandleRequest() {

	fmt.Println("Starting server at port: 8080")
	router.HandleFunc("/teams", controllers.GetTeams).Methods("GET")
	router.HandleFunc("/teams/{id}", controllers.GetTeamsById).Methods("GET")
	router.HandleFunc("/teams", controllers.CreateTeam).Methods("POST")
	router.HandleFunc("/teams/{id}", controllers.DeleteTeam).Methods("DELETE")
	router.HandleFunc("/teams/{id}", controllers.UpdateTeam).Methods("PUT")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func main() {

	data.OpenDb()
	defer data.Db.Close()
	HandleRequest()

}
