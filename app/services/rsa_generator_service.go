package services

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// RSAGenerator generates private public key pair
type RSAGenerator struct{}

func (yo *RSAGenerator) Generate(size int, privateKeyPath string, publicKeyPath string) {
	privateKey, err := yo.generateKeyPair(size) //2048
	if err != nil {
		fmt.Println("Failed to Generate private key:", err)
		return
	}

	publicKey := &privateKey.PublicKey

	if err := yo.savePrivateKeyToFile(privateKey, privateKeyPath); err != nil {
		fmt.Println("Failed to save private key:", err)
		return
	}
	//todo concurency here

	if err := yo.savePublicKeyToFile(publicKey, publicKeyPath); err != nil {
		fmt.Println("Failed to save public key:", err)
		return
	}
	//todo concurency here

	fmt.Println("Private and public key files have been generated successfully.")
}

func (*RSAGenerator) generateKeyPair(bits int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)

	//todo concurency here
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func (yo *RSAGenerator) savePrivateKeyToFile(privateKey *rsa.PrivateKey, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer yo.closeFile(file)

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	if err := pem.Encode(file, privateKeyPEM); err != nil {
		return err
	}

	return nil
}

func (yo *RSAGenerator) savePublicKeyToFile(publicKey *rsa.PublicKey, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer yo.closeFile(file)

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}

	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	if err := pem.Encode(file, publicKeyPEM); err != nil {
		return err
	}

	return nil
}

func (*RSAGenerator) LoadPrivateKeyFromFile(privateKeyPath string) (*rsa.PrivateKey, error) {
	privateKeyData, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(privateKeyData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing the private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func (*RSAGenerator) LoadPublicKeyFromFile(publicKeyPath string) (*rsa.PublicKey, error) {
	publicKeyData, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(publicKeyData)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing the public key")
	}

	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}

func (*RSAGenerator) closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		panic(err.Error())
	}
}

// IRSAGenerator is interface for RSAGenerator
type IRSAGenerator interface {
	Generate(size int, privateKeyPath string, publicKeyPath string)
	LoadPrivateKeyFromFile(privateKeyPath string) (*rsa.PrivateKey, error)
	LoadPublicKeyFromFile(publicKeyPath string) (*rsa.PublicKey, error)
}

func NewRSAGenerator() IRSAGenerator {
	return &RSAGenerator{}
}
