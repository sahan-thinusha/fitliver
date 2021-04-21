package doctor


import (
	"errors"
	"fitliver/pkg/env"
	"fitliver/pkg/model"
	"strings"
)

func ApproveDoctor(id int,approve bool,role string)  (*model.Doctor,error){
	db :=env.RDB

	if strings.EqualFold(role,env.ADMIN){
		err := db.Model(&model.Doctor{}).Where("id = ?", id).Update("is_new",approve).Error
		if err!=nil{
			return nil,err
		}else{
			d := model.Doctor{}
			db.Model(&model.Doctor{}).First(&d,id)
			return &d,nil
		}
	}else{
		return nil,errors.New("permission denied...")
	}

}