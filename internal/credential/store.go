package credential

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Values struct {
	Username    string `json:"username"`
	APIPassword string `json:"api_password"`
	APIKey      string `json:"api_key,omitempty"`
}

func Save(path string, values Values, force bool) error {
	if _, err := os.Stat(path); err == nil && !force {
		return fmt.Errorf("credentials already exist")
	}
	plaintext, err := json.Marshal(values)
	if err != nil {
		return err
	}
	encoded, err := encrypt(plaintext)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return fmt.Errorf("create credentials directory: %w", err)
	}
	return os.WriteFile(path, []byte(encoded), 0o600)
}

func Load(path string) (Values, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Values{}, fmt.Errorf("read credentials: %w", err)
	}
	plaintext, err := decrypt(string(data))
	if err != nil {
		return Values{}, fmt.Errorf("credentials file corrupted or unreadable: %w", err)
	}
	var values Values
	if err := json.Unmarshal(plaintext, &values); err != nil {
		return Values{}, fmt.Errorf("parse credentials: %w", err)
	}
	return values, nil
}

func Delete(path string) error {
	if err := os.Remove(path); errors.Is(err, os.ErrNotExist) {
		return nil
	} else if err != nil {
		return fmt.Errorf("remove credentials: %w", err)
	}
	return nil
}

func encrypt(plaintext []byte) (string, error) {
	block, err := aes.NewCipher(deriveKey())
	if err != nil {
		return "", fmt.Errorf("create cipher: %w", err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("create GCM: %w", err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("generate nonce: %w", err)
	}
	sealed := gcm.Seal(nonce, nonce, plaintext, nil)
	return base64.StdEncoding.EncodeToString(sealed), nil
}

func decrypt(encoded string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, fmt.Errorf("decode credentials: %w", err)
	}
	block, err := aes.NewCipher(deriveKey())
	if err != nil {
		return nil, fmt.Errorf("create cipher: %w", err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("create GCM: %w", err)
	}
	if len(data) < gcm.NonceSize() {
		return nil, errors.New("credentials ciphertext is too short")
	}
	nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errors.New("decryption failed")
	}
	return plaintext, nil
}

func deriveKey() []byte {
	data, err := os.ReadFile("/etc/machine-id")
	if err != nil {
		home, homeErr := os.UserHomeDir()
		if homeErr != nil {
			data = []byte("rss-cli")
		} else {
			data = []byte(home)
		}
	}
	sum := sha256.Sum256(data)
	return sum[:]
}
