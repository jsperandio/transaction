package model

type OperationType int64

const (
	CompraAVista    OperationType = iota + 1
	CompraParcelada               //  2
	Saque                         //  3
	Pagamento                     //  4
)

func (ot OperationType) String() string {
	return [...]string{"Compra a Vista", "Compra Parcelada", "Saque", "Pagamento"}[ot-1]
}

func (ot OperationType) IsValid() bool {
	return ot > 0 && ot < 5
}

func (ot OperationType) Index() int64 {
	return int64(ot)
}
