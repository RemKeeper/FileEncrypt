package main

import (
	"github.com/wumansgy/goEncrypt/aes"
	"github.com/wumansgy/goEncrypt/des"
	"github.com/wumansgy/goEncrypt/rsa"
)

// DesEncrypt DES加密函数 传入待加密路径和8位密钥 String类型
func DesEncrypt(FileURI, SecretKey string) error {
	file, _, err := ReadFIle(FileURI)
	if err != nil {
		return err
	}
	encryptByte, err := des.DesCbcEncrypt(file, []byte(SecretKey), nil)
	if err != nil {
		return err
	}
	return WriteFile(FileURI+".DesEncrypt", encryptByte)
}

// DesDecrypt DES解密部分 传入加密文件路径和8位密钥
func DesDecrypt(FileURI, SecretKey string) error {
	file, _, err := ReadFIle(FileURI)
	if err != nil {
		return err
	}
	decryptByte, err := des.DesCbcDecrypt(file, []byte(SecretKey), nil)
	if err != nil {
		return err
	}

	return WriteFile(RemoveEncryptionSuffix(FileURI), decryptByte)
}

// TripleDesEncrypt 三重Des加密部分 传入加密文件路径和24位密钥
func TripleDesEncrypt(FileURI, SecretKey string) error {
	file, _, err := ReadFIle(FileURI)
	if err != nil {
		return err
	}
	encrypt, err := des.TripleDesEncrypt(file, []byte(SecretKey), nil)
	if err != nil {
		return err
	}
	return WriteFile(FileURI+".TripleDesEncrypt", encrypt)
}

// TripleDesDecrypt 三重DES解密部分，传入加密文件路径和24位密钥
func TripleDesDecrypt(FileURI, SecretKey string) error {
	file, _, err := ReadFIle(FileURI)
	if err != nil {
		return err
	}
	decrypt, err := des.TripleDesDecrypt(file, []byte(SecretKey), nil)
	if err != nil {
		return err
	}
	return WriteFile(RemoveEncryptionSuffix(FileURI), decrypt)
}

// AesCBCEncrypt AES-CBC模式加密部分 16位密钥
func AesCBCEncrypt(FileURI, SecretKey string) error {
	fIle, _, err := ReadFIle(FileURI)
	if err != nil {
		return err
	}
	encrypt, err := aes.AesCbcEncrypt(fIle, []byte(SecretKey), nil)
	if err != nil {
		return err
	}
	return WriteFile(FileURI+".AesCbcEncrypt", encrypt)
}

// AesCBCDecrypt AES-CBC模式解密部分 16位密钥
func AesCBCDecrypt(FileURI, SecretKey string) error {
	fIle, _, err := ReadFIle(FileURI)
	if err != nil {
		return err
	}
	decrypt, err := aes.AesCbcDecrypt(fIle, []byte(SecretKey), nil)
	if err != nil {
		return err
	}
	return WriteFile(RemoveEncryptionSuffix(FileURI), decrypt)
}

// AesCTREncrypt AES-CTR加密 16位密钥
func AesCTREncrypt(FileURI, SecretKey string) error {
	file, _, err := ReadFIle(FileURI)
	if err != nil {
		return err
	}
	decrypt, err := aes.AesCtrEncrypt(file, []byte(SecretKey), nil)
	if err != nil {
		return err
	}
	return WriteFile(FileURI+".AesCtrEncrypt", decrypt)
}

// AesCTRDecrypt AES-CTR解密部分 16位密钥
func AesCTRDecrypt(FileURI, SecretKey string) error {
	file, _, err := ReadFIle(FileURI)
	if err != nil {
		return err
	}
	decrypt, err := aes.AesCtrDecrypt(file, []byte(SecretKey), nil)
	if err != nil {
		return err
	}
	return WriteFile(RemoveEncryptionSuffix(FileURI), decrypt)
}

// RsaEncrypt RSA加密函数 自动生成密钥
func RsaEncrypt(FileURI string) error {
	file, _, err := ReadFIle(FileURI)
	if err != nil {
		return err
	}
	rsaBase64Key, err := rsa.GenerateRsaKeyBase64(1024)
	if err != nil {
		return err
	}
	EncryptHex, err := rsa.RsaEncryptToBase64(file, rsaBase64Key.PublicKey)
	if err != nil {
		return err
	}
	err = WriteFile(RemoveFileName(FileURI)+"PrivateKey", []byte(rsaBase64Key.PrivateKey))
	if err != nil {
		return err
	}

	return WriteFile(FileURI+".RsaEncrypt", []byte(EncryptHex))
}

// RsaDecrypt RSA解密函数 传入私钥
func RsaDecrypt(FileURI string, PrivateKey []byte) error {
	file, _, err := ReadFIle(FileURI)
	if err != nil {
		return err
	}
	Decrypt, err := rsa.RsaDecryptByBase64(string(file), string(PrivateKey))
	if err != nil {
		return err
	}
	return WriteFile(RemoveEncryptionSuffix(FileURI), Decrypt)
}
