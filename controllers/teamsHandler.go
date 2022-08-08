package controllers

import (
	"encoding/json"
	"net/http"
	"rodrigo/go-REST/data"
	"rodrigo/go-REST/entities"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func GetTeams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Teams)
}

func GetTeamsById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //params basicamente contém todos parâmetros que podem ser utilizados ex: localhost:8080/{id}/{outro_parametro_possivel}/{outro}
	//poderia ser router.HandleFunc("/teams/{id}/{param}", controllers.GetTeamsById).Methods("GET")

	for _, team := range data.Teams {
		if team.ID == params["id"] { //params["id"] é uma string q vc digita na url localhost:8080/teams/1/oe, esse 1 vem como string nao como int
			json.NewEncoder(w).Encode(team) // Encode > escreve o que ta na variavel user para JSON
		}
	}
}

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var team entities.Team

	_ = json.NewDecoder(r.Body).Decode(&team)          // Aqui eu decodifico o body da requisicao, que estará em JSON, contendo os dados do user
	team.ID = strconv.FormatInt(time.Now().Unix(), 10) //gerando numero aleatorio baseado no tempo e convertendo pra string
	data.Teams = append(data.Teams, team)
	json.NewEncoder(w).Encode(team) // Encode > escreve o que ta na variavel (team) para JSON

}

func DeleteTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, team := range data.Teams {
		if team.ID == params["id"] {
			data.Teams = append(data.Teams[:index], data.Teams[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(data.Teams)
}

func UpdateTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, team := range data.Teams {

		if team.ID == params["id"] {
			data.Teams = append(data.Teams[:index], data.Teams[index+1:]...)
			var team entities.Team
			team.ID = params["id"]
			_ = json.NewDecoder(r.Body).Decode(&team)
			data.Teams = append(data.Teams, team)
			json.NewEncoder(w).Encode(data.Teams)
		}
	}
}
