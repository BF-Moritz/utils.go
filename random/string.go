package random

import (
	"math/rand"

	"github.com/BF-Moritz/utils.go/consts"
	"github.com/BF-Moritz/utils.go/types"
)

const ValidChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func GenerateRandomString(length uint32) (out string) {
	out = ""
	for i := uint32(0); i < length; i++ {
		out += string(ValidChars[int(rand.Uint32())%len(ValidChars)])
	}
	return out
}

func GenerateAdvancedString(config *types.RandomStringConfig) (out string) {
	availableChars := ""
	if config.Override != nil {
		availableChars = *config.Override
	} else {
		if config.Lowercase {
			availableChars += consts.LowerCaseChars
		}
		if config.Uppercase {
			availableChars += consts.UpperCaseChars
		}
		if config.Numbers {
			availableChars += consts.NumberChars
		}
		if config.Symbols {
			availableChars += consts.SymbolChars
		}
	}

	out = ""
	if len(availableChars) <= 0 {
		return
	}

	for i := uint32(0); i < config.Length; i++ {
		out += string(availableChars[int(rand.Uint32())%len(availableChars)])
	}

	return
}
