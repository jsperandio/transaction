package model

type OperationType int

const (
	CompraAVista    OperationType = iota + 1
	CompraParcelada               //  2
	Saque                         //  3
	Pagamento                     //  4
)

func (ot OperationType) String() string {
	return [...]string{"CompraAVista", "CompraParcelada", "Saque", "Pagamento"}[ot-1]
}

func (ot OperationType) EnumIndex() int {
	return int(ot)
}
