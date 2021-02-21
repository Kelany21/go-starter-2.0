package helpers

import (
	"github.com/jinzhu/gorm"
)

// Param 分页参数
type Param struct {
	DB      *gorm.DB
	Page    int
	Limit   int
	OrderBy []string
	Filters []string
	Preload []string
	ShowSQL bool
}

// Paginator 分页返回
type Paginator struct {
	TotalRecord  int `json:"total_record"`
	RecordsCount int `json:"records_count"`
	TotalPage    int `json:"total_page"`
	Offset       int `json:"offset"`
	Limit        int `json:"limit"`
	Page         int `json:"page"`
	PrevPage     int `json:"prev_page"`
	NextPage     int `json:"next_page"`
	//Records     interface{} `json:"records"`
}
