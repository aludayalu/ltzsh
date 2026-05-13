package keys

var (
	Number0 KeyVector
	Number1 KeyVector
	Number2 KeyVector
	Number3 KeyVector
	Number4 KeyVector
	Number5 KeyVector
	Number6 KeyVector
	Number7 KeyVector
	Number8 KeyVector
	Number9 KeyVector
)

func InitializeNumbers() {
	Number0 = GetNextKeyVector()
	Number1 = GetNextKeyVector()
	Number2 = GetNextKeyVector()
	Number3 = GetNextKeyVector()
	Number4 = GetNextKeyVector()
	Number5 = GetNextKeyVector()
	Number6 = GetNextKeyVector()
	Number7 = GetNextKeyVector()
	Number8 = GetNextKeyVector()
	Number9 = GetNextKeyVector()
}