package etronics

import (
	"net/url"
)

const (
	methodGetJWTToken        = `api/v1/GetJWTToken`
	methodGetConsumerDevices = `QueryReport/GetConsumerDevices`
	methodGetArchiveJson     = `Reports/GetArchiveJson`
)

func (c *Connection) login(URL url.URL, username, password string) (string, error) {
	return "", nil
}

func (c *Connection) consumerDevices(URL url.URL) ([]byte, error) {
	return c.get(URL)
}

func (c *Connection) archive(URL url.URL) ([]byte, error) {
	return c.get(URL)
}

func (c *Connection) get(URL url.URL) ([]byte, error) {
	return nil, nil
}
