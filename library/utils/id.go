package utils

import (
	"sync"

	"github.com/sony/sonyflake"
)

var (
	flake     *sonyflake.Sonyflake
	flakeOnce sync.Once
)

func getSonyFlake() *sonyflake.Sonyflake {
	flakeOnce.Do(
		func() {
			flake = sonyflake.NewSonyflake(sonyflake.Settings{})
		},
	)
	return flake
}

func RandomID() (uint64, error) {
	id, err := getSonyFlake().NextID()
	if err != nil {
		return 0, err
	}
	return id, nil
}
