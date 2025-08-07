package entity

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// Конвертация bigint Postgres и time.Duration Go
type Duration time.Duration

func (d Duration) Value() (driver.Value, error) {
	return driver.Value(int64(d)), nil
}

func (d *Duration) Scan(raw interface{}) error {
	switch v := raw.(type) {
	case int64:
		*d = Duration(v)
	case nil:
		*d = Duration(0)
	default:
		return fmt.Errorf("cannot sql.Scan() Duration from: %#v", v)
	}
	return nil
}
