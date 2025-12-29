package sorting
import "fmt"
import "strconv"
// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
    return fmt.Sprintf("This is the number %.1f", f)
	panic("Please implement DescribeNumber")
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
	panic("Please implement DescribeNumberBox")
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	// 使用类型断言判断是否为 FancyNumber 类型
	if fancy, ok := fnb.(FancyNumber); ok {
		// 如果是 FancyNumber 类型，解析字符串中的数字并返回
		// 这里假设 FancyNumber 的 Value() 是一个合法的数字字符串
		val, err := strconv.Atoi(fancy.Value())
		if err == nil {
			return val
		}
	}
	// 如果不是 FancyNumber，返回 0
	return 0
    }

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
    val := ExtractFancyNumber(fnb)
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(val))
	panic("Please implement DescribeFancyNumberBox")
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
    switch v := i.(type) {
	case int:
    	return DescribeNumber(float64(v))
    case float64:
        return DescribeNumber(v)
	case NumberBox:
    	return DescribeNumberBox(v)
	case FancyNumberBox:
    	return DescribeFancyNumberBox(v)
	default:
    	return "Return to sender"
	}
}
