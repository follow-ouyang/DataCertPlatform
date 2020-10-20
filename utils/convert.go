package utils

import (
	"bytes"
	"encoding/binary"
)

/**
将一个int64类型转换为[]byte字节切片
*/
func Int64ToByte(num int64) ([]byte, error) {
	//Buffer：缓冲区。增益
	buff := new(bytes.Buffer) //通过new实例化一个缓冲区
	//buff.write  通过一系列的write方法向缓冲区写入数据
	//buff.Bytes()  通过Bytes方法从缓冲区中获取数据
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		return nil, err
	}
	//从缓冲区当中读取数据
	return buff.Bytes(), nil

}

/**
将字符串转换为[]byte字节切片
 */
func StringToBytes(data string) []byte {
	return []byte(data)
}