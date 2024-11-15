package codec

import (
	"encoding/hex"

	"github.com/google/uuid"
)

func uuidHex(u uuid.UUID) string {
	return hex.EncodeToString(u[:])
}

func UUIDHex() string {
	return uuidHex(uuid.New())
}
