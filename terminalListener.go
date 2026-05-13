package main

import (
	"bytes"
	"ltz/keys"
	"ltz/shared"
	"os"
	"golang.org/x/term"
	"unicode/utf8"
)

// file written originally by ChatGPT

func terminalListener(events chan<- shared.Event, listener_cleanup *func()) {
	fd := int(os.Stdin.Fd())

	oldState, _ := term.MakeRaw(fd)

	os.Stdout.Write([]byte("\x1b[?1049h")) // alternate screen buffer

	os.Stdout.Write([]byte("\x1b[?25l")) // hides cursor

	os.Stdout.Write([]byte("\x1b[?2004h")) // enable bracketed paste

	os.Stdout.Write([]byte("\x1b[?1003h\x1b[?1006h"))

	cleanup := func() {
		os.Stdout.Write([]byte("\x1b[?1003l\x1b[?1006l"))
		os.Stdout.Write([]byte("\x1b[0m")) // resets terminal styles
		os.Stdout.Write([]byte("\x1b[?2004l")) // disables bracketed paste
		os.Stdout.Write([]byte("\x1b[?25h")) // unhides cursor
		os.Stdout.Write([]byte("\x1b[?1049l")) // restores screen buffer
		os.Stdout.Sync()
		term.Restore(fd, oldState)
	}

	*listener_cleanup = cleanup

	defer cleanup()

	buf := make([]byte, 128)
	pending := make([]byte, 0, 256)
	inPaste := false
	pasteBuffer := make([]byte, 0, 256)
	pasteEnd := []byte("\x1b[201~")

	emitKey := func(key keys.KeyVector) {
		events <- shared.Event{Type: shared.ENUM_EVENT_KEY, KeyData: &shared.KeyEventData{Key: key}}
	}

	emitKeyWithData := func(key keys.KeyVector, data *string) {
		events <- shared.Event{Type: shared.ENUM_EVENT_KEY, KeyData: &shared.KeyEventData{Key: key, Data: data}}
	}

	controlKey := func(b byte) (keys.KeyVector, bool) {
		switch b {
			case 0:
				return keys.And(keys.CTRL, keys.At), true
			case 8:
				return keys.BackSpace, true
			case 9:
				return keys.Tab, true
			case 10, 13:
				return keys.Enter, true
			case 28:
				return keys.And(keys.CTRL, keys.BackSlash), true
			case 29:
				return keys.And(keys.CTRL, keys.SquareBracketClose), true
			case 30:
				return keys.And(keys.CTRL, keys.Caret), true
			case 31:
				return keys.And(keys.CTRL, keys.Underscore), true
			case 127:
				return keys.BackSpace, true
		}

		if b >= 1 && b <= 26 {
			return keys.And(keys.CTRL, keys.GetNthCapitalAlphabetKV(int(b - 1))), true
		}

		return keys.KeyVector{}, false
	}

	modifierPrefix := func(mod int) keys.KeyVector {
		switch mod {
			case 2:
				return keys.Shift
			case 3:
				return keys.Alt
			case 4:
				return keys.And(keys.Alt, keys.Shift)
			case 5:
				return keys.CTRL
			case 6:
				return keys.And(keys.CTRL, keys.Shift)
			case 7:
				return keys.And(keys.CTRL, keys.Alt)
			case 8:
				return keys.And(keys.CTRL, keys.Alt, keys.Shift)
		}

		return keys.KeyVector{}
	}

	parseSingleKey := func(data []byte) (keys.KeyVector, int, bool) {
		if len(data) == 0 {
			return keys.KeyVector{}, 0, true
		}

		b := data[0]

		if key, ok := controlKey(b); ok {
			return key, 1, false
		}

		if b < 0x80 {
			return keys.AsciiKey(b), 1, false
		}

		return keys.KeyVector{}, 1, false
	}

	parseSS3 := func(data []byte) (int, bool) {
		if len(data) < 3 {
			return 0, true
		}

		if data[0] != 0x1b || data[1] != 'O' {
			return 0, false
		}

		var key keys.KeyVector

		switch data[2] {
			case 'A':
				key = keys.ArrowUp
			case 'B':
				key = keys.ArrowDown
			case 'C':
				key = keys.ArrowRight
			case 'D':
				key = keys.ArrowLeft
			case 'H':
				key = keys.Home
			case 'F':
				key = keys.End
			case 'P':
				key = keys.F1
			case 'Q':
				key = keys.F2
			case 'R':
				key = keys.F3
			case 'S':
				key = keys.F4
			default:
				return 3, false
		}

		emitKey(key)
		return 3, false
	}

	parseCSI := func(data []byte) (int, bool) {
		if len(data) < 3 {
			return 0, true
		}

		if data[0] != 0x1b || data[1] != '[' {
			return 0, false
		}

		if data[2] == '<' {
			j := 3
			nums := [3]int{}
			k := 0
			val := 0
			haveVal := false

			for j < len(data) {
				c := data[j]

				if c >= '0' && c <= '9' {
					val = val*10 + int(c-'0')
					haveVal = true
				} else if c == ';' {
					if k < 3 {
						if haveVal {
							nums[k] = val
						} else {
							nums[k] = 0
						}
						k++
					}
					val = 0
					haveVal = false
				} else if c == 'M' || c == 'm' {
					if k < 3 {
						if haveVal {
							nums[k] = val
						} else {
							nums[k] = 0
						}
					}

					events <- shared.Event{
						Type: shared.ENUM_EVENT_MOUSE,
						MouseData: &shared.MouseEventData{
							Button:  nums[0],
							X:       nums[1],
							Y:       nums[2],
							Pressed: map[bool]int{true: 1, false: 0}[c == 'M'],
						},
					}

					return j + 1, false
				} else {
					return j + 1, false
				}

				j++
			}

			return 0, true
		}

		params := make([]int, 0, 4)
		val := 0
		haveVal := false
		j := 2

		appendParam := func() {
			if haveVal {
				params = append(params, val)
			} else {
				params = append(params, 0)
			}
			val = 0
			haveVal = false
		}

		for j < len(data) {
			c := data[j]

			if c >= '0' && c <= '9' {
				val = val*10 + int(c-'0')
				haveVal = true
			} else if c == ';' {
				appendParam()
			} else if c >= '@' && c <= '~' {
				if haveVal || len(params) > 0 {
					appendParam()
				}

				final := c
				mod := 1

				if len(params) >= 2 && params[1] != 0 {
					mod = params[1]
				}

				prefix := modifierPrefix(mod)
				var key keys.KeyVector

				switch final {
					case 'A':
						key = keys.And(prefix, keys.ArrowUp)
					case 'B':
						key = keys.And(prefix, keys.ArrowDown)
					case 'C':
						key = keys.And(prefix, keys.ArrowRight)
					case 'D':
						key = keys.And(prefix, keys.ArrowLeft)
					case 'H':
						key = keys.And(prefix, keys.Home)
					case 'F':
						key = keys.And(prefix, keys.End)
					case 'P':
						key = keys.And(prefix, keys.F1)
					case 'Q':
						key = keys.And(prefix, keys.F2)
					case 'R':
						key = keys.And(prefix, keys.F3)
					case 'S':
						key = keys.And(prefix, keys.F4)
					case 'Z':
						if len(params) == 0 {
							key = keys.And(keys.Shift, keys.Tab)
						} else {
							key = keys.And(prefix, keys.Tab)
						}
					case '~':
						if len(params) == 0 {
							return j + 1, false
						}

						switch params[0] {
							case 1:
								key = keys.And(prefix, keys.Home)
							case 2:
								key = keys.And(prefix, keys.Insert)
							case 3:
								key = keys.And(prefix, keys.Delete)
							case 4:
								key = keys.And(prefix, keys.End)
							case 5:
								key = keys.And(prefix, keys.PageUp)
							case 6:
								key = keys.And(prefix, keys.PageDown)
							case 11:
								key = keys.And(prefix, keys.F1)
							case 12:
								key = keys.And(prefix, keys.F2)
							case 13:
								key = keys.And(prefix, keys.F3)
							case 14:
								key = keys.And(prefix, keys.F4)
							case 15:
								key = keys.And(prefix, keys.F5)
							case 17:
								key = keys.And(prefix, keys.F6)
							case 18:
								key = keys.And(prefix, keys.F7)
							case 19:
								key = keys.And(prefix, keys.F8)
							case 20:
								key = keys.And(prefix, keys.F9)
							case 21:
								key = keys.And(prefix, keys.F10)
							case 23:
								key = keys.And(prefix, keys.F11)
							case 24:
								key = keys.And(prefix, keys.F12)
							case 200:
								inPaste = true
								pasteBuffer = pasteBuffer[:0]
								return j + 1, false
							default:
								return j + 1, false
						}
					default:
						return j + 1, false
				}

				emitKey(key)
				return j + 1, false
			} else {
				return j + 1, false
			}

			j++
		}

		return 0, true
	}

	parseEscape := func(data []byte) (int, bool) {
		if len(data) == 0 {
			return 0, true
		}

		if data[0] != 0x1b {
			return 0, false
		}

		if len(data) == 1 {
			return 0, true
		}

		if data[1] == '[' {
			if consumed, needMore := parseCSI(data); needMore {
				return 0, true
			} else if consumed > 0 {
				return consumed, false
			}
		}

		if data[1] == 'O' {
			if consumed, needMore := parseSS3(data); needMore {
				return 0, true
			} else if consumed > 0 {
				return consumed, false
			}
		}

		key, used, needMore := parseSingleKey(data[1:])
		if needMore {
			return 0, true
		}
		if key.Equals(keys.KeyVector{}) {
			return 1 + used, false
		}

		emitKey(keys.And(keys.Alt, key))
		return 1 + used, false
	}

	for {
		n, err := os.Stdin.Read(buf)

		if err != nil {
			close(events)
			return
		}

		pending = append(pending, buf[:n]...)

		i := 0
		for i < len(pending) {
			b := pending[i]

			if inPaste {
				if len(pending[i:]) < len(pasteEnd) && bytes.HasPrefix(pasteEnd, pending[i:]) {
					break
				}
				if bytes.HasPrefix(pending[i:], pasteEnd) {
					inPaste = false
					pasted := string(pasteBuffer)
					events <- shared.Event{Type: shared.ENUM_EVENT_KEY, KeyData: &shared.KeyEventData{Key: keys.PASTE, Data: &pasted}}
					i += len(pasteEnd)
					continue
				}
				pasteBuffer = append(pasteBuffer, b)
				i++
				continue
			}

			if b == 0x1b {
				consumed, needMore := parseEscape(pending[i:])
				if needMore {
					break
				}

				if consumed > 0 {
					i += consumed
					continue
				}

				emitKey(keys.ESC)
				i++
				continue
			}

			if key, ok := controlKey(b); ok {
				emitKey(key)

				i++
				continue
			}

			if b < 0x80 {
				emitKey(keys.AsciiKey(b))
				i++
				continue
			}

			if !utf8.FullRune(pending[i:]) {
				break
			}

			r, size := utf8.DecodeRune(pending[i:])
			unicode_char := string(r)
			emitKeyWithData(keys.UnicodeCharacter, &unicode_char)
			i += size
		}

		if i > 0 {
			pending = append(pending[:0], pending[i:]...)
		}

		if len(pending) == 1 && pending[0] == 0x1b {
			emitKey(keys.ESC)
			pending = pending[:0]
		}
	}
}