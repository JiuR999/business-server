package Time

import (
	"database/sql/driver"
	"strings"
	"time"
)

const (
	LOCAL_DAY_FORMAT = "2006-01-02"
	NIL_TIME         = "0001-01-01 00:00:00"
	NIL_DAY          = "0001-01-01"
)

type LocalDay time.Time

func (t *LocalDay) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*t = LocalDay(time.Time{})
		return
	}
	now, err := time.Parse(`"`+LOCAL_DAY_FORMAT+`"`, string(data))
	*t = LocalDay(now)
	return
}

func (t *LocalDay) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(LOCAL_DAY_FORMAT)+2)
	b = append(b, '"')
	b = time.Time(*t).AppendFormat(b, LOCAL_DAY_FORMAT)
	b = append(b, '"')
	return b, nil
}

func (t *LocalDay) Value() (driver.Value, error) {
	if t == nil || t.String() == NIL_DAY {
		return nil, nil
	}
	return []byte(time.Time(*t).Format(LOCAL_DAY_FORMAT)), nil
}

func (t *LocalDay) Scan(v interface{}) error {
	value := v.(time.Time)
	if strings.Contains(value.String(), NIL_TIME) {
		*t = LocalDay(time.Time{})
		return nil
	}
	tTime, _ := time.Parse(`2006-01-02 15:04:05 +0800 CST`, value.String())
	*t = LocalDay(tTime)
	return nil
}

func (t LocalDay) String() string {
	return time.Time(t).Format(LOCAL_DAY_FORMAT)
}
