// +build integration

package etronics

import (
	"encoding/json"
	"flag"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/vitpelekhaty/httptracer"
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
	t.Log("validating flags...")

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

	t.Log("preparing http client...")

	client := &http.Client{Timeout: timeout}

	if strings.TrimSpace(tracePath) != "" {
		tf, err := os.Create(tracePath)

		if err != nil {
			t.Fatal(err)
		}

		defer func() {
			tf.WriteString("]")
			tf.Close()
		}()

		_, err = tf.WriteString("[")

		if err != nil {
			t.Fatal(err)
		}

		emptyTrace := true

		client = httptracer.Trace(client, httptracer.WithBodies(true),
			httptracer.WithCallback(func(entry *httptracer.Entry) {
				if entry == nil {
					return
				}

				b, err := json.Marshal(entry)

				if err != nil {
					t.Fatal(err)
				}

				if !emptyTrace {
					_, err = tf.WriteString(",")

					if err != nil {
						t.Fatal(err)
					}
				}

				_, err = tf.Write(b)

				if err != nil {
					t.Fatal(err)
				}

				emptyTrace = false
			}))
	}

	conn, err := NewConnection(client)

	if err != nil {
		t.Fatal(err)
	}

	t.Log("Logging in...")

	err = conn.Open(rawURL, user, password)

	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		t.Log("disconnecting...")

		conn.Close()
	}()

	t.Log("Reading a list of devices...")

	d, err := conn.ConsumerDevices()

	if err != nil {
		t.Fatal(err)
	}

	devchan := ParseConsumerDevices(d)

	var devCount uint

	t.Log("Reading archives...")

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
