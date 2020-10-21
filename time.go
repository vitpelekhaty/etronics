package etronics

import (
	"fmt"
	"strings"
	"time"
)

// QueryTime представление времени в запросе к API АИСКУТЭ Энерготроника
type QueryTime time.Time

// ArchiveTime представление времени в архиве показаний приборов учета
type ArchiveTime time.Time

const (
	layoutQuery   = `02.01.2006 03`
	layoutArchive = `2006-01-02T15:04:05`
)

// UnmarshalJSON реализация интерфейса Unmarshaler для типа QueryTime
func (t *QueryTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	v, err := time.Parse(layoutQuery, s)

	*t = QueryTime(v)

	return
}

// MarshalJSON реализация интерфейса Marshaller для типа QueryTime
func (t QueryTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.String())), nil
}

// String возвращает строковое представление значения типа QueryTime
func (t *QueryTime) String() string {
	v := time.Time(*t)
	return v.Format(layoutQuery)
}

// UnmarshalJSON реализация интерфейса Unmarshaler для типа ArchiveTime
func (t *ArchiveTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	v, err := time.Parse(layoutArchive, s)

	*t = ArchiveTime(v)

	return
}

// MarshalJSON реализация интерфейса Marshaller для типа ArchiveTime
func (t ArchiveTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.String())), nil
}

// String возвращает строковое представление значения типа ArchiveTime
func (t *ArchiveTime) String() string {
	v := time.Time(*t)
	return v.Format(layoutArchive)
}
