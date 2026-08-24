package main

// Secret-Vault-Manager: encrypted key-value store
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Vault struct {
	Path string
	Key  []byte
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: secret-vault-manager <vault-file> <key-hex> [get|set|list] [name] [value]")
		os.Exit(1)
	}
	vaultPath := os.Args[1]
	keyHex := os.Args[2]

	key, err := hex.DecodeString(keyHex)
	if err != nil || len(key) != 32 {
		fmt.Fprintln(os.Stderr, "key must be 64 hex chars (32 bytes)")
		os.Exit(1)
	}

	v := Vault{Path: vaultPath, Key: key}
	data := v.load()

	cmd := "list"
	if len(os.Args) > 3 {
		cmd = os.Args[3]
	}

	switch cmd {
	case "list":
		names := make([]string, 0, len(data))
		for n := range data {
			names = append(names, n)
		}
		fmt.Printf("%d secrets:\n", len(names))
		for _, n := range names {
			val := data[n]
			masked := val
			if len(masked) > 4 {
				masked = masked[:2] + "****" + masked[len(masked)-2:]
			}
			fmt.Printf("  %s = %s\n", n, masked)
		}
	case "get":
		if len(os.Args) < 5 {
			fmt.Fprintln(os.Stderr, "need name")
			os.Exit(1)
		}
		name := os.Args[4]
		val, ok := data[name]
		if !ok {
			fmt.Fprintln(os.Stderr, "not found")
			os.Exit(1)
		}
		fmt.Println(val)
	case "set":
		if len(os.Args) < 6 {
			fmt.Fprintln(os.Stderr, "need name and value")
			os.Exit(1)
		}
		name := os.Args[4]
		value := os.Args[5]
		data[name] = value
		v.save(data)
		fmt.Printf("set %s (%d chars)\n", name, len(value))
	default:
		fmt.Println("unknown command:", cmd)
		os.Exit(1)
	}
}

func (v Vault) load() map[string]string {
	data := map[string]string{}
	raw, err := os.ReadFile(v.Path)
	if err != nil {
		return data
	}
	if len(raw) < 12 {
		return data
	}
	block, _ := aes.NewCipher(v.Key)
	gcm, _ := cipher.NewGCM(block)
	nonce := raw[:12]
	ct := raw[12:]
	pt, err := gcm.Open(nil, nonce, ct, nil)
	if err != nil {
		return data
	}
	json.Unmarshal(pt, &data)
	return data
}

func (v Vault) save(data map[string]string) {
	pt, _ := json.Marshal(data)
	block, _ := aes.NewCipher(v.Key)
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, 12)
	io.ReadFull(rand.Reader, nonce)
	ct := gcm.Seal(nil, nonce, pt, nil)
	out := append(nonce, ct...)
	os.WriteFile(v.Path, out, 0600)
}
