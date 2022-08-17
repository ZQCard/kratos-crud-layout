package entity

import (
	"gorm.io/gorm"
	"time"
)

const (
	// ServiceNameDeleted 已经删除
	ServiceNameDeleted = "1"
	// ServiceNameUnDeleted 未删除
	ServiceNameUnDeleted = "2"
)

type ServiceNameEntity struct {
	Id        int64
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt string
}

func (ServiceNameEntity) TableName() string {
	return "serviceName"
}

// Paginate 分页
func Paginate(page, pageSize int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}

// UnDelete 非删除数据
func UnDelete() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("deleted_at = ''")
	}
}

func HasDelete() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("deleted_at != ''")
	}
}
