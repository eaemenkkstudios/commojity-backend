package data

import (
	"errors"
	"strings"
)

type GeneData []byte

// Decode emoji code into gene data
func (g *GeneData) Decode(emojiCode string) error {
	result := []byte{}
	for _, r := range emojiCode {
		if v := strings.Index(DICTIONARY, string(r)); v != -1 {
			result = append(result, byte(v/EMOJI_BYTE))
		} else {
			return errors.New("illegal character on emoji code")
		}
	}
	*g = result
	return nil
}

// Encode gene data into emoji code
func (g *GeneData) Encode() string {
	result := ""
	for _, b := range *g {
		result += DICTIONARY[int(b)*EMOJI_BYTE : (int(b)+1)*EMOJI_BYTE]
	}
	return result
}
