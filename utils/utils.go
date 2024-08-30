package utils

import (
	"crypto/md5"
	"encoding"
	"fmt"
	"io"
	"strconv"
	"time"

	"heavenMs-NAP-golang/global"
)

func Retry(times int, wait time.Duration, fn func() error) error {
	if times <= 0 {
		times = 1
	}
	for i := 0; i < times; i++ {
		if err := fn(); err == nil {
			return nil
		} else if i == times-1 {
			return err
		}
		time.Sleep(wait)
	}

	return nil
}

func CalcStringMD5(str string) string {
	if str == "" {
		return ""
	}
	data := []byte(str)
	b := md5.Sum(data)
	return fmt.Sprintf("%x", b)
}

// InputToBytes 将输入转换成字节
func InputToBytes(input any) ([]byte, error) {
	var (
		data = make([]byte, 0)
		err  error
	)
	switch d := input.(type) {
	case []byte:
		data = d
	case *[]byte:
		data = *d
	case string:
		data = []byte(d)
	case *string:
		data = []byte(*d)
	case int:
		data = strconv.AppendInt(data, int64(d), 10)
	case *int:
		data = strconv.AppendInt(data, int64(*d), 10)
	case int8:
		data = strconv.AppendInt(data, int64(d), 10)
	case *int8:
		data = strconv.AppendInt(data, int64(*d), 10)
	case int16:
		data = strconv.AppendInt(data, int64(d), 10)
	case *int16:
		data = strconv.AppendInt(data, int64(*d), 10)
	case int32:
		data = strconv.AppendInt(data, int64(d), 10)
	case *int32:
		data = strconv.AppendInt(data, int64(*d), 10)
	case int64:
		data = strconv.AppendInt(data, d, 10)
	case *int64:
		data = strconv.AppendInt(data, *d, 10)
	case uint:
		data = strconv.AppendUint(data, uint64(d), 10)
	case *uint:
		data = strconv.AppendUint(data, uint64(*d), 10)
	case uint8:
		data = strconv.AppendUint(data, uint64(d), 10)
	case *uint8:
		data = strconv.AppendUint(data, uint64(*d), 10)
	case uint16:
		data = strconv.AppendUint(data, uint64(d), 10)
	case *uint16:
		data = strconv.AppendUint(data, uint64(*d), 10)
	case uint32:
		data = strconv.AppendUint(data, uint64(d), 10)
	case *uint32:
		data = strconv.AppendUint(data, uint64(*d), 10)
	case uint64:
		data = strconv.AppendUint(data, d, 10)
	case *uint64:
		data = strconv.AppendUint(data, *d, 10)
	case float32:
		data = strconv.AppendFloat(data, float64(d), 'f', -1, 32)
	case *float32:
		data = strconv.AppendFloat(data, float64(*d), 'f', -1, 32)
	case float64:
		data = strconv.AppendFloat(data, d, 'f', -1, 64)
	case *float64:
		data = strconv.AppendFloat(data, *d, 'f', -1, 64)
	case io.Reader:
		data, err = io.ReadAll(d)
	case encoding.BinaryMarshaler:
		data, err = d.MarshalBinary()
	default:
		err = global.ErrInput
	}
	return data, err
}
