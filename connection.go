package etronics

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strconv"
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
	if client == nil {
		return nil, errors.New("undefined HTTP client")
	}

	return &Connection{client: client}, nil
}

// Open устанавливает соединение с API и авторизуется в API
//
// Параметры:
//
//   URL - адрес API
//   username - имя пользователя
//   password - пароль пользователя
func (c *Connection) Open(rawURL string, username, password string) error {
	return c.OpenWithContext(context.Background(), rawURL, username, password)
}

// OpenWithContext устанавливает соединение с API и авторизуется в API
//
// Параметры:
//
//   URL - адрес API
//   username - имя пользователя
//   password - пароль пользователя
func (c *Connection) OpenWithContext(ctx context.Context, rawURL string, username, password string) error {
	c.Close()

	if c.client == nil {
		return errors.New("undefined HTTP client")
	}

	_, err := url.Parse(rawURL)

	if err != nil {
		return err
	}

	c.baseURL = rawURL

	loginURL, err := join(rawURL, methodGetJWTToken)

	if err != nil {
		return err
	}

	u, err := url.Parse(loginURL)

	if err != nil {
		return err
	}

	token, err := c.login(ctx, u, username, password)

	if err != nil {
		return err
	}

	c.token = token

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
	return c.ConsumerDevicesWithContext(context.Background())
}

// ConsumerDevicesWithContext возвращает результат обращения к методу GetConsumerDevices API.
// Для разбора ответа используется функция ParseConsumerDevices
func (c *Connection) ConsumerDevicesWithContext(ctx context.Context) ([]byte, error) {
	if c.client == nil {
		return nil, errors.New("undefined HTTP client")
	}

	if !c.Connected() {
		return nil, errors.New("no connection")
	}

	rawURL, err := join(c.baseURL, methodGetConsumerDevices)

	if err != nil {
		return nil, err
	}

	u, err := url.Parse(rawURL)

	if err != nil {
		return nil, err
	}

	return c.consumerDevices(ctx, u)
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
	return c.ArchiveWithContext(context.Background(), deviceID, archive, start, end)
}

// ArchiveWithContext возвращает результат обращения к методу GetArchiveJson API
//
// Параметры:
//   deviceID - идентификатор прибора учета тепловой энергии в АИСКУТЭ Энерготроника
//   archive - тип возвращаемого архива показаний
//   start, end - период показаний
//
// Для разбора ответа используется функция ParseArchive
func (c *Connection) ArchiveWithContext(ctx context.Context, deviceID int, archive DataArchive, start,
	end time.Time) ([]byte, error) {
	if c.client == nil {
		return nil, errors.New("undefined HTTP client")
	}

	if !c.Connected() {
		return nil, errors.New("no connection")
	}

	rawURL, err := join(c.baseURL, methodGetArchiveJson)

	if err != nil {
		return nil, err
	}

	u, err := url.Parse(rawURL)

	if err != nil {
		return nil, err
	}

	s := QueryTime(start)
	e := QueryTime(end)

	query := url.Values{}

	query.Add("DeviceId", strconv.Itoa(deviceID))
	query.Add("ArchiveType", strconv.Itoa(int(archive)))
	query.Add("DtFrom", s.String())
	query.Add("DtTo", e.String())

	u.RawQuery = query.Encode()

	return c.archive(ctx, u)
}
