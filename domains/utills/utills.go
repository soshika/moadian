package utills

var (
	UtillsService utillsServiceInterface = &utillsService{}
)

type utillsService struct{}

type utillsServiceInterface interface {
	XorArrays([]byte, []byte) ([]byte, error)
}

func (s *utillsService) XorArrays(a, b []byte) ([]byte, error) {
	// سرویس XOR کردن دو آرایه بایت
	var c []byte

	aLen := len(a)
	bLen := len(b)
	min := 0
	size := 0
	if aLen > bLen {
		size = aLen
	} else {
		size = bLen
	}

	if size == aLen {
		min = bLen
		copy(c[min:], a[min:size])
	} else {
		min = aLen
		copy(c[min:], b[min:size])
	}

	for i := 0; i < min; i++ {
		c[i] = a[i] ^ b[i]
	}

	return c, nil
}
