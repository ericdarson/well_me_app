package entity

type GetSaldo struct {
	Reponse GetSaldoModel `json:"BCAAPI02" binding:"required"`
}
type GetSaldoModel struct {
	Response      GetSaldoResponse      `json:"BCAAPI02_O_RESP" binding:"required"`
	ErrorResponse GetSaldoErrorResponse `json:"BCAAPI02_O_ERR_RESP" binding:"required"`
}

type GetSaldoResponse struct {
	NoRek string  `json:"BCAAPI02" binding:"required"`
	Saldo float64 `json:"BCAAPI02_SALDO" binding:"required"`
}
type GetSaldoErrorResponse struct {
	ErrorCode    string               `json:"BCAAPI02_O_ERR_CODE" binding:"required"`
	ErrorMessage GetSaldoErrorMessage `json:"BCAAPI02_O_ERR_MSG" binding:"required"`
}

type GetSaldoErrorMessage struct {
	Indonesian string `json:"BCAAPI02_O_ERR_MSG_IND" binding:"required"`
	English    string `json:"BCAAPI02_O_ERR_MSG_ENG" binding:"required"`
}
