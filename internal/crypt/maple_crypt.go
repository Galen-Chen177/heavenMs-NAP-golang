package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"

	"heavenMs-NAP-golang/global"
	"heavenMs-NAP-golang/utils"
)

var funnyBytes = []byte{0xEC, 0x3F, 0x77, 0xA4, 0x45, 0xD0,
	0x71, 0xBF, 0xB7, 0x98, 0x20, 0xFC, 0x4B, 0xE9, 0xB3, 0xE1,
	0x5C, 0x22, 0xF7, 0x0C, 0x44, 0x1B, 0x81, 0xBD, 0x63, 0x8D, 0xD4,
	0xC3, 0xF2, 0x10, 0x19, 0xE0, 0xFB, 0xA1, 0x6E, 0x66, 0xEA,
	0xAE, 0xD6, 0xCE, 0x06, 0x18, 0x4E, 0xEB, 0x78, 0x95, 0xDB,
	0xBA, 0xB6, 0x42, 0x7A, 0x2A, 0x83, 0x0B, 0x54, 0x67, 0x6D, 0xE8, 0x65,
	0xE7, 0x2F, 0x07, 0xF3, 0xAA, 0x27, 0x7B, 0x85, 0xB0, 0x26, 0xFD,
	0x8B, 0xA9, 0xFA, 0xBE, 0xA8, 0xD7, 0xCB, 0xCC,
	0x92, 0xDA, 0xF9, 0x93, 0x60, 0x2D, 0xDD, 0xD2, 0xA2,
	0x9B, 0x39, 0x5F, 0x82, 0x21, 0x4C, 0x69, 0xF8, 0x31, 0x87, 0xEE,
	0x8E, 0xAD, 0x8C, 0x6A, 0xBC, 0xB5, 0x6B, 0x59, 0x13, 0xF1, 0x04,
	0x00, 0xF6, 0x5A, 0x35, 0x79, 0x48, 0x8F, 0x15, 0xCD, 0x97, 0x57, 0x12, 0x3E, 0x37,
	0xFF, 0x9D, 0x4F, 0x51, 0xF5, 0xA3, 0x70, 0xBB, 0x14, 0x75, 0xC2,
	0xB8, 0x72, 0xC0, 0xED, 0x7D, 0x68, 0xC9, 0x2E, 0x0D, 0x62, 0x46, 0x17, 0x11, 0x4D,
	0x6C, 0xC4, 0x7E, 0x53, 0xC1, 0x25, 0xC7, 0x9A, 0x1C, 0x88, 0x58, 0x2C,
	0x89, 0xDC, 0x02, 0x64, 0x40, 0x01, 0x5D, 0x38, 0xA5, 0xE2, 0xAF, 0x55,
	0xD5, 0xEF, 0x1A, 0x7C, 0xA7, 0x5B, 0xA6, 0x6F, 0x86, 0x9F, 0x73,
	0xE6, 0x0A, 0xDE, 0x2B, 0x99, 0x4A, 0x47, 0x9C, 0xDF, 0x09, 0x76,
	0x9E, 0x30, 0x0E, 0xE4, 0xB2, 0x94, 0xA0, 0x3B, 0x34, 0x1D, 0x28, 0x0F,
	0x36, 0xE3, 0x23, 0xB4, 0x03, 0xD8, 0x90, 0xC8, 0x3C, 0xFE, 0x5E,
	0x32, 0x24, 0x50, 0x1F, 0x3A, 0x43, 0x8A, 0x96, 0x41, 0x74, 0xAC, 0x52, 0x33, 0xF0,
	0xD9, 0x29, 0x80, 0xB1, 0x16, 0xD3, 0xAB, 0x91, 0xB9,
	0x84, 0x7F, 0x61, 0x1E, 0xCF, 0xC5, 0xD1, 0x56, 0x3D, 0xCA, 0xF4,
	0x05, 0xC6, 0xE5, 0x08, 0x49,
}

// NewMapleCypher 返回一个枫叶传说专用传输加密器
func NewMapleCypher(key, iv []byte, mapleVersion uint16) (*MapleCypher, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	oriIv := make([]byte, len(iv))
	copy(oriIv, iv)
	return &MapleCypher{
		iv:           iv,
		oriIv:        oriIv,
		cipher:       block,
		mapleVersion: (mapleVersion>>8)&0xFF | (mapleVersion<<8)&0xFF00,
	}, nil
}

// MapleCypher 枫叶传说专用传输加密器
type MapleCypher struct {
	iv           []byte
	oriIv        []byte
	cipher       cipher.Block
	mapleVersion uint16
}

// OriginalIv 返回该暗号的初始向量
func (c *MapleCypher) OriginalIv() []byte {
	return c.oriIv
}

// Crypt 进行加密
func (c *MapleCypher) Crypt(dst, src []byte) error {
	if len(dst) < len(src) {
		return global.ErrSmallerLen
	}

	remaining := len(src)
	llength := 0x5B0
	start := 0

	for remaining > 0 {
		myIv := c.multiplyBytes(c.iv[:4])
		if remaining < llength {
			llength = remaining
		}
		for x := start; x < (start + llength); x++ {
			if (x-start)%len(myIv) == 0 {
				newIv := make([]byte, len(myIv))
				c.cipher.Encrypt(newIv, myIv)
				copy(myIv, newIv)
			}
			dst[x] = src[x] ^ myIv[(x-start)%len(myIv)]
		}
		start += llength
		remaining -= llength
		llength = 0x5B4
		c.updateIv()
	}
	return nil
}

func (c *MapleCypher) updateIv() {
	c.iv = c.getNewIv(c.iv)
}

