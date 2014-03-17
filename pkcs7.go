package crypto_padding

import (
    "bytes"
    "errors"
    "fmt"
)

// http://tools.ietf.org/html/rfc5652#section-6.3
type PKCS7 struct {}

func (padding PKCS7) Pad(data []byte, blockSize int) (output []byte, err error) {
    if blockSize < 1 || blockSize >= 256 {
        return output, errors.New(fmt.Sprintf("blocksize is out of bounds: %v", blockSize))
    }
    var paddingBytes = padSize(len(data), blockSize)
    paddingSlice := bytes.Repeat([]byte{byte(paddingBytes)}, paddingBytes)
    output = append(data, paddingSlice...)
    return output, nil
}

func (padding PKCS7) Unpad(data []byte, blockSize int) (output []byte, err error) {
    var dataLen = len(data)
    if dataLen % blockSize != 0 {
        return output, errors.New("data's length isn't a multiple of blockSize")
    }
    var paddingBytes = int(data[dataLen - 1])
    if paddingBytes > blockSize || paddingBytes <= 0 {
        return output, errors.New(fmt.Sprintf("invalid padding found: %v", paddingBytes))
    }
    var pad = data[dataLen - paddingBytes:dataLen - 1]
    for _, v := range pad {
        if int(v) != paddingBytes {
            return output, errors.New("invalid padding found")
        }
    }
    output = data[0:dataLen - paddingBytes]
    return output, nil
}
