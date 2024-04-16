package data

import (
	"database/sql"
	"time"
)

type ModuleInfo struct {
	ID             int64     `json:"id"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
	ModuleName     string    `json:"moduleName"`
	ModuleDuration int32     `json:"year,omitempty"`
	Runtime        Runtime   `json:"runtime,omitempty,string"`
	ExamType       []string  `json:"examType"`
	Version        int32     `json:"version"`
}
type ModuleInfoModel struct {
	DB *sql.DB
}
