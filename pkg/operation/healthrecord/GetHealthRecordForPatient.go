package healthrecord



import (
	"fitliver/pkg/env"
	"fitliver/pkg/model"
)

func GetHealthRecordByPatient(id int64)  ([]*model.HealthRecord,error){
	db :=env.RDB
	healthRecord := model.HealthRecord{}
	healthRecords :=[]*model.HealthRecord{}
	healthRecord.PreloadHealthRecord(db).Model(model.HealthRecord{}).Where("patient_id = ?",id).Find(&healthRecords)
	return healthRecords,nil
}