package etronics

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"
)

func TestParseConsumerDevices(t *testing.T) {
	_, exec, _, ok := runtime.Caller(0)

	if !ok {
		t.FailNow()
	}

	path := filepath.Join(filepath.Dir(exec), "/testdata/responses/devices.json")

	devices, err := ioutil.ReadFile(path)

	if err != nil {
		t.Fatal(err)
	}

	dev := ParseConsumerDevices(devices)

	for v := range dev {
		if v.error != nil {
			t.Fatal(v.error)
		}
	}
}
