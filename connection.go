package etronics

import (
	"net/http"
	"time"
)

// DataArchive тип архива показаний
type DataArchive byte

const (
	// UnknownArchive неизвестный тип архива
	UnknownArchive DataArchive = 0
	// HourArchive часовой архив
	HourArchive DataArchive = 1
	// DailyArchive суточный архив
	DailyArchive DataArchive = 2
)

// Connection соединение с API АИСКУТЭ Энерготроника
type Connection struct {
	client *http.Client

	baseURL string
	token   string
}

// NewConnection возвращает настроенный экземпляр соединения с API Энерготроника
func NewConnection(client *http.Client) (*Connection, error) {
	return nil, nil
}

// Open устанавливает соединение с API и авторизуется в API
//
// Параметры:
//
//   URL - адрес API
//   username - имя пользователя
//   password - пароль пользователя
func (c *Connection) Open(URL string, username, password string) error {
	c.Close()

	return nil
}

// Close закрывает соединение с API
func (c *Connection) Close() {
	c.token = ""
}

// Connected возвращает true, если установлено соединение с API и выполнена авторизация
func (c *Connection) Connected() bool {
	return c.token != ""
}

// ConsumerDevices возвращает результат обращения к методу GetConsumerDevices API.
// Для разбора ответа используется функция ParseConsumerDevices
func (c *Connection) ConsumerDevices() ([]byte, error) {
	return nil, nil
}

// Archive возвращает результат обращения к методу GetArchiveJson API
//
// Параметры:
//   deviceID - идентификатор прибора учета тепловой энергии в АИСКУТЭ Энерготроника
//   archive - тип возвращаемого архива показаний
//   start, end - период показаний
//
// Для разбора ответа используется функция ParseArchive
func (c *Connection) Archive(deviceID int, archive DataArchive, start, end time.Time) ([]byte, error) {
	return nil, nil
}
