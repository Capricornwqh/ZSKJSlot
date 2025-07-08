package entity_pgsql

import (
	"database/sql/driver"
	"errors"

	jsoniter "github.com/json-iterator/go"
)

// JSON类型，用于存储不同操作类型的详细信息
type JSONB map[string]any

// Value 实现driver.Valuer接口
func (j JSONB) Value() (driver.Value, error) {
	return jsoniter.Marshal(j)
}

// Scan 实现sql.Scanner接口
func (j *JSONB) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return jsoniter.Unmarshal(bytes, &j)
}
