package strand
import "strings"

func ToRNA(dna string) string {
	var b strings.Builder
	b.Grow(len(dna))

	for i := 0; i < len(dna); i++ {
		switch dna[i] {
		case 'G':
			b.WriteByte('C')
		case 'C':
			b.WriteByte('G')
		case 'T':
			b.WriteByte('A')
		case 'A':
			b.WriteByte('U')
		default:
			b.WriteByte(dna[i])
		}
	}
	return b.String()
}
