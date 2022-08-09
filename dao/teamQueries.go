package dao

import (
	"database/sql"
	"log"
	"rodrigo/go-REST/data"
	"rodrigo/go-REST/entities"
)

func ListAllTeams() []entities.Team {
	var teamset []entities.Team
	results, err := data.Db.Query("SELECT * FROM Teams")
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()
	for results.Next() {
		var team entities.Team
		err = results.Scan(&team.ID, &team.Name, &team.City, &team.State, &team.Stadium, &team.LeagueDivision)
		if err != nil {
			panic(err.Error())
		}
		teamset = append(teamset, team)
	}
	return teamset
}

func CheckExistanceById(id string) bool {
	var found_id string
	err := data.Db.QueryRow("SELECT id FROM Teams WHERE id = ? LIMIT 1", id).Scan(&found_id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return false
	}
	return true
}

func InsertTeam(team entities.Team) {
	insert, err := data.Db.Query(`INSERT INTO Teams(name,city,state,stadium,LeagueDivision)
	VALUES (?,?,?,?,?)`, team.Name, team.City, team.State, team.Stadium, team.LeagueDivision)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func DeleteTeamById(id string) {
	delete, err := data.Db.Query("DELETE FROM Teams WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}
	defer delete.Close()
}

func UpdateTeamById(id string, team entities.Team) {

	update, err := data.Db.Query(`UPDATE Teams
	SET 
		name = ?,
		city = ?,
		state = ?,
		stadium = ?,
		LeagueDivision = ?
	WHERE
		id = ?;`, team.Name, team.City, team.State, team.Stadium, team.LeagueDivision, id)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
	defer update.Close()
}
