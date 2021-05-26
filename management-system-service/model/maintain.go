package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Maintain 维护表单数据库模型
type Maintain struct {
	gorm.Model
	UserID      uint      `gorm:"not null"` //维护表单发起用户
	EquipmentID uint      `gorm:"not null"` //维护设备id
	Date        time.Time `gorm:"not null"` //维护日期
	Status      uint8     `gorm:"not null"` //设备状态
	Remark      string    `gorm:"not null"` //异常信息
}

//HTTPAddMaintainInfo HTTP消息模型AddMaintain
type HTTPAddMaintainInfo struct {
	UserID      uint   `form:"user_id" json:"user_id" binding:"required"`
	EquipmentID uint   `form:"equipment_id" json:"equipment_id" binding:"required"`
	Date        uint64 `form:"date" json:"date" binding:"required"`
	Status      uint8  `form:"status" json:"status" binding:"required"`
	Remark      string `form:"remark" json:"remark" binding:"required"`
}

//HTTPAddMaintainInfo HTTP回复模型AddMaintain
type HTTPAddMaintainResponse struct {
	ID uint `form:"id" json:"id" binding:"required"`
}

//HTTPGetMaintainInfo HTTP回复模型GetMaintain
type HTTPGetMaintainResponse struct {
	UserID      uint   `form:"user_id" json:"user_id" binding:"required"`
	EquipmentID uint   `form:"equipment_id" json:"equipment_id" binding:"required"`
	Date        uint64 `form:"date" json:"date" binding:"required"`
	Status      uint8  `form:"status" json:"status" binding:"required"`
	Remark      string `form:"remark" json:"remark" binding:"required"`
}
