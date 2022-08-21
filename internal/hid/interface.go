package hid

import (
	"context"
	"time"
)

type Device interface {
	Open(ctx context.Context) error  // Open file handle
	Close(ctx context.Context) error // Close file handle
	Info() DeviceInfo
	SendFeatureReport(ctx context.Context, v []byte) (int, error) // ioctl set report
	GetFeatureReport(ctx context.Context, v []byte) (int, error)  // ioctl get report
	Read(ctx context.Context, v []byte, t time.Duration) (int, error)
	Write(ctx context.Context, v []byte) (int, error)
}

type DeviceInfo struct {
	VendorID  uint16
	ProductID uint16
	Revision  uint16

	SubClass uint8
	Protocol uint8

	Interface uint8
	Bus       int
	Device    int
}
