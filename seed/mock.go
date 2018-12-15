package seed

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sandbox-api/model"
)

func AllFoods() ([]model.Food, error) {
	file, err := os.Open("/go/src/sandbox-api/seed/food50.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	byteJson, _ := ioutil.ReadAll(file)
	foods := Foods{}
	json.Unmarshal(byteJson, &foods)
	return foods.Foods, nil
}

func AllMotions() ([]model.Motion, error) {
	file, err := os.Open("/go/src/sandbox-api/seed/motion.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	byteJson, _ := ioutil.ReadAll(file)
	motions := Motions{}
	json.Unmarshal(byteJson, &motions)
	return motions.Motions, nil
}

type Motions struct {
	Motions []model.Motion `json:"Motions"`
}

type Foods struct {
	Foods []model.Food `json:"Foods"`
}
