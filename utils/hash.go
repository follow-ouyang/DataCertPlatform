package utils

import (
	"crypto/md5"
	sha2562 "crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
)

/*
对一个字符串进行hash计算
 */
func Md5HashString(data string) string {
	md5Hash := md5.New()
	md5Hash.Write([]byte(data))
	bytes := md5Hash.Sum(nil)
	return hex.EncodeToString(bytes)
}

//io：input output的缩写
func Md5HashReader(reader io.Reader) (string,error) {
	md5Hash := md5.New()
	readerBytes,err := ioutil.ReadAll(reader)
	if err != nil {
		return "",err
	}
	md5Hash.Write(readerBytes)
	hashBytes := md5Hash.Sum(nil)
	return hex.EncodeToString(hashBytes),nil
}

/*
读取io流中的数据进行sha256计算，返回sha256 hash值
 */
func Sha256HashReader(reader io.Reader) (string,error) {
	sha256Hash := sha2562.New()
	readerBytes,err := ioutil.ReadAll(reader)
	if err != nil {
		return "",err
	}
	sha256Hash.Write(readerBytes)
	hashBytes := sha256Hash.Sum(nil)
	return hex.EncodeToString(hashBytes),nil

}