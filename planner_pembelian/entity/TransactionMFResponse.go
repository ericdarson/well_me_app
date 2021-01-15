package entity

type TransactionMF struct {
	Reponse TransactionMFParent `json:"BCAAPI06" binding:"required"`
}
type TransactionMFParent struct {
	ErrorResponse TransactionMFErrorResponse `json:"BCAAPI06_O_ERR_RESP" binding:"required"`
	Response      TransactionMFResponse      `json:"BCAAPI06_O_RESP" binding:"required"`
}

type TransactionMFErrorResponse struct {
	ErrorCode    string                    `json:"BCAAPI06_O_ERR_CODE" binding:"required"`
	ErrorMessage TransactionMFErrorMessage `json:"BCAAPI06_O_ERR_MSG" binding:"required"`
}
type TransactionMFErrorMessage struct {
	Indonesian string `json:"BCAAPI06_O_ERR_MSG_IND" binding:"required"`
	English    string `json:"BCAAPI06_O_ERR_MSG_ENG" binding:"required"`
}
type TransactionMFResponse struct {
	User TransactionMFUser `json:"BCAAPI06_I_USER" binding:"required"`
}

type TransactionMFUser struct {
	NoRek        string                 `json:"BCAAPI06_NO_REK" binding:"required"`
	FuncID       string                 `json:"BCAAPI06_FUNC_ID" binding:"required"`
	ListTrn      []TransactionMFListTrn `json:"BCAAPI06_LIST_TRN" binding:"required"`
	TotalNominal float64                `json:"BCAAPI06_TOTAL_NOMINAL" binding:"required"`
}

type TransactionMFListTrn struct {
	Desc    string  `json:"BCAAPI06_DESC" binding:"required"`
	Nominal float64 `json:"BCAAPI06_NOMINAL" binding:"required"`
}
