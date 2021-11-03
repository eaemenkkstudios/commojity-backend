package main

import (
	"errors"
	"math"
	"math/big"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/x448/float16"
)

const dictionary = "🍕🍔🍟🌭🍿🧂🥓🥚🍳🧇🥞🧈🍞🥐🥨🥯🥖🫓🧀🥗🥙🥪🌮🌯🫔🥫🍖🍗🥩🍠🥟🥠🥡🍱🍘🍙🍚🍛🍜🦪🍣🍤🍥🥮🍢🧆🥘🍲🫕🍝🥣🥧🍦🍧🍨🍩🍪🎂🍰🧁🍫🍬🍭🍡"

func emojiCodeToInt(emojiCode string) (*big.Int, error) {
	const EMOJI_BYTE = 4
	result := big.NewInt(0)
	for i, c := range emojiCode {
		if v := strings.Index(dictionary, string(c)); v != -1 {
			base := big.NewInt(int64(len(dictionary) / EMOJI_BYTE))
			base.Mul(
				base.Exp(base,
					big.NewInt(int64(len(emojiCode)/EMOJI_BYTE-1-i/EMOJI_BYTE)),
					nil,
				),
				big.NewInt(int64(v/EMOJI_BYTE)),
			)
			result.Add(result, base)
		} else {
			return nil, errors.New("illegal character on emoji code")
		}
	}
	return result, nil
}

type Distro struct {
	Inputs         float16.Float16
	Maintenance    float16.Float16
	Transportation float16.Float16
}

func (d *Distro) Fill(number *big.Int) {
	if len(number.Bytes()) >= 2 {
		d.Inputs = float16.Fromfloat32(math.Float32frombits(uint32(number.Bytes()[0])<<8 | uint32(number.Bytes()[1])))
		if len(number.Bytes()) >= 4 {
			d.Maintenance = float16.Fromfloat32(math.Float32frombits(uint32(number.Bytes()[1])<<8 | uint32(number.Bytes()[2])))
			if len(number.Bytes()) >= 6 {
				d.Transportation = float16.Fromfloat32(math.Float32frombits(uint32(number.Bytes()[3])<<8 | uint32(number.Bytes()[4])))
			}
		}
	}
}

func main() {
	// New Gin instance
	r := gin.Default()

	r.GET("/gene", func(c *gin.Context) {
		geneParam := c.Request.URL.Query()["g"]
		if len(geneParam) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "gene not present"})
		}
		if number, err := emojiCodeToInt(geneParam[0]); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			d := Distro{}
			d.Fill(number)
			c.JSON(http.StatusOK, gin.H{"number": number.String(), "distro": d})
		}
	})

	r.Run()
}
