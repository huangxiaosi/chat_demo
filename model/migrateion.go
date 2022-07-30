package model

func migrateion() {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{})
	if err != nil {
		return
	}
}

//func migrateion() {
//	//自动迁移模式
//	err := DB.Set("gorm:table_options", "charset=utf8mb4").
//		AutoMigrate(&User{})
//	if err != nil {
//		return
//	}
//	//DB.Model(&Task{}).AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE")
//}
