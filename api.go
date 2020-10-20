package etronics

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

const (
	methodGetJWTToken        = `api/v1/GetJWTToken`
	methodGetConsumerDevices = `QueryReport/GetConsumerDevices`
	methodGetArchiveJson     = `Reports/GetArchiveJson`
)

func (c *Connection) login(u *url.URL, user, password string) (string, error) {
	form := url.Values{}

	form.Add("Email", user)
	form.Add("Password", password)

	req, err := http.NewRequest("POST", u.String(), strings.NewReader(form.Encode()))

	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	resp, err := c.client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		token, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return "", err
		}

		return string(token), nil
	}

	message, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	if strings.TrimSpace(string(message)) != "" {
		return "", fmt.Errorf("login error %d (%s): %s", resp.StatusCode, resp.Status, string(message))
	}

	return "", fmt.Errorf("login error %d (%s)", resp.StatusCode, resp.Status)
}

func (c *Connection) consumerDevices(u *url.URL) ([]byte, error) {
	return c.get(u)
}

func (c *Connection) archive(u *url.URL) ([]byte, error) {
	return c.get(u)
}

func (c *Connection) get(u *url.URL) ([]byte, error) {
	req, err := http.NewRequest("GET", u.String(), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		return b, nil
	}

	if strings.TrimSpace(string(b)) != "" {
		return nil, fmt.Errorf("%d (%s): %s", resp.StatusCode, resp.Status, string(b))
	}

	return nil, fmt.Errorf("%d (%s)", resp.StatusCode, resp.Status)
}

func join(rawURL string, elem string) (string, error) {
	u, err := url.Parse(rawURL)

	if err != nil {
		return rawURL, err
	}

	u.Path = path.Join(u.Path, elem)

	return u.String(), nil
}
