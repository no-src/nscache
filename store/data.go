package store

import "time"

// Data the store data object
type Data struct {
	// Data the real user data
	Data []byte
	// ExpireTime the expiration time of the data
	ExpireTime time.Time
}

// NewData create a store data object
func NewData(data []byte, expiration time.Duration) *Data {
	return &Data{
		Data:       data,
		ExpireTime: time.Now().Add(expiration),
	}
}

// IsExpired check if the data is expired or not
func (d *Data) IsExpired() bool {
	return d.ExpireTime.Before(time.Now())
}
