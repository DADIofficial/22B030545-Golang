package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	TD "tsis1/pkg/Total_Drama_Characters"

	"github.com/gorilla/mux"
)

const NOT_FOUND = "these characters doesn't exist"

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("API is working"))
}

func Total_Drama(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/json")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(TD.Characters)
}

func find_Gender(gender string) ([]TD.Person, bool) {
	var characters []TD.Person
	for i := range TD.Characters {
		if TD.Characters[i].Gender == gender {
			characters = append(characters, TD.Characters[i])
		}
	}
	return characters, len(characters) > 0

}

func find_gender(w http.ResponseWriter, r *http.Request) {
	argv := mux.Vars(r)
	gender, ok := argv["Gender"]
	if !ok {
		http.Error(w, NOT_FOUND, http.StatusBadRequest)
		return
	}

	characters, ok := find_Gender(gender)
	if !ok {
		http.Error(w, NOT_FOUND, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "applications/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(characters)
}

func find_Team(team string) ([]TD.Person, bool) {
	var characters []TD.Person
	for i := range TD.Characters {
		if TD.Characters[i].F_team == team {
			characters = append(characters, TD.Characters[i])
		}
	}
	return characters, len(characters) > 0

}

func find_team(w http.ResponseWriter, r *http.Request) {
	argv := mux.Vars(r)
	team, ok := argv["Favteam"]
	if !ok {
		http.Error(w, NOT_FOUND, http.StatusBadRequest)
		return
	}

	characters, ok := find_Team(team)
	if !ok {
		http.Error(w, NOT_FOUND, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "applications/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(characters)
}

func find_Hair(team string) ([]TD.Person, bool) {
	var characters []TD.Person
	for i := range TD.Characters {
		if TD.Characters[i].Hair == team {
			characters = append(characters, TD.Characters[i])
		}
	}
	return characters, len(characters) > 0

}

func find_hair(w http.ResponseWriter, r *http.Request) {
	argv := mux.Vars(r)
	team, ok := argv["Hair"]
	if !ok {
		http.Error(w, NOT_FOUND, http.StatusBadRequest)
		return
	}

	characters, ok := find_Hair(team)
	if !ok {
		http.Error(w, NOT_FOUND, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "applications/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(characters)
}

func find_Position(team string) ([]TD.Person, bool) {
	var characters []TD.Person
	for i := range TD.Characters {
		if TD.Characters[i].Position == team {
			characters = append(characters, TD.Characters[i])
		}
	}
	return characters, len(characters) > 0

}

func find_position(w http.ResponseWriter, r *http.Request) {
	argv := mux.Vars(r)
	team, ok := argv["Position"]
	if !ok {
		http.Error(w, NOT_FOUND, http.StatusBadRequest)
		return
	}

	characters, ok := find_Position(team)
	if !ok {
		http.Error(w, NOT_FOUND, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "applications/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(characters)
}

func find_Power(team int) ([]TD.Person, bool) {
	var characters []TD.Person
	for i := range TD.Characters {
		if TD.Characters[i].Power_10 == team {
			characters = append(characters, TD.Characters[i])
		}
	}
	return characters, len(characters) > 0

}

func find_power(w http.ResponseWriter, r *http.Request) {
	argv := mux.Vars(r)
	team_str, ok := argv["Power"]
	if !ok {
		http.Error(w, NOT_FOUND, http.StatusBadRequest)
		return
	}

	team, err := strconv.Atoi(team_str)
	if err != nil {
		http.Error(w, NOT_FOUND, http.StatusBadRequest)
		return
	}

	characters, ok := find_Power(team)
	if !ok {
		http.Error(w, NOT_FOUND, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "applications/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(characters)
}

func find_Season(team int) ([]TD.Person, bool) {
	var characters []TD.Person
	for i := range TD.Characters {
		if TD.Characters[i].Number_seasons == team {
			characters = append(characters, TD.Characters[i])
		}
	}
	return characters, len(characters) > 0

}

func find_season(w http.ResponseWriter, r *http.Request) {
	argv := mux.Vars(r)
	team_str, ok := argv["Number_seasons"]
	if !ok {
		http.Error(w, NOT_FOUND, http.StatusBadRequest)
		return
	}

	team, err := strconv.Atoi(team_str)
	if err != nil {
		http.Error(w, NOT_FOUND, http.StatusBadRequest)
		return
	}

	characters, ok := find_Season(team)
	if !ok {
		http.Error(w, NOT_FOUND, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "applications/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(characters)
}

func find_Char(team string) ([]TD.Person, bool) {
	var characters []TD.Person
	for i := range TD.Characters {
		if TD.Characters[i].Name == team {
			characters = append(characters, TD.Characters[i])
		}
	}
	return characters, len(characters) > 0

}

func find_char(w http.ResponseWriter, r *http.Request) {
	argv := mux.Vars(r)
	team, ok := argv["Name"]
	if !ok {
		http.Error(w, NOT_FOUND, http.StatusBadRequest)
		return
	}

	characters, ok := find_Char(team)
	if !ok {
		http.Error(w, NOT_FOUND, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "applications/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(characters)
}
