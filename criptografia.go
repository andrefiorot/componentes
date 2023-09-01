package componentes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

const (
	nonceKey = "Vaxdk3l1XXad"
)

// Recebe uma bytea e realiza a criptografia
func Encrypt(plaintext []byte, key string) string {

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err.Error())
	}
	nonce := []byte(nonceKey)
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	//	fmt.Printf("Ciphertext: %x\n", ciphertext)

	return fmt.Sprintf("%x", ciphertext)
}

// Desincripta os dados
func Decrypt(MensagemCriptografada string, key string) ([]byte, error) {
	// key := []byte("blogPostGeekHunterblogPostGeekHu")
	//fmt.Printf("Ciphertext 2: %x\n", MensagemCriptografada)

	ciphertext, _ := hex.DecodeString(string(MensagemCriptografada))

	nonce := []byte(nonceKey)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("Erro ", err)
		return nil, err
		//		panic(err.Error())
	}
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println("Erro p ", err)
		return nil, err
		//		panic(err.Error())
	}
	// fmt.Printf("Plaintext: %s\n", string(plaintext))
	return plaintext, nil
}
