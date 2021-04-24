package booking

import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func BookingServiceCreate(bookingService *model.BookingService,email string)  (*model.BookingService,error){
	db :=env.RDB
	user := model.User{}
	user.PreloadDoctor(db).Model(model.User{}).Where("email = ?",email).First(&user)
	bookingService.Doctor = user.Doctor
	err := db.Create(bookingService).Error
	if err!=nil{
		return nil,err
	}else{
		return bookingService,nil
	}
}