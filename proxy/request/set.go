package request

import "time"

// SetRequest the request of the set command
type SetRequest struct {
	Value      any           `json:"value"`
	Expiration time.Duration `json:"expiration"`
}
