package utils

import (
	"crypto/md5"
	"fmt"
	"time"
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
