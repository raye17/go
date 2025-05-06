package dao

// func CreateFileInfo(tx *gorm.DB, info *model.FileInfo) error {
// 	if tx == nil {
// 		tx = app.ModuleClients.DciDB
// 	}
// 	err := tx.Model(&model.FileInfo{}).Create(&info).Error
// 	return err
// }

// func UpdateFileInfo(tx *gorm.DB, info *model.FileInfo) error {
// 	if tx == nil {
// 		tx = app.ModuleClients.DciDB
// 	}
// 	err := tx.Model(&model.FileInfo{}).Where("id =?", info.Model.ID).Updates(&info).Error
// 	return err
// }
