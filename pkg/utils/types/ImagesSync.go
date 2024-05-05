package types

import "time"

type ImagesSync struct {
	Id             uint      `gorm:"primarykey"`
	SourceRegistry string    `gorm:"type:varchar(50);comment:'原仓库url';"`
	SourceRepo     string    `gorm:"type:varchar(100);comment:'原仓库镜像名称';"`
	SourceTag      string    `gorm:"type:varchar(50);comment:'原仓库镜像标签';"`
	DestRegistry   string    `gorm:"type:varchar(50);comment:'目标仓库url';"`
	DestRepo       string    `gorm:"type:varchar(100);comment:'目标仓库镜像名称';"`
	DestTag        string    `gorm:"type:varchar(50);comment:'目标仓库镜像标签';"`
	SyncStatus     string    `gorm:"type:varchar(1);comment:'同步状态：0-未同步 1-同步成功 2-同步失败';"`
	CreateTime     time.Time `gorm:"comment:'创建时间';"`
	UpdateTime     time.Time `gorm:"comment:'更新时间';"`
}
