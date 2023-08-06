package services

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

const (
	PRIVATE_KEY_FILE_NAME string = "storage/private.key"
	PUBLIC_KEY_FILE_NAME  string = "storage/public.key"
)

// declare
type IRSAGeneartor interface {
	Generate(size int)
	LoadPrivateKeyFromFile() (*rsa.PrivateKey, error)
	LoadPublicKeyFromFile() (*rsa.PublicKey, error)
}

// define
type RSAGeneartor struct{}

func (rsaGen RSAGeneartor) Generate(size int) {
	privateKey, err := rsaGen.generateKeyPair(size) //2048
	if err != nil {
		fmt.Println("Failed to generate private key:", err)
		return
	}

	publicKey := &privateKey.PublicKey

	privateKeyFile := PRIVATE_KEY_FILE_NAME
	publicKeyFile := PUBLIC_KEY_FILE_NAME

	if err := rsaGen.savePrivateKeyToFile(privateKey, privateKeyFile); err != nil {
		fmt.Println("Failed to save private key:", err)
		return
	}

	if err := rsaGen.savePublicKeyToFile(publicKey, publicKeyFile); err != nil {
		fmt.Println("Failed to save public key:", err)
		return
	}

	fmt.Println("Private and public key files have been generated successfully.")
}

func (RSAGeneartor) generateKeyPair(bits int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
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

func (RSAGeneartor) LoadPrivateKeyFromFile() (*rsa.PrivateKey, error) {
	privateKeyData, err := os.ReadFile(PRIVATE_KEY_FILE_NAME)
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

func (RSAGeneartor) LoadPublicKeyFromFile() (*rsa.PublicKey, error) {
	publicKeyData, err := os.ReadFile(PUBLIC_KEY_FILE_NAME)
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
