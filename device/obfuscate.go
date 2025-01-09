package device

import (
	"os"
)

var obfuscateEnabled bool
var obfuscateKey []byte

func InitObfuscation(logger *Logger) {
	ENV_WG_OBFUSCATE_KEY := os.Getenv("WG_OBFUSCATE_KEY")
	if ENV_WG_OBFUSCATE_KEY != "" {
		obfuscateEnabled = true
		obfuscateKey = []byte(ENV_WG_OBFUSCATE_KEY)
		logger.Verbosef("XOR Obfuscation enabled")
	}
}

// XORs bytes in place with the obfuscation key if obfuscation is enabled
func ObfuscateBytes(inBytes []byte) {
	if !obfuscateEnabled {
		return
	}

	keyLen := len(obfuscateKey)
	for i := 0; i < len(inBytes); i++ {
		inBytes[i] ^= obfuscateKey[i%keyLen]
	}
}

// Calls ObfuscateBytes
func DeobfuscateBytes(inBytes []byte) {
	ObfuscateBytes(inBytes)
}
