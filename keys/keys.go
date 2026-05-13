package keys

var counter = 0

type KeyVector struct {
	UpperHalf    uint64
	LowerHalf    uint64
}

var (
	CTRL KeyVector
	At   KeyVector
	BackSpace KeyVector
	Tab KeyVector
	Enter KeyVector
	BackSlash KeyVector
	SquareBracketOpen KeyVector
	SquareBracketClose KeyVector
	Caret KeyVector
	Underscore KeyVector
	Shift KeyVector
	Alt KeyVector
	UnicodeCharacter KeyVector
	Space KeyVector
	Minus KeyVector
	EqualTo KeyVector
	Plus KeyVector
	CurlyBracketOpen KeyVector
	CurlyBracketClose KeyVector
	Pipe KeyVector
	SemiColon KeyVector
	Colon KeyVector
	SingleQuote KeyVector
	DoubleQuote KeyVector
	Comma KeyVector
	AngleBracketOpen KeyVector
	AngleBracketClose KeyVector
	Dot KeyVector
	Slash KeyVector
	QuestionMark KeyVector
	BackTick KeyVector
	Tilde KeyVector
	Exclamation KeyVector
	HashTag KeyVector
	Dollar KeyVector
	Percentage KeyVector
	Ampercent KeyVector
	Asterisk KeyVector
	RoundBracketOpen KeyVector
	RoundBracketClose KeyVector
	ArrowUp KeyVector
	ArrowDown KeyVector
	ArrowRight KeyVector
	ArrowLeft KeyVector
	Home KeyVector
	End KeyVector
	F1 KeyVector
	F2 KeyVector
	F3 KeyVector
	F4 KeyVector
	Insert KeyVector
	Delete KeyVector
	PageUp KeyVector
	PageDown KeyVector
	F5 KeyVector
	F6 KeyVector
	F7 KeyVector
	F8 KeyVector
	F9 KeyVector
	F10 KeyVector
	F11 KeyVector
	F12 KeyVector
	ESC KeyVector
	PASTE KeyVector
)

func And(keys ...KeyVector)KeyVector {
	out := KeyVector{}

	for _, key := range keys {
		out.UpperHalf |= key.UpperHalf
		out.LowerHalf |= key.LowerHalf
	}

	return out
}

func (a KeyVector)Equals(b KeyVector) bool {
	return a.LowerHalf == b.LowerHalf && a.UpperHalf == b.UpperHalf
}

func InitializeKeys() {
	if (counter != 0) {
		return
	}

	InitializeAlphabets() // must be initialized first
	InitializeNumbers()

	CTRL = GetNextKeyVector()
	At = GetNextKeyVector()
	BackSpace = GetNextKeyVector()
	Tab = GetNextKeyVector()
	Enter = GetNextKeyVector()
	BackSlash = GetNextKeyVector()
	SquareBracketOpen = GetNextKeyVector()
	SquareBracketClose = GetNextKeyVector()
	Caret = GetNextKeyVector()
	Underscore = GetNextKeyVector()
	Shift = GetNextKeyVector()
	Alt = GetNextKeyVector()
	UnicodeCharacter = GetNextKeyVector()
	Space = GetNextKeyVector()
	Minus = GetNextKeyVector()
	EqualTo = GetNextKeyVector()
	Plus = GetNextKeyVector()
	CurlyBracketOpen = GetNextKeyVector()
	CurlyBracketClose = GetNextKeyVector()
	Pipe = GetNextKeyVector()
	SemiColon = GetNextKeyVector()
	Colon = GetNextKeyVector()
	SingleQuote = GetNextKeyVector()
	DoubleQuote = GetNextKeyVector()
	Comma = GetNextKeyVector()
	AngleBracketOpen = GetNextKeyVector()
	AngleBracketClose = GetNextKeyVector()
	Dot = GetNextKeyVector()
	Slash = GetNextKeyVector()
	QuestionMark = GetNextKeyVector()
	BackTick = GetNextKeyVector()
	Tilde = GetNextKeyVector()
	Exclamation = GetNextKeyVector()
	HashTag = GetNextKeyVector()
	Dollar = GetNextKeyVector()
	Percentage = GetNextKeyVector()
	Ampercent = GetNextKeyVector()
	Asterisk = GetNextKeyVector()
	RoundBracketOpen = GetNextKeyVector()
	RoundBracketClose = GetNextKeyVector()
	ArrowUp = GetNextKeyVector()
	ArrowDown = GetNextKeyVector()
	ArrowRight = GetNextKeyVector()
	ArrowLeft = GetNextKeyVector()
	Home = GetNextKeyVector()
	End = GetNextKeyVector()
	F1 = GetNextKeyVector()
	F2 = GetNextKeyVector()
	F3 = GetNextKeyVector()
	F4 = GetNextKeyVector()
	Delete = GetNextKeyVector()
	Insert = GetNextKeyVector()
	PageUp = GetNextKeyVector()
	PageDown = GetNextKeyVector()
	F5 = GetNextKeyVector()
	F6 = GetNextKeyVector()
	F7 = GetNextKeyVector()
	F8 = GetNextKeyVector()
	F9 = GetNextKeyVector()
	F10 = GetNextKeyVector()
	F11 = GetNextKeyVector()
	F12 = GetNextKeyVector()
	ESC = GetNextKeyVector()
	PASTE = GetNextKeyVector()
}

func GetNextKeyVector() KeyVector {
	count := counter
	counter+=1
	
	if (count < 64) {
		return KeyVector{UpperHalf: 0, LowerHalf: 1 << count}
	} else {
		return KeyVector{UpperHalf: 1 << (count - 64)}
	}
}