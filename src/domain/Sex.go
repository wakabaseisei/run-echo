package domain

import "errors"

// 性別
type Sex struct {
	value int
}

func newSex(v int) (*Sex, error) {
	if isValidSex(v) {
		return nil, errors.New("Sex must be set 0, 1, 2 or 9")
	}
	return &Sex{value: v}, nil
}

// 性別はintで扱い、国際規格"ISO 5218"に従う。
// 不明が0・男性が1・女性が2・その他が3
func isValidSex(v int) bool {
	if v == 0 || v == 1 || v == 2 || v == 9 {
		return true
	}
	return false
}
