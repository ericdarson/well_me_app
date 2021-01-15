package entity

type TransactionMFRequest struct {
	Parent TransactionMFRequestParent `json:"BCAAPI06"`
}

type TransactionMFRequestParent struct {
	Response TransactionMFRequestResp `json:"BCAAPI06_O_RESP"`
}
type TransactionMFRequestResp struct {
	User TransactionMFRequestUser `json:"BCAAPI06_I_USER"`
}

type TransactionMFRequestUser struct {
	FuncID       string                        `json:"BCAAPI06_FUNC_ID"`
	NoRek        string                        `json:"BCAAPI06_NO_REK"`
	ListTrans    []TransactionMFRequestListTrn `json:"BCAAPI06_LIST_TRN"`
	TotalNominal float64                       `json:"BCAAPI06_TOTAL_NOMINAL"`
}

type TransactionMFRequestListTrn struct {
	Desc    string  `json:"BCAAPI06_DESC" binding:"required"`
	Nominal float64 `json:"BCAAPI06_NOMINAL" binding:"required"`
}
