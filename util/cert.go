package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path"
)

var pwd, _ = os.Getwd()

func GenRsaCertPair() {
	// Generate RSA private key with size 2048 bits
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Failed to generate RSA private key:", err)
		return
	}

	// Encode private key to PEM format
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// Write private key to file
	privateKeyFile, err := os.Create("private_key.pem")
	if err != nil {
		fmt.Println("Failed to create private key file:", err)
		return
	}
	defer privateKeyFile.Close()
	if err := pem.Encode(privateKeyFile, privateKeyPEM); err != nil {
		fmt.Println("Failed to write private key to file:", err)
		return
	}
	fmt.Println("Private key generated and saved to private_key.pem")

	// Generate public key from private key
	publicKey := privateKey.PublicKey

	// Encode public key to PEM format
	publicKeyDER, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		fmt.Println("Failed to encode public key:", err)
		return
	}
	publicKeyPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyDER,
	}

	// Write public key to file
	publicKeyFile, err := os.Create("public_key.pem")
	if err != nil {
		fmt.Println("Failed to create public key file:", err)
		return
	}
	defer publicKeyFile.Close()
	if err := pem.Encode(publicKeyFile, publicKeyPEM); err != nil {
		fmt.Println("Failed to write public key to file:", err)
		return
	}
	fmt.Println("Public key generated and saved to public_key.pem")
}

func ReadPriKey(filePath string) (*rsa.PrivateKey, error) {
	// Read the key file
	keyBytes, err := os.ReadFile(path.Join(pwd, filePath))
	if err != nil {
		return nil, err
	}

	// Decode PEM block
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	// Parse RSA private key
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// ReadPubKey 	读取公钥
func ReadPubKey(filePath string) (*rsa.PublicKey, error) {
	// Read the key file
	keyBytes, err := os.ReadFile(path.Join(pwd, filePath))
	if err != nil {
		return nil, err
	}

	return BuildPubKey(string(keyBytes))
}

// BuildPubKey 	构建公钥
func BuildPubKey(pubKey string) (*rsa.PublicKey, error) {

	// Decode PEM block
	block, _ := pem.Decode([]byte(pubKey))
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	// Parse RSA public key
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// Type assertion to RSA public key
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}

	return rsaPublicKey, nil
}
