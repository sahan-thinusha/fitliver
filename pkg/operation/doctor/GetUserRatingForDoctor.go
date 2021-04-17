package doctor

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func GetUserRatingForADoctor(id int64,email string)  (*model.Rating,error){
	db :=env.RDB
	user := model.User{}
	db.Model(model.User{}).Where("email = ?",email).First(&user)
	rate := model.Rating{}
	rate.PreloadRating(db).Where("user_id = ? AND doctor = ?",user.ID,id).First(&rate)
	return &rate,nil
}