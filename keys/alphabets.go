package keys

import "fmt"

var (
	A KeyVector
	B KeyVector
	C KeyVector
	D KeyVector
	E KeyVector
	F KeyVector
	G KeyVector
	H KeyVector
	I KeyVector
	J KeyVector
	K KeyVector
	L KeyVector
	M KeyVector
	N KeyVector
	O KeyVector
	P KeyVector
	Q KeyVector
	R KeyVector
	S KeyVector
	T KeyVector
	U KeyVector
	V KeyVector
	W KeyVector
	X KeyVector
	Y KeyVector
	Z KeyVector
	A_LowerCase KeyVector
	B_LowerCase KeyVector
	C_LowerCase KeyVector
	D_LowerCase KeyVector
	E_LowerCase KeyVector
	F_LowerCase KeyVector
	G_LowerCase KeyVector
	H_LowerCase KeyVector
	I_LowerCase KeyVector
	J_LowerCase KeyVector
	K_LowerCase KeyVector
	L_LowerCase KeyVector
	M_LowerCase KeyVector
	N_LowerCase KeyVector
	O_LowerCase KeyVector
	P_LowerCase KeyVector
	Q_LowerCase KeyVector
	R_LowerCase KeyVector
	S_LowerCase KeyVector
	T_LowerCase KeyVector
	U_LowerCase KeyVector
	V_LowerCase KeyVector
	W_LowerCase KeyVector
	X_LowerCase KeyVector
	Y_LowerCase KeyVector
	Z_LowerCase KeyVector
)

func InitializeAlphabets() {
	if (counter != 0) {
		fmt.Println("Initialize 'InitializeAlphabets()' first. It can only be initialized once.")
		return
	}

	A = GetNextKeyVector()
	B = GetNextKeyVector()
	C = GetNextKeyVector()
	D = GetNextKeyVector()
	E = GetNextKeyVector()
	F = GetNextKeyVector()
	G = GetNextKeyVector()
	H = GetNextKeyVector()
	I = GetNextKeyVector()
	J = GetNextKeyVector()
	K = GetNextKeyVector()
	L = GetNextKeyVector()
	M = GetNextKeyVector()
	N = GetNextKeyVector()
	O = GetNextKeyVector()
	P = GetNextKeyVector()
	Q = GetNextKeyVector()
	R = GetNextKeyVector()
	S = GetNextKeyVector()
	T = GetNextKeyVector()
	U = GetNextKeyVector()
	V = GetNextKeyVector()
	W = GetNextKeyVector()
	X = GetNextKeyVector()
	Y = GetNextKeyVector()
	Z = GetNextKeyVector()
	A_LowerCase = GetNextKeyVector()
	B_LowerCase = GetNextKeyVector()
	C_LowerCase = GetNextKeyVector()
	D_LowerCase = GetNextKeyVector()
	E_LowerCase = GetNextKeyVector()
	F_LowerCase = GetNextKeyVector()
	G_LowerCase = GetNextKeyVector()
	H_LowerCase = GetNextKeyVector()
	I_LowerCase = GetNextKeyVector()
	J_LowerCase = GetNextKeyVector()
	K_LowerCase = GetNextKeyVector()
	L_LowerCase = GetNextKeyVector()
	M_LowerCase = GetNextKeyVector()
	N_LowerCase = GetNextKeyVector()
	O_LowerCase = GetNextKeyVector()
	P_LowerCase = GetNextKeyVector()
	Q_LowerCase = GetNextKeyVector()
	R_LowerCase = GetNextKeyVector()
	S_LowerCase = GetNextKeyVector()
	T_LowerCase = GetNextKeyVector()
	U_LowerCase = GetNextKeyVector()
	V_LowerCase = GetNextKeyVector()
	W_LowerCase = GetNextKeyVector()
	X_LowerCase = GetNextKeyVector()
	Y_LowerCase = GetNextKeyVector()
	Z_LowerCase = GetNextKeyVector()
}

func GetNthCapitalAlphabetKV(n int) KeyVector {
	return KeyVector{UpperHalf: 0, LowerHalf: 1 << n}
}

func GetNthLowerAlphabetKV(n int) KeyVector {
	return KeyVector{UpperHalf: 0, LowerHalf: 1 << (26 + n)}
}