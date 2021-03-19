package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"softwareDB/models"
	"softwareDB/service"
)

// ListSoftware godoc
// @Summary List details of all software
// @Description List details of all software
// @Tags Software
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Software
// @Router /software [get]
func ListSoftware(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(service.FindAll())
}

// CreateSoftware godoc
// @Summary Create a new software entry
// @Description Create a new software entry with the input paylod
// @Tags Software
// @Accept  json
// @Produce  json
// @Param software body models.Software true "Create software"
// @Success 200 {object} models.Software
// @Router /software [post]
func CreateSoftware(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var software models.Software
	json.NewDecoder(request.Body).Decode(&software)
	software = service.Add(software)
	json.NewEncoder(writer).Encode(software)

}

// GetSoftwareByName godoc
// @Summary Get details for a given software name
// @Description Get details of software corresponding to the input software name
// @Tags Software
// @Accept  json
// @Produce  json
// @Param name path string true "Name of the software"
// @Success 200 {array} models.Software
// @Router /software/{name} [get]
func GetSoftwareByName(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	softwareName := params["name"]
	software := service.FindByName(softwareName)
	json.NewEncoder(writer).Encode(software)
}

// UpdateSoftware godoc
// @Summary Update software identified by the given id
// @Description Update the software corresponding to the input id
// @Tags Software
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the software to be updated"
// @Param software body models.Software true "Updated software"
// @Success 200 {object} models.Software
// @Router /software/{id} [put]
func UpdateSoftware(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	softwareId := params["id"]
	var software models.Software
	json.NewDecoder(request.Body).Decode(&software)
	json.NewEncoder(writer).Encode(service.Update(softwareId, &software))
}

// DeleteSoftwareById godoc
// @Summary Delete software identified by the given id
// @Description Delete the software corresponding to the input id
// @Tags Software
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the software to be deleted"
// @Success 204 "No Content"
// @Router /software/{id} [delete]
func DeleteSoftwareById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	softwareId := params["id"]
	service.Delete(softwareId)
	writer.WriteHeader(http.StatusNoContent)
	return
}

// Shutdown godoc
// @Summary Shut the server down
// @Description Shut the server down in case of when the application is running in background
// @Tags Shutdown
// @Accept  json
// @Produce  json
// @Accept  json
// @Produce  json
// @Success 200
// @Router /shutdown [get]
func Shutdown(writer http.ResponseWriter, request *http.Request) {
	os.Exit(3)
}
