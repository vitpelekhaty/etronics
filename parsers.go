package etronics

import (
	"bytes"
	"encoding/json"

	"github.com/guregu/null"
)

// Device прибора учета тепловой энергии
type Device struct {
	// ID идентификатор прибора учета в АИСКУТЭ
	ID int `json:"DeviceId"`
	// Address адерс места установки прибора учета
	Address string `json:"Address"`
	// Serial серийный номер прибора учета
	Serial string `json:"SerialNumber"`
	// Type тип прибора учета
	Type string `json:"DeviceTypeName"`
	// Vendor производитель прибора учета
	Vendor string `json:"VendorName"`
	// Inputs тепловые вводы прибора учета
	Inputs []byte `json:"Vvods"`
}

// Archive строка архива показаний прибора учета тепловой энергии
type Archive struct {
	// Device прибор учета
	Device string `json:"name"`
	// Serial серийный номер прибора учета
	Serial string `json:"sn"`
	// Type тип архива
	Type DataArchive `json:"archiveType"`
	// Input номер теплового ввода теплового канала
	Input byte `json:"inputNum"`
	// Channel номер канала прибора учета
	Channel byte `json:"channelNum"`
	// Time время полученных показаний прибора учета
	Time ArchiveTime `json:"dt"`
	// M масса теплоносителя по трубе
	M null.Float `json:"M,omitempty"`
	// V объем теплоносителя по трубе
	V null.Float `json:"V,omitempty"`
	// P давление в трубе
	P null.Float `json:"P,omitempty"`
	// T температура теплоносителя по трубе
	T null.Float `json:"T,omitempty"`
	// Ti время наработки прибора учета в часах
	Ti null.Float `json:"Ti,omitempty"`
	// Thw температура холодной воды
	Thw null.Float `json:"Txv,omitempty"`
	// Q тепловая энергия по всему вводу
	Q null.Float `json:"Q,omitempty"`
	// Q1 тепловая энергия по отоплению
	Q1 null.Float `json:"Q1,omitempty"`
	// Q2 тепловая энергия по ГВС
	Q2 null.Float `json:"Q2,omitempty"`
	// BadRow признак отсутствия записи в приборе учета
	BadRow bool `json:"is_bad_row"`
}

// ParseConsumerDevices возвращает результат разбора ответа при вызове метода GetConsumerDevices API
func ParseConsumerDevices(data []byte) <-chan struct {
	*Device
	error
} {
	out := make(chan struct {
		*Device
		error
	})

	go func() {
		defer close(out)

		decoder := json.NewDecoder(bytes.NewReader(data))

		_, err := decoder.Token()

		if err != nil {
			out <- struct {
				*Device
				error
			}{nil, err}

			return
		}

		for decoder.More() {
			var device Device

			err := decoder.Decode(&device)

			if err != nil {
				out <- struct {
					*Device
					error
				}{nil, err}

				break
			}

			out <- struct {
				*Device
				error
			}{&device, nil}
		}
	}()

	return out
}

// ParseArchive возвращает результат разбора ответа при вызове метода GetArchiveJson API
func ParseArchive(data []byte) <-chan struct {
	*Archive
	error
} {
	out := make(chan struct {
		*Archive
		error
	})

	go func() {
		defer close(out)

		decoder := json.NewDecoder(bytes.NewReader(data))

		_, err := decoder.Token()

		if err != nil {
			out <- struct {
				*Archive
				error
			}{nil, err}

			return
		}

		for decoder.More() {
			var archive Archive

			err := decoder.Decode(&archive)

			if err != nil {
				out <- struct {
					*Archive
					error
				}{nil, err}

				break
			}

			out <- struct {
				*Archive
				error
			}{&archive, nil}
		}
	}()

	return out
}
