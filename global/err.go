package global

import "errors"

var (
	InvalidInput = errors.New("invalid input")

	InvalidEncodeInput   = errors.New("ProtocolEncoder: invalid input")
	InvalidDecoderOutput = errors.New("ProtocolDecoder: invalid output")

	SmallerLen = errors.New("MapleCypher: crypt error: output smaller than input")

	InvalidPacket = errors.New("MapleDecoder: invalid packet")
)
