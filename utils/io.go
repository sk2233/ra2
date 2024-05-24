/*
@author: sk
@date: 2023/9/2
*/
package utils

import (
	"encoding/binary"
	"io"
	"os"
)

func ReadUin32(reader io.Reader) uint32 {
	bs := ReadBytes(reader, 4)
	return binary.LittleEndian.Uint32(bs)
}

func ReadBytes(reader io.Reader, count int) []byte {
	bs := make([]byte, count)
	if num, err := reader.Read(bs); num == 0 {
		// 最后一次读取后会返回 io.EOF 错误，但这次读取的数据是正常的，需要额外判断 读取数目是否为 0
		HandleErr(err)
	}
	return bs
}

func OpenFile(path string) *os.File {
	open, err := os.Open(path)
	HandleErr(err)
	return open
}
