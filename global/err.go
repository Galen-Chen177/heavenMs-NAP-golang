package global

import "errors"

var (
	ErrInput         = errors.New("invalid input")
	ErrEncodeInput   = errors.New("ProtocolEncoder: invalid input")
	ErrDecoderOutput = errors.New("ProtocolDecoder: invalid output")
	ErrSmallerLen    = errors.New("MapleCypher: crypt error: output smaller than input")
	ErrPacket        = errors.New("MapleDecoder: invalid packet")
)
