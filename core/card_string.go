// Code generated by "stringer -type Card"; DO NOT EDIT.

package core

import "strconv"

const _Card_name = "ChopsticksTempuraSashimiDumplingsSquidNigiriSalmonNigiriEggNigiriWasabiMakiRoll1MakiRoll2MakiRoll3Pudding"

var _Card_index = [...]uint8{0, 10, 17, 24, 33, 44, 56, 65, 71, 80, 89, 98, 105}

func (i Card) String() string {
	if i < 0 || i >= Card(len(_Card_index)-1) {
		return "Card(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Card_name[_Card_index[i]:_Card_index[i+1]]
}
