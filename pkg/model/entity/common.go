package entity

import (
	"5z7Game/pkg/constant"
	"database/sql/driver"
	"fmt"
	"time"
)

type BaseModel struct {
	ID       int    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	// 对time进行格式化
	CreatedAt Timestamp `gorm:"column:created_at" json:"created_at"`
	UpdatedAt Timestamp `gorm:"column:created_at" json:"updated_at"`
}

// Timestamp 自定义gorm的时间戳格式
type Timestamp struct {
	time.Time
}

// MarshalJSON 解析
func (t Timestamp) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format(constant.TimeFormat))
	return []byte(formatted), nil
}

// Value
func (t Timestamp) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan 转换时间戳
func (t *Timestamp) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Timestamp{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t *Timestamp) String() string {
	return t.Format(constant.TimeFormat)
}

// Columns 定义更新用的字段,其实就是个map
type Columns map[string]interface{}