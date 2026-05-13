package keys

// AI assisted

func AsciiKey(b byte) KeyVector {
	switch {
	case b >= 'a' && b <= 'z':
		return GetNthLowerAlphabetKV(int(b - 'a'))

	case b >= 'A' && b <= 'Z':
		return GetNthCapitalAlphabetKV(int(b - 'A'))
	}

	switch b {
		case '0':
			return Number0
		case '1':
			return Number1
		case '2':
			return Number2
		case '3':
			return Number3
		case '4':
			return Number4
		case '5':
			return Number5
		case '6':
			return Number6
		case '7':
			return Number7
		case '8':
			return Number8
		case '9':
			return Number9

		case ' ':
			return Space

		case '-':
			return Minus
		case '_':
			return Underscore

		case '=':
			return EqualTo
		case '+':
			return Plus

		case '[':
			return SquareBracketOpen
		case '{':
			return CurlyBracketOpen

		case ']':
			return SquareBracketClose
		case '}':
			return CurlyBracketClose

		case '\\':
			return BackSlash
		case '|':
			return Pipe

		case ';':
			return SemiColon
		case ':':
			return Colon

		case '\'':
			return SingleQuote
		case '"':
			return DoubleQuote

		case ',':
			return Comma
		case '<':
			return AngleBracketOpen

		case '.':
			return Dot
		case '>':
			return AngleBracketClose

		case '/':
			return Slash
		case '?':
			return QuestionMark

		case '`':
			return BackTick
		case '~':
			return Tilde
		case '!':
			return Exclamation
		case '@':
			return At
		case '#':
			return HashTag
		case '$':
			return Dollar
		case '%':
			return Percentage
		case '^':
			return Caret
		case '&':
			return Ampercent
		case '*':
			return Asterisk
		case '(':
			return RoundBracketOpen
		case ')':
			return RoundBracketClose
	}

	return KeyVector{}
}