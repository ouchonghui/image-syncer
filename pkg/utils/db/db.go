package db

import (
	"fmt"
	"github.com/AliyunContainerService/image-syncer/pkg/utils/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

// InitDB 初始化数据库
func InitDB(dbConf types.MyQLConf) (err error) {
	dsn := dbConf.Username + ":" + dbConf.Password + "@tcp(" + dbConf.Host + ":" + dbConf.Port + ")/" + dbConf.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		return err
	}
	DB.Set("gorm:table_options", "COMMENT='镜像同步状态表' ENGINE=InnoDB CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci").AutoMigrate(&types.ImagesSync{})
	db, _ := DB.DB()
	return db.Ping()
}

func Close() error {
	db, _ := DB.DB()
	err := db.Close()
	if err != nil {
		return fmt.Errorf("failed to close MySQL: %v", err)
	}
	return nil
}
