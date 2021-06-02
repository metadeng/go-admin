package models

import (
	"odmp-admin/common/models"
)

type SysFileInfo struct {
	models.Model
	Type    string `json:"type" gorm:"type:varchar(255);comment:类型"`     //
	Name    string `json:"name" gorm:"type:varchar(255);comment:名称"`     //
	Size    string `json:"size" gorm:"type:int(11);comment:大小"`          //
	PId     int    `json:"pId" gorm:"type:int(11);comment:目录"`           //
	Source  string `json:"source" gorm:"type:varchar(255);comment:来源"`   //
	Url     string `json:"url" gorm:"type:varchar(255);comment:地址"`      //
	FullUrl string `json:"fullUrl" gorm:"type:varchar(255);comment:全地址"` //
	models.ControlBy
	models.ModelTime
}

func (SysFileInfo) TableName() string {
	return "sys_file_info"
}

func (e *SysFileInfo) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysFileInfo) GetId() interface{} {
	return e.Id
}
