package etronics

import (
	"fmt"
	"strings"
	"time"
)

// Time представление времени в АИСКУТЭ Энерготроника
type Time time.Time

const layout = `02.01.2006 15:04:05`

// UnmarshalJSON реализация интерфейса Unmarshaler для типа Time
func (t *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	v, err := time.Parse(layout, s)

	*t = Time(v)

	return
}

// MarshalJSON реализация интерфейса Marshaller для типа Time
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.String())), nil
}

// String возвращает строковое представление значения типа Time
func (t *Time) String() string {
	v := time.Time(*t)
	return v.Format(layout)
}
