package handlers

import (
	"TECHTEST_BE/models"
	"TECHTEST_BE/utils"
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// GetCakesHandler handles the GET request to /cakes
func GetCakesHandler(w http.ResponseWriter, r *http.Request) {
	cakes, err := models.GetCakes()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Sort cakes by rating and title
	sort.Slice(cakes, func(i, j int) bool {
		if cakes[i].Rating == cakes[j].Rating {
			return cakes[i].Title < cakes[j].Title
		}
		return cakes[i].Rating > cakes[j].Rating
	})

	utils.RespondWithJSON(w, http.StatusOK, cakes)
}

// GetCakeHandler handles the GET request to /cakes/{id}
func GetCakeIDHandler(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	cake, err := models.GetCake(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if cake == nil {
		utils.RespondWithError(w, http.StatusNotFound, "Cake not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, cake)
}

// AddCakeHandler handles the POST request to /cakes
func AddCakeHandler(w http.ResponseWriter, r *http.Request) {
	var cake models.Cake
	err := json.NewDecoder(r.Body).Decode(&cake)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = models.AddCake(&cake)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Cake added successfully")
}

// UpdateCakeHandler handles the PUT request to /cakes/{id}
func UpdateCakeHandler(w http.ResponseWriter, r *http.Request) {
	idParam := GetURLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var cake models.Cake
	err = json.NewDecoder(r.Body).Decode(&cake)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = models.UpdateCake(id, &cake)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Cake not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Cake updated successfully")
}

// DeleteCakeHandler handles the DELETE request to /cakes/{id}
func DeleteCakeHandler(w http.ResponseWriter, r *http.Request) {
	idParam := GetURLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	err = models.DeleteCake(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Cake not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Cake deleted successfully")
}

// GetURLParam extracts a URL parameter from the request URL
func GetURLParam(r *http.Request, paramName string) string {
	splitPath := strings.Split(r.URL.Path, "/")
	for i := range splitPath {
		if splitPath[i] == paramName && i+1 < len(splitPath) {
			return splitPath[i+1]
		}
	}
	return ""
}
