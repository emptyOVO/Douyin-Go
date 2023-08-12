package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"log"
)

type Keys struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func loadKeys() (*Keys, error) {
	//****实际pem文件不能放在服务器中而应该加密保护在本地，这里为了测试图方便直接放在同一包下并且引用了绝对路径****
	privateBytes, err := ioutil.ReadFile("C:\\Users\\高毅飞\\Desktop\\gitsubmit\\Douyin-Go\\Douyin\\utils\\private.pem")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	privateBlock, _ := pem.Decode(privateBytes)
	if privateBlock == nil {
		log.Println(err)
		return nil, errors.New("failed to decode private key")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(privateBlock.Bytes)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		log.Println(err)
		return nil, errors.New("invalid private key type")
	}
	publicBytes, err := ioutil.ReadFile("C:\\Users\\高毅飞\\Desktop\\gitsubmit\\Douyin-Go\\Douyin\\utils\\public.pem")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	publicBlock, _ := pem.Decode(publicBytes)
	if publicBlock == nil {
		log.Println(err)
		return nil, errors.New("failed to decode public key")
	}
	publicKey, err := x509.ParsePKIXPublicKey(publicBlock.Bytes)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		log.Println(err)
		return nil, errors.New("invalid public key type")
	}
	return &Keys{
		PrivateKey: rsaPrivateKey,
		PublicKey:  rsaPublicKey,
	}, nil
}
