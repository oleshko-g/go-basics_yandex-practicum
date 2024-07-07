package mathSlice

type Slice []Element
type Element int

func SumSlice(s Slice) (sum Element) {
	for _, v := range s {
		sum += v
	}
	return
}

func MapSlice(s Slice, op func(Element) Element) {
	for i, v := range s {
		s[i] = op(v)
	}
}

func FoldSlice(s Slice, op func(Element, Element) Element, init Element) (res Element) {
	res = op(init, s[0])
	for i := 1; i < len(s); i++ {
		res = op(res, s[i])
	}
	return
}
