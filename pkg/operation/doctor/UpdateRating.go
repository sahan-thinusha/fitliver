package doctor



import (
	"errors"
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func RatingUpdate(rating *model.Rating,email string)  (*model.Rating,error){
	db :=env.RDB
	user := model.User{}
	db.Model(model.User{}).Where("email = ?",email).First(&user)
	rate := model.Rating{}
	rating.PreloadRating(db).First(&rate,rating.ID)
	if rate.User.ID == user.ID{
		rate.Rate = rating.Rate
		err := db.Save(rate).Error
		if err!=nil{
			return nil,err
		}else{
			return &rate,nil
		}
	}else{
		return nil,errors.New("permission denied...")
	}

}