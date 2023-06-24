package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildDirNameForDevice(t *testing.T) {
	name := buildDirNameForDevice(&adb.DeviceInfo{
		Model:      "foo1",
		Serial:     "bar2",
		Product:    "ignored",
		Usb:        "ignored",
		DeviceInfo: "ignored",
	})
	assert.Equal(t, "foo1-bar2", name)

	name = buildDirNameForDevice(&adb.DeviceInfo{
		Model:  "-f-o-o_!@#$",
		Serial: "bar%^&*()",
	})
	assert.Equal(t, "-f-o-o_-bar_", name)
}
