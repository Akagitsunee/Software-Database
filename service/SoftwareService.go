package service

import (
	"softwareDB/models"
	"softwareDB/repository"
	"sort"
	"strconv"
)

var repo = repository.NewJsonRepo()
// "constructor" variable
var _ = repo.Init()

func Add(software models.Software) models.Software {
	softwareList := FindAll()
	sort.Slice(softwareList, func(i, j int) bool {
		return softwareList[i].Id > softwareList[j].Id
	})

	idToInt, _ := strconv.Atoi(softwareList[0].Id)
	software.Id = strconv.Itoa(idToInt + 1)

	repo.Add(software)
	return software
}

func FindAll() []models.Software  {
	return repo.FindAll()
}

func FindByName(name string) []models.Software {
	return repo.FindByName(name)
}

func Update(id string, updatedSoftware *models.Software) *models.Software {
	softwareList := repo.FindAll()

	for _, software := range softwareList {
		if software.Id == id {
			software.Name = updatedSoftware.Name
			software.Description = updatedSoftware.Description
			software.Homepage = updatedSoftware.Homepage
			software.Version = updatedSoftware.Version
			software.License = updatedSoftware.License
			repo.Update(software)
			return &software
		}
	}
	return nil
}

func Delete(id string) {
	repo.Delete(id)
}
