package etronics

import (
	"context"
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

func TestParseArchive(t *testing.T) {
	_, exec, _, ok := runtime.Caller(0)

	if !ok {
		t.FailNow()
	}

	path := filepath.Join(filepath.Dir(exec), "/testdata/responses/archive.json")

	archive, err := ioutil.ReadFile(path)

	if err != nil {
		t.Fatal(err)
	}

	a := ParseArchive(archive)

	for v := range a {
		if v.error != nil {
			t.Fatal(v.error)
		}
	}
}

func TestParseArchiveWithCancel(t *testing.T) {
	_, exec, _, ok := runtime.Caller(0)

	if !ok {
		t.FailNow()
	}

	path := filepath.Join(filepath.Dir(exec), "/testdata/responses/archive.json")
	const expectedCount = 63

	archive, err := ioutil.ReadFile(path)

	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	a := ParseArchiveWithContext(ctx, archive)

	var count int

	for v := range a {
		if v.error != nil {
			t.Fatal(v.error)
		}

		count++

		cancel()
	}

	t.Logf("read items: %d...", count)

	if !(count < expectedCount) {
		t.Errorf("%s: %d step(s), you cannot stop me", path, count)
	}
}
