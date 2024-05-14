package saviyntconnectors

import (
	"embed"
)

//go:embed RingCentral/src/*
var RingCentral embed.FS

//go:embed RingCentral/src/Connection_2023-07-26_07-08-58(UTC).json
var RingCentralConnection []byte
