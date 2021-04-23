package availability

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func DeleteAvailability(id int64,email string)  (*model.Availability,error){
	db :=env.RDB
	user := model.User{}
	db.Model(model.User{}).Where("email = ?",email).First(&user)
	availability := model.Availability{}
	availability.PreloadAvailability(db).First(&availability,id)
		err := db.Delete(availability).Error
		if err!=nil{
			return nil,err
		}else{
			return &availability,nil
		}


}