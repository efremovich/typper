package main

import (
	"fmt"
	"runtime"
	"time"
	"unicode"

	"github.com/micmonay/keybd_event"
)

func main() {
	symbols := map[rune]int{}
	symbols['1'] = keybd_event.VK_1
	symbols['2'] = keybd_event.VK_2
	symbols['3'] = keybd_event.VK_3
	symbols['4'] = keybd_event.VK_4
	symbols['5'] = keybd_event.VK_5
	symbols['6'] = keybd_event.VK_6
	symbols['7'] = keybd_event.VK_7
	symbols['8'] = keybd_event.VK_8
	symbols['9'] = keybd_event.VK_9
	symbols['0'] = keybd_event.VK_0
	symbols['й'] = keybd_event.VK_Q
	symbols['ц'] = keybd_event.VK_W
	symbols['у'] = keybd_event.VK_E
	symbols['к'] = keybd_event.VK_R
	symbols['е'] = keybd_event.VK_T
	symbols['н'] = keybd_event.VK_Y
	symbols['г'] = keybd_event.VK_U
	symbols['ш'] = keybd_event.VK_I
	symbols['щ'] = keybd_event.VK_O
	symbols['з'] = keybd_event.VK_P
	symbols['х'] = keybd_event.VK_LEFTBRACE
	symbols['ъ'] = keybd_event.VK_RIGHTBRACE
	symbols['ф'] = keybd_event.VK_A
	symbols['ы'] = keybd_event.VK_S
	symbols['в'] = keybd_event.VK_D
	symbols['а'] = keybd_event.VK_F
	symbols['п'] = keybd_event.VK_G
	symbols['р'] = keybd_event.VK_H
	symbols['о'] = keybd_event.VK_J
	symbols['л'] = keybd_event.VK_K
	symbols['д'] = keybd_event.VK_L
	symbols['ж'] = keybd_event.VK_SEMICOLON
	symbols['э'] = keybd_event.VK_APOSTROPHE
	symbols['я'] = keybd_event.VK_Z
	symbols['ч'] = keybd_event.VK_X
	symbols['с'] = keybd_event.VK_C
	symbols['м'] = keybd_event.VK_V
	symbols['и'] = keybd_event.VK_B
	symbols['т'] = keybd_event.VK_N
	symbols['ь'] = keybd_event.VK_M
	symbols[','] = keybd_event.VK_COMMA
	symbols['.'] = keybd_event.VK_DOT
	symbols['m'] = keybd_event.VK_M
	symbols['q'] = keybd_event.VK_Q
	symbols['w'] = keybd_event.VK_W
	symbols['e'] = keybd_event.VK_E
	symbols['r'] = keybd_event.VK_R
	symbols['t'] = keybd_event.VK_T
	symbols['y'] = keybd_event.VK_Y
	symbols['u'] = keybd_event.VK_U
	symbols['i'] = keybd_event.VK_I
	symbols['o'] = keybd_event.VK_O
	symbols['p'] = keybd_event.VK_P
	symbols['a'] = keybd_event.VK_A
	symbols['s'] = keybd_event.VK_S
	symbols['d'] = keybd_event.VK_D
	symbols['f'] = keybd_event.VK_F
	symbols['g'] = keybd_event.VK_G
	symbols['h'] = keybd_event.VK_H
	symbols['j'] = keybd_event.VK_J
	symbols['k'] = keybd_event.VK_K
	symbols['l'] = keybd_event.VK_L
	symbols['z'] = keybd_event.VK_Z
	symbols['x'] = keybd_event.VK_X
	symbols['c'] = keybd_event.VK_C
	symbols['v'] = keybd_event.VK_V
	symbols['b'] = keybd_event.VK_B
	symbols['n'] = keybd_event.VK_N
	symbols['_'] = keybd_event.VK_SPACE

	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// For linux, it is very important wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}
	//strKaka := `Found on the instruction sheet of a Conair Pro Style 1600 hair dryer: "WARNING: Do not use in shower. Never use while sleeping." Found on Bat Man The Animated Series Armor Set Halloween costume box: "PARENT: Please exercise caution, mask and chest plate are not protective; cape does not enable wearer to fly." Found in a television set's owner's manual: "Do not pour liquids into your television set."`
	strKaka := []rune(`Ancient Greeks wove marjoram into funeral wreaths and put them on the graves of loved ones. The wreaths served as prayers for the happiness of the deceased in a future life. Breaking of a glass is traditional in some wedding ceremonies. This custom symbolizes different things. To some it is the destruction of the temple in Jerusalem, and for some it is the represents the fragility of a relationship.`)

	for _, letter := range strKaka {
		if unicode.IsSpace(letter) {

			letter = rune('_')
		}

		kb.HasSHIFT(false)
		if unicode.IsLetter(letter) {
			if unicode.IsUpper(letter) {
				letter = unicode.ToLower(letter)
				kb.HasSHIFT(true)
			}
		}

		kb.SetKeys(symbols[letter])
		fmt.Println(letter)

		err = kb.Launching()
		if err != nil {
			panic(err)
		}

		time.Sleep(200 * time.Millisecond)
	}

}
