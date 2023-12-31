package payload

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

type Kind uint8

const (
	KindMeta Kind = iota + 1
	KindContent
	KindLayout
	KindIndex
	KindAttributes
	KindDumb
)

type Compression uint8

const (
	CompressionNone Compression = iota + 1
	CompressionZstd
)

type PayloadHeader struct {
	StoredSize  uint64
	PlainSize   uint64
	Checksum    [8]uint8
	NumRecords  uint32
	Version     uint16
	Kind        Kind
	Compression Compression
}

func ReadPayloadHeader(r io.Reader) (PayloadHeader, error) {
	payloadHeaderHeader := PayloadHeader{}
	err := binary.Read(r, binary.BigEndian, &payloadHeaderHeader)
	if err != nil {
		return PayloadHeader{}, err
	}

	return payloadHeaderHeader, nil
}

func (p PayloadHeader) Print() {
	fmt.Printf("Payload: %s [Records: %d Compression: %s, Savings: %.2f%%, Size: %s]\n",
		strings.TrimPrefix(p.Kind.String(), "Kind"), p.NumRecords, strings.TrimPrefix(p.Compression.String(), "Compression"), 100-(float64(p.StoredSize)/float64(p.PlainSize)*100), formatBytes(p.PlainSize))
}
