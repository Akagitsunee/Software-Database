package repository

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
	"softwareDB/models"
)

type JsonRepository struct {
	softwareList []models.Software
}

func (repository *JsonRepository) Init() bool {
	file, _ := ioutil.ReadFile("software.json")
	json.Unmarshal(file, &repository.softwareList)
	return true
}

func (repository *JsonRepository) Add(software models.Software) {
	repository.softwareList = append(repository.softwareList, software)
	repository.writeFile()
}

func (repository *JsonRepository) FindAll() []models.Software {
	return repository.softwareList
}

func (repository *JsonRepository) FindByName(name string) []models.Software {
	var softwareListByName []models.Software

	for _, software := range repository.softwareList {
		match, _ := regexp.Match(".{0,}" + name + ".{0,}", []byte(software.Name))
		if match {
			softwareListByName = append(softwareListByName, software)
		}
	}
	return softwareListByName
}

func (repository *JsonRepository) Update(software models.Software)  {
	for i, sf := range repository.softwareList {
		if sf.Id == software.Id {
			repository.softwareList[i] = software
		}
	}
	repository.writeFile()
}

func (repository *JsonRepository) Delete(id string)  {
	for i, sf := range repository.softwareList {
		if sf.Id == id {
			repository.softwareList = append(repository.softwareList[:i], repository.softwareList[i+1:]...)
		}
	}
	repository.writeFile()
}

func (repository *JsonRepository) writeFile()  {
	file, _ := json.Marshal(repository.softwareList)
	ioutil.WriteFile("softwareback.jsons", file, 0644)
}

func NewJsonRepo() JsonRepository {
	return JsonRepository{}
}
