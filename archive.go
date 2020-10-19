package etronics

import (
	"fmt"
	"strconv"
	"strings"
)

// UnmarshalJSON реализация интерфейса Unmarshaler для типа DataArchive
func (a *DataArchive) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)

	var i int
	i, err = strconv.Atoi(s)

	if err != nil {
		*a = UnknownArchive
		return
	}

	switch i {
	case int(HourArchive):
		*a = HourArchive
	case int(DailyArchive):
		*a = DailyArchive
	default:
		*a = UnknownArchive
		err = fmt.Errorf("unknown archive type %d", i)
	}

	return
}
