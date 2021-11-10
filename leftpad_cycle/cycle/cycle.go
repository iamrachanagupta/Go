package cycle

import leftpad "github.com/imrachanagupta/GO/leftpad_cycle"

var DEFAULT_CHAR = ' '

func DoublePad(s string, i int) string {
	return leftpad.Format(s+s, i) // This will create cross reference of the packages that is not allowed in GO.
}
