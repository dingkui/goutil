package encry

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
)

var RsaCoder = NewCoder(func() Coder { return Rsa{} }, nil)

type Rsa struct{}
type rsaUtil struct{}

var RsaUtil = rsaUtil{}

func ReadAll(r io.Reader) ([]byte, error) {
	b := make([]byte, 0, 512)
	for {
		if len(b) == cap(b) {
			// Add more capacity (let append pick how much).
			b = append(b, 0)[:len(b)]
		}
		n, err := r.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return b, err
		}
	}
}

func (Rsa) Encode(reader io.Reader, writer io.Writer, pass []byte) error {
	data, err := ReadAll(reader)
	if err != nil {
		return err
	}
	v15, err := RsaUtil.Encrypt(data, pass)
	if err != nil {
		return err
	}
	writer.Write(v15)
	return nil
}
func (Rsa) Decode(reader io.Reader, writer io.Writer, pass []byte) error {
	cipherInfo, err := ReadAll(reader)
	if err != nil {
		return err
	}
	v15, err := RsaUtil.Decrypt(cipherInfo, pass)
	if err != nil {
		return err
	}
	writer.Write(v15)
	return nil
}

const prKeyType = "RSA PRIVATE KEY"
const puKeyType = "PUBLIC KEY"

func (rsaUtil) GenerateKeys(bits int) (publicKey []byte, privateKey []byte, e error) {
	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}

	// 生成pem公钥
	puKeyBytes, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	puKeyBf := new(bytes.Buffer)
	err = pem.Encode(puKeyBf, &pem.Block{Type: puKeyType, Bytes: puKeyBytes})
	if err != nil {
		return nil, nil, err
	}

	//生成pem私钥
	var prKeyBytes = x509.MarshalPKCS1PrivateKey(key)
	prKeyBf := new(bytes.Buffer)
	err = pem.Encode(prKeyBf, &pem.Block{Type: prKeyType, Bytes: prKeyBytes})
	if err != nil {
		return nil, nil, err
	}

	return puKeyBf.Bytes(), prKeyBf.Bytes(), nil
}
func (rsaUtil) getPublicKey(pubKeyPEM []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pubKeyPEM)
	if block == nil || block.Type != puKeyType {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		return nil, fmt.Errorf("not an RSA public key")
	}
}
func (r rsaUtil) loadPublicKey(fileName string) (*rsa.PublicKey, error) {
	pubKeyPEM, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return r.getPublicKey(pubKeyPEM)
}

func (r rsaUtil) loadPrivateKey(fileName string) (*rsa.PrivateKey, error) {
	privKeyPEM, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return r.getPrivateKey(privKeyPEM)
}

func (r rsaUtil) getPrivateKey(privKeyPEM []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privKeyPEM)
	if block == nil || block.Type != prKeyType {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func (r rsaUtil) Decrypt(cipherInfo []byte, privKeyPEM []byte) ([]byte, error) {
	privKey, err := r.getPrivateKey(privKeyPEM)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, privKey, cipherInfo)
}
func (r rsaUtil) Encrypt(data []byte, pubKeyPEM []byte) ([]byte, error) {
	pubKey, err := RsaUtil.getPublicKey(pubKeyPEM)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptPKCS1v15(rand.Reader, pubKey, data)
}

func test() {
	puKey, prKey, err := RsaUtil.GenerateKeys(2048)
	if err != nil {
		fmt.Println("Error generating RSA key pair:", err)
		return
	}
	ioutil.WriteFile("private_key.pem", prKey, 0644)
	ioutil.WriteFile("public_key.pem", puKey, 0644)

	fmt.Println("RSA Key pair generated and saved.")
}
