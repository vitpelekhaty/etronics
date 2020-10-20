package etronics

import (
	"bytes"
	"encoding/json"
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
	Time Time `json:"dt"`
	// M масса теплоносителя по трубе
	M float32 `json:"M"`
	// V объем теплоносителя по трубе
	V float32 `json:"V"`
	// P давление в трубе
	P float32 `json:"P"`
	// T температура теплоносителя по трубе
	T float32 `json:"T"`
	// Ti время наработки прибора учета в часах
	Ti float32 `json:"Ti"`
	// Thw температура холодной воды
	Thw float32 `json:"Txv"`
	// Q тепловая энергия по всему вводу
	Q float32 `json:"Q"`
	// Q1 тепловая энергия по отоплению
	Q1 float32 `json:"Q1"`
	// Q2 тепловая энергия по ГВС
	Q2 float32 `json:"Q2"`
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
func ParseArchive(data []byte) (chan *Archive, error) {
	return nil, nil
}
