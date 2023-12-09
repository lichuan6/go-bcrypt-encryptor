package bcryptencryptor

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordEncryptor interface {
	Encrypt(password string) (string, error)
	Compare(encryptedPassword, password string) error
	CompareHashAndPassword(hashedPassword, password string) error
}

type BcryptEncryptor struct {
	cost int
}

func NewBcryptEncryptor(cost int) *BcryptEncryptor {
	return &BcryptEncryptor{cost: cost}
}

func (b *BcryptEncryptor) Encrypt(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (b *BcryptEncryptor) Compare(encryptedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (b *BcryptEncryptor) CompareHashAndPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
