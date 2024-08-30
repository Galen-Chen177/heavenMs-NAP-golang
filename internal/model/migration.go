package model

import "time"

// id tag name time isSuccess
type Migration struct {
	ID        int `gorm:"primarykey"`
	Tag       int
	Name      string `gorm:"type:varchar(255);not null"`
	Md5       string `gorm:"type:varchar(255);not null"`
	Content   string `gorm:"-"`
	ExecTime  time.Time
	IsSuccess bool   `gorm:"type:bool"`
	Remark    string `gorm:"type:varchar(255);not null"`
}

func NewMigration(tag int, name, md5, cont string) Migration {
	return Migration{
		Tag:     tag,
		Name:    name,
		Md5:     md5,
		Content: cont,
	}
}

type MigrationSlice []Migration

func (m MigrationSlice) Len() int {
	return len(m)
}

func (m MigrationSlice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m MigrationSlice) Less(i, j int) bool {
	return m[i].Tag < m[j].Tag
}
