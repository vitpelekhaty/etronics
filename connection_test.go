// +build integration

package etronics

import (
	"flag"
	"net/http"
	"strings"
	"testing"
	"time"
)

var (
	user           string
	password       string
	rawURL         string
	limit          uint
	tracePath      string
	strStart       string
	strEnd         string
	strTimeout     string
	strDataArchive string
)

func init() {
	flag.StringVar(&user, "user", "", "username or email")
	flag.StringVar(&password, "pwd", "", "password")
	flag.StringVar(&rawURL, "url", "", "API url")
	flag.StringVar(&tracePath, "trace", "", "path to a tracing file")
	flag.StringVar(&strStart, "from", "", "a beginning of a measurement period")
	flag.StringVar(&strEnd, "to", "", "end of measurement period")
	flag.StringVar(&strTimeout, "timeout", "30s", "request timeout")
	flag.StringVar(&strDataArchive, "archive", "HourArchive", "type of archive")

	flag.UintVar(&limit, "limit", 0, "limit a number of devices")
}

func TestConnection(t *testing.T) {
	if strings.TrimSpace(rawURL) == "" {
		t.Fatal("no API url")
	}

	start, err := time.Parse(layoutQuery, strStart)

	if err != nil {
		t.Fatal(err)
	}

	end, err := time.Parse(layoutQuery, strEnd)

	if err != nil {
		t.Fatal(err)
	}

	timeout, err := time.ParseDuration(strTimeout)

	if err != nil {
		t.Fatal(err)
	}

	archiveType, err := ParseDataArchive(strDataArchive)

	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{Timeout: timeout}

	conn, err := NewConnection(client)

	if err != nil {
		t.Fatal(err)
	}

	err = conn.Open(rawURL, user, password)

	if err != nil {
		t.Fatal(err)
	}

	defer conn.Close()

	d, err := conn.ConsumerDevices()

	if err != nil {
		t.Fatal(err)
	}

	devchan := ParseConsumerDevices(d)

	var devCount uint

	for device := range devchan {
		if device.error != nil {
			t.Fatal(device.error)
		}

		if limit > 0 && devCount > limit {
			continue
		}

		a, err := conn.Archive(device.ID, archiveType, start, end)

		if err != nil {
			t.Fatal(err)
		}

		archan := ParseArchive(a)

		for archive := range archan {
			if archive.error != nil {
				t.Fatalf("error reading archive of the device %d: %v", device.ID, archive.error)
			}
		}

		devCount++
	}

}
