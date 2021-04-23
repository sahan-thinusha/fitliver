package workout

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func GetWorkouts()  ([]*model.Workout,error){
	db :=env.RDB
	woprkouts :=[]*model.Workout{}
	db.Model(model.Workout{}).Find(&woprkouts)
	return woprkouts,nil
}