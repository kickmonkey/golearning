package clock
import "fmt"
type Clock struct {
    minutes int 
}

const dayMinutes = 24 * 60
func normalize(m int) int {
	m %= dayMinutes
	if m < 0 {
		m += dayMinutes
	}
	return m
}
func New(h, m int) Clock {
	return Clock{minutes: normalize(h*60 + m)}
}

func (c Clock) Add(m int) Clock {
	return Clock{minutes: normalize(c.minutes + m)}
}

func (c Clock) Subtract(m int) Clock {
	return Clock{minutes: normalize(c.minutes - m)}
}

func (c Clock) String() string {
	h := c.minutes / 60
	min := c.minutes % 60
	return fmt.Sprintf("%02d:%02d", h, min)
}
