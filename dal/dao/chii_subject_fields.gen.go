// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"time"

	"github.com/bangumi/server/internal/model"
)

const TableNameSubjectField = "chii_subject_fields"

// SubjectField mapped from table <chii_subject_fields>
type SubjectField struct {
	Sid      model.SubjectID `gorm:"column:field_sid;type:mediumint(8) unsigned;primaryKey;autoIncrement:true"`
	Tid      uint16          `gorm:"column:field_tid;type:smallint(6) unsigned;not null"`
	Tags     []byte          `gorm:"column:field_tags;type:mediumtext;not null"`
	Rate1    uint32          `gorm:"column:field_rate_1;type:mediumint(8) unsigned;not null"`
	Rate2    uint32          `gorm:"column:field_rate_2;type:mediumint(8) unsigned;not null"`
	Rate3    uint32          `gorm:"column:field_rate_3;type:mediumint(8) unsigned;not null"`
	Rate4    uint32          `gorm:"column:field_rate_4;type:mediumint(8) unsigned;not null"`
	Rate5    uint32          `gorm:"column:field_rate_5;type:mediumint(8) unsigned;not null"`
	Rate6    uint32          `gorm:"column:field_rate_6;type:mediumint(8) unsigned;not null"`
	Rate7    uint32          `gorm:"column:field_rate_7;type:mediumint(8) unsigned;not null"`
	Rate8    uint32          `gorm:"column:field_rate_8;type:mediumint(8) unsigned;not null"`
	Rate9    uint32          `gorm:"column:field_rate_9;type:mediumint(8) unsigned;not null"`
	Rate10   uint32          `gorm:"column:field_rate_10;type:mediumint(8) unsigned;not null"`
	Airtime  uint8           `gorm:"column:field_airtime;type:tinyint(1) unsigned;not null"`
	Rank     uint32          `gorm:"column:field_rank;type:int(10) unsigned;not null"`
	Year     int32           `gorm:"column:field_year;type:year(4);not null"`        // 放送年份
	Mon      int8            `gorm:"column:field_mon;type:tinyint(2);not null"`      // 放送月份
	WeekDay  int8            `gorm:"column:field_week_day;type:tinyint(1);not null"` // 放送日(星期X)
	Date     time.Time       `gorm:"column:field_date;type:date;not null"`           // 放送日期
	Redirect model.SubjectID `gorm:"column:field_redirect;type:mediumint(8) unsigned;not null"`
}

// TableName SubjectField's table name
func (*SubjectField) TableName() string {
	return TableNameSubjectField
}