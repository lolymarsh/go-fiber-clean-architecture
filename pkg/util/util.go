package util

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

func NewUUID(prefix string) string {
	return strings.ToUpper(prefix) + "_" + uuid.New().String()
}

func NewMiliSecondEpoch() int64 {
	return time.Now().UnixNano()
}
