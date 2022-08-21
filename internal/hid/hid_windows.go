package hid

import (
	"context"
	"time"

	"git.sr.ht/~errnoh/go.hid/devices/hid"
)

type USB struct {
	*hid.Device
}

func (d *USB) Open(ctx context.Context) error { return nil }
func (d *USB) Close(ctx context.Context) error {
	d.Device.Close()
	return nil
}
func (d *USB) Info() DeviceInfo {
	return DeviceInfo{
		VendorID:  0x0fd9,
		ProductID: 0x6c,
	}
}

func (d *USB) SendFeatureReport(ctx context.Context, v []byte) (int, error) {
	return d.Device.SendFeatureReport(v)
}

func (d *USB) GetFeatureReport(ctx context.Context, v []byte) (int, error) {
	return d.Device.GetFeatureReport(v)
}

func (d *USB) Read(ctx context.Context, v []byte, t time.Duration) (n int, err error) {
	return d.ReadTimeout(v, int(t/time.Millisecond))
}

func (d *USB) Write(ctx context.Context, v []byte) (n int, err error) {
	return d.Device.Write(v)
}

func Devices() (devs []Device, err error) {
	var (
		arr = hid.Devices(0x0fd9, 0, 0, 0)
	)
	if len(arr) > 0 {
		for i := 1; i < len(arr); i++ {
			arr[i].Close()
		}
		devs = []Device{&USB{arr[0]}}
	}

	return
}
