package gen

type Options struct {
	Length       uint
	UseLowercase bool
	UseUppercase bool
	UseNumbers   bool
	UseSpecials  bool
}

func NewOptions(
	length uint,
	useLowercase bool,
	useUppercase bool,
	useNumbers bool,
	useSpecials bool,
) *Options {
	return &Options{
		Length:       length,
		UseLowercase: useLowercase,
		UseUppercase: useUppercase,
		UseNumbers:   useNumbers,
		UseSpecials:  useSpecials,
	}
}
