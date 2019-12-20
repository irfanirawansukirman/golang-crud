package models

type ResponseMap struct {
	code        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	message  	string    `gorm:"size:255;not null;unique" json:"nickname"`
	data		User
}

// func ResMap(code uint32, message string, data *[]models.User) ResponseMap {
// 	mapmap := ResponseMap{
// 		code: code,
// 		message: message,
// 		data: data}
// 	return mapmap
// }