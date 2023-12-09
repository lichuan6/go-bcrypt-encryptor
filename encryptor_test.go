package bcryptencryptor

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestBcryptEncryptor(t *testing.T) {
	encryptor := NewBcryptEncryptor(bcrypt.DefaultCost)

	password := "myPassword"

	// 测试密码加密
	encryptedPassword, err := encryptor.Encrypt(password)
	if err != nil {
		t.Errorf("Failed to encrypt password: %s", err)
	}

	// 测试密码比较（正确的密码）
	err = encryptor.Compare(encryptedPassword, password)
	if err != nil {
		t.Errorf("Password comparison failed for correct password: %s", err)
	}

	// 测试密码比较（错误的密码）
	err = encryptor.Compare(encryptedPassword, "wrongPassword")
	if err == nil {
		t.Errorf("Password comparison succeeded for incorrect password")
	}

	// 测试哈希和密码比较（正确的密码）
	err = encryptor.CompareHashAndPassword(encryptedPassword, password)
	if err != nil {
		t.Errorf("Hash and password comparison failed for correct password: %s", err)
	}

	// 测试哈希和密码比较（错误的密码）
	err = encryptor.CompareHashAndPassword(encryptedPassword, "wrongPassword")
	if err == nil {
		t.Errorf("Hash and password comparison succeeded for incorrect password")
	}
}
