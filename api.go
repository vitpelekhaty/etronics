package etronics

import (
	"net/url"
	"path"
)

const (
	methodGetJWTToken        = `api/v1/GetJWTToken`
	methodGetConsumerDevices = `QueryReport/GetConsumerDevices`
	methodGetArchiveJson     = `Reports/GetArchiveJson`
)

func (c *Connection) login(u *url.URL, username, password string) (string, error) {
	return "", nil
}

func (c *Connection) consumerDevices(u *url.URL) ([]byte, error) {
	return c.get(u)
}

func (c *Connection) archive(u *url.URL) ([]byte, error) {
	return c.get(u)
}

func (c *Connection) get(u *url.URL) ([]byte, error) {
	return nil, nil
}

func join(rawURL string, elem string) (string, error) {
	u, err := url.Parse(rawURL)

	if err != nil {
		return rawURL, err
	}

	u.Path = path.Join(u.Path, elem)

	return u.String(), nil
}
