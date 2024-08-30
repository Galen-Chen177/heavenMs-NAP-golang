package crypt

import (
	"io"

	"heavenMs-NAP-golang/global"
	"heavenMs-NAP-golang/utils"
)

// ------------------ ProtocolEncoder

// ProtocolEncoder 协议编码器
type ProtocolEncoder interface {
	Encode(input any, output io.Writer) error
}

type UnimplementedProtocolEncoder struct {
}

func (u UnimplementedProtocolEncoder) Encode(input any, output io.Writer) error {
	//TODO implement me

	data, err := utils.InputToBytes(input)
	if err != nil {
		return global.InvalidEncodeInput
	}
	_, err = output.Write(data)
	return err
}

// ------------------ ProtocolDecoder

// ProtocolDecoder 协议解码器
type ProtocolDecoder interface {
	Decode(input io.Reader, output io.Writer) error
}

type UnimplementedProtocolDecoder struct {
}

func (u UnimplementedProtocolDecoder) Decode(input io.Reader, output io.Writer) error {
	_, err := io.Copy(output, input)
	return err
}

// ------------- UnimplementedProtocolCodec

// ProtocolCodecer 协议编码解码器
type ProtocolCodecer interface {
	ProtocolEncoder
	ProtocolDecoder
}

var _ ProtocolCodecer = new(UnimplementedProtocolCodec)

type UnimplementedProtocolCodec struct {
}

func (u UnimplementedProtocolCodec) Encode(input any, output io.Writer) error {
	//TODO implement me
	data, err := utils.InputToBytes(input)
	if err != nil {
		return global.InvalidEncodeInput
	}
	_, err = output.Write(data)
	return err
}

func (u UnimplementedProtocolCodec) Decode(input io.Reader, output io.Writer) error {
	_, err := io.Copy(output, input)
	return err
}