// PacketHeader 根据数据长度,生成包头部信息
func (c *MapleCypher) PacketHeader(length int) []byte {
	iiv := (uint16(c.iv[3])&0xFF | (uint16(c.iv[2])<<8)&0xFF00) ^ c.mapleVersion
	mlength := (uint16(length<<8) & 0xFF00) | uint16(length>>8) ^ iiv

	packetHeader := make([]byte, 4)
	binary.BigEndian.PutUint16(packetHeader, iiv)
	binary.BigEndian.PutUint16(packetHeader[2:], mlength)

	return packetHeader
}

func (c *MapleCypher) multiplyBytes(bytes []byte) []byte {
	newBytes := make([]byte, len(bytes)*4)
	for i := 0; i < 4; i++ {
		copy(newBytes[i*len(bytes):(i+1)*len(bytes)], bytes)
	}
	return newBytes
}

func (c *MapleCypher) getNewIv(oldIv []byte) []byte {
	newIv := []byte{0xf2, 0x53, 0x50, 0xc6}
	for x := 0; x < 4; x++ {
		funnyShit(oldIv[x], newIv)
	}
	return newIv
}

// PacketLength 获取包长度
// @param header 头部信息
func (c *MapleCypher) PacketLength(header int) int {
	packetLength := (header >> 16) ^ (header & 0xFFFF)
	packetLength = (packetLength << 8 & 0xFF00) | (packetLength >> 8 & 0xFF)
	return packetLength
}

// CheckPacketHeader 检测包头是否合法
func (c *MapleCypher) CheckPacketHeader(header int) bool {
	packetLength := (header >> 16) ^ (header & 0xFFFF)
	packetLength = ((packetLength << 8) & 0xFF00) | ((packetLength >> 8) & 0xFF)
	return c.CheckHashHeader([]byte{byte(header >> 24), byte(header >> 16)})

}

// CheckHashHeader 检测该消息是否包含头信息
func (c *MapleCypher) CheckHashHeader(packet []byte) bool {
	return (packet[0]^c.iv[2]&0xFF) == (byte(c.mapleVersion>>8)&0xFF) && (packet[1]^c.iv[3]&0xFF) == (byte(c.mapleVersion)&0xFF)
}

func funnyShit(inputByte byte, in []byte) {
	elina := in[1]
	anna := inputByte
	moritz := funnyBytes[elina&0xFF]
	moritz -= inputByte
	in[0] += moritz
	moritz = in[2]
	moritz ^= funnyBytes[anna&0xFF]
	elina -= moritz & 0xFF
	in[1] = elina
	elina = in[3]
	moritz = elina
	elina -= in[0] & 0xFF
	moritz = funnyBytes[moritz&0xFF]
	moritz += inputByte
	moritz ^= in[2]
	in[2] = moritz
	elina += funnyBytes[anna&0xFF] & 0xFF
	in[3] = elina

	merry := int(in[0]) & 0xFF
	merry |= int(in[1]) << 8 & 0xFF00
	merry |= int(in[2]) << 16 & 0xFF0000
	merry |= int(in[3]) << 24 & 0xFF000000
	retValue := uint32(merry) >> 0x1d
	merry <<= 3
	retValue |= uint32(merry)

	in[0] = byte(retValue & 0xFF)
	in[1] = byte((retValue >> 8) & 0xFF)
	in[2] = byte((retValue >> 16) & 0xFF)
	in[3] = byte((retValue >> 24) & 0xFF)
}

// MapleEncrypt 枫叶传说专用加密方法
func MapleEncrypt(data []byte) {
	for j := 0; j < 6; j++ {
		remember := byte(0)
		dataLength := (byte)(len(data) & 0xFF)
		if j%2 == 0 {
			for i := 0; i < len(data); i++ {
				cur := data[i]
				cur = utils.RollLeft(cur, 3)
				cur += dataLength
				cur ^= remember
				remember = cur
				cur = utils.RollRight(cur, int(dataLength&0xFF))
				cur = (^cur) & 0xFF
				cur += 0x48
				dataLength--
				data[i] = cur
			}
		} else {
			for i := len(data) - 1; i >= 0; i-- {
				cur := data[i]
				cur = utils.RollLeft(cur, 4)
				cur += dataLength
				cur ^= remember
				remember = cur
				cur ^= 0x13
				cur = utils.RollRight(cur, 3)
				dataLength--
				data[i] = cur
			}
		}
	}
}

// MapleDecrypt 枫叶传说专用解密方法
func MapleDecrypt(data []byte) {
	for j := 1; j <= 6; j++ {
		remember := byte(0)
		dataLength := byte(len(data) & 0xFF)
		nextRemember := byte(0)

		if j%2 == 0 {
			for i := 0; i < len(data); i++ {
				cur := data[i]
				cur = cur - 72
				cur = (^cur) & 0xFF
				cur = utils.RollLeft(cur, int(dataLength&0xFF))
				nextRemember = cur
				cur = cur ^ remember
				remember = nextRemember
				cur = cur - dataLength
				cur = utils.RollRight(cur, 3)
				data[i] = cur
				dataLength = dataLength - 1
			}
		} else {
			for i := len(data) - 1; i >= 0; i-- {
				cur := data[i]
				cur = utils.RollLeft(cur, 3)
				cur = cur ^ 0x13
				nextRemember = cur
				cur = cur ^ remember
				remember = nextRemember
				cur = cur - dataLength
				cur = utils.RollRight(cur, 4)
				data[i] = cur
				dataLength = dataLength - 1
			}
		}
	}
}
