package controllers

import (
	"encoding/json"
	"net/http"
	"rodrigo/go-REST/dao"
	"rodrigo/go-REST/entities"

	"github.com/gorilla/mux"
)

func GetTeams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	teamset := dao.ListAllTeams()
	json.NewEncoder(w).Encode(teamset)
}

func GetTeamsById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //params basicamente contém todos parâmetros que podem ser utilizados ex: localhost:8080/{id}/{outro_parametro_possivel}/{outro}
	//poderia ser router.HandleFunc("/teams/{id}/{param}", controllers.GetTeamsById).Methods("GET")
	teamset := dao.ListAllTeams()
	var found bool
	for _, team := range teamset {
		if team.ID == params["id"] { //params["id"] é uma string q vc digita na url localhost:8080/teams/1/oe, esse 1 vem como string nao como int
			found = true
			json.NewEncoder(w).Encode(team) // Encode > escreve o que ta na variavel team para JSON
		}
	}
	if !found {
		w.WriteHeader(404)
	}
}

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var team entities.Team
	_ = json.NewDecoder(r.Body).Decode(&team) // Aqui eu decodifico o body da requisicao, que estará em JSON, contendo os dados do user
	dao.InsertTeam(team)

}

func DeleteTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if dao.CheckExistanceById(params["id"]) {
		dao.DeleteTeamById(params["id"])
	} else {
		w.WriteHeader(404)
	}
}

func UpdateTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var team entities.Team
	_ = json.NewDecoder(r.Body).Decode(&team)
	params := mux.Vars(r)
	if dao.CheckExistanceById(params["id"]) {
		dao.UpdateTeamById(params["id"], team)
	} else {
		w.WriteHeader(404)
	}
}
