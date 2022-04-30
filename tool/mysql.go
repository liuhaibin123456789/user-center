package tool

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"user-center/model"
)

var GDb *gorm.DB //gorm的db对象

func LinkMysql() error {
	//连接数据库，关闭默认启动事务
	db, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/user_center?charset=utf8mb4&loc=Local&parseTime=true"), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		return err
	}
	GDb = db
	err = createTables()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// CreateTables 初始化表格
func createTables() error {
	//校验表是否已经存在，存在就不创建
	tx := GDb.Begin()

	//遇到panic时，回滚
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	//用户隐私数据表
	if !tx.Migrator().HasTable(&model.User{}) {
		err := tx.AutoMigrate(&model.User{})
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	//用户非隐私数据表
	if !tx.Migrator().HasTable(&model.UserSide{}) {
		err := tx.AutoMigrate(&model.UserSide{})
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
