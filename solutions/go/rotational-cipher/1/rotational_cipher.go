package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
    	shift := shiftKey % 26
	out := make([]rune, 0, len(plain))

	for _, r := range plain {
		switch {
		case r >= 'A' && r <= 'Z':
			r = 'A' + (r-'A'+rune(shift))%26
		case r >= 'a' && r <= 'z':
			r = 'a' + (r-'a'+rune(shift))%26
		}
		out = append(out, r)
	}

	return string(out)
}
