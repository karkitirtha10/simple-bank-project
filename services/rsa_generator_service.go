package services

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// define
type RSAGeneartor struct{}

func (rsaGen RSAGeneartor) Generate(size int, privateKeyPath string, publicKeyPath string) {
	privateKey, err := rsaGen.generateKeyPair(size) //2048
	if err != nil {
		fmt.Println("Failed to generate private key:", err)
		return
	}

	publicKey := &privateKey.PublicKey

	if err := rsaGen.savePrivateKeyToFile(privateKey, privateKeyPath); err != nil {
		fmt.Println("Failed to save private key:", err)
		return
	}
	//todo concurency here

	if err := rsaGen.savePublicKeyToFile(publicKey, publicKeyPath); err != nil {
		fmt.Println("Failed to save public key:", err)
		return
	}
	//todo concurency here

	fmt.Println("Private and public key files have been generated successfully.")
}

func (RSAGeneartor) generateKeyPair(bits int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)

	//todo concurency here
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func (RSAGeneartor) savePrivateKeyToFile(privateKey *rsa.PrivateKey, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

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

func (RSAGeneartor) savePublicKeyToFile(publicKey *rsa.PublicKey, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

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

func (RSAGeneartor) LoadPrivateKeyFromFile(privateKeyPath string) (*rsa.PrivateKey, error) {
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

func (RSAGeneartor) LoadPublicKeyFromFile(publicKeyPath string) (*rsa.PublicKey, error) {
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

// declare
type IRSAGeneartor interface {
	Generate(size int, privateKeyPath string, publicKeyPath string)
	LoadPrivateKeyFromFile(privateKeyPath string) (*rsa.PrivateKey, error)
	LoadPublicKeyFromFile(publicKeyPath string) (*rsa.PublicKey, error)
}
