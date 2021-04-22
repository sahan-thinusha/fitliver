package doctor

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func GetNewDoctors()  ([]*model.Doctor,error){
	db :=env.RDB
	doctors := []*model.Doctor{}
	db.Model(model.Doctor{}).Where("is_new = ?",env.STATUS_NEW).Scan(&doctors)
	return doctors,nil

}
