package data

import (
	"rodrigo/go-REST/entities"
)

var Teams []entities.Team

func GenerateTeams() []entities.Team {
	var teamset = []entities.Team{
		{ID: "1", Name: "Sport Recife", City: "Recife", State: "PE", Stadium: "Ilha do Retiro", LeagueDivision: "B"},
		{ID: "2", Name: "Nautico", City: "Recife", State: "PE", Stadium: "Aflitos", LeagueDivision: "B"},
		{ID: "3", Name: "Santa Cruz", City: "Recife", State: "PE", Stadium: "Arruda", LeagueDivision: "D"},
		{ID: "4", Name: "Guarani", City: "Campinas", State: "SP", Stadium: "Brinco de Ouro", LeagueDivision: "B"},
		{ID: "5", Name: "Ponte Preta", City: "Campinas", State: "SP", Stadium: "Moises Lucarelli", LeagueDivision: "B"},
	}

	return teamset
}
