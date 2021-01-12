package response

type AnnualReportResponse struct {
	ErrorSchema  ErrorSchema                `json:"error_schema"`
	OutputSchema []AnnualReportOutputSchema `json:"output_schema"`
}

type AnnualReportOutputSchema struct {
	BcaId           string             `json:"bcaId"`
	Nama            string             `json:"nama"`
	JoinedTime      string             `json:"joinedTime"`
	InvestmentTimes string             `json:"investmentTimes"`
	ReportTarget    AnnualReportTarget `json:"annualReportTarget"`
	TopReksadana    []string           `json:"topReksadana"`
	TargetTitle     string             `json:"targetTitle"`
	InvestmentTitle string             `json:"investmentTitle"`
	JoinTitle       string             `json:"joinTitle"`
}

type AnnualReportTarget struct {
	BestTargetName   string `json:"bestTargetName"`
	BestTargetAmount string `json:"bestTargetAmount"`
	FinishedTarget   string `json:"finishedTargetAmount"`
	OnProgressTarget string `json:"onProgressTarget"`
}
