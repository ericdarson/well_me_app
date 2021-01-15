package response

type AnnualReportResponse struct {
	ErrorSchema  ErrorSchema                `json:"error_schema"`
	OutputSchema []AnnualReportOutputSchema `json:"output_schema"`
}

type AnnualReportOutputSchema struct {
	BcaId            string             `json:"bcaId"`
	Nama             string             `json:"nama"`
	JoinedTime       string             `json:"joinedTime"`
	InvestmentTimes  string             `json:"investmentTimes"`
	ReportTarget     AnnualReportTarget `json:"annualReportTarget"`
	TopReksadana     []string           `json:"topReksadana"`
	DetailTarget     TargetDetail       `json:"detailTarget"`
	DetailInvestment InvestmentDetail   `json:"detailInvestment"`
	DetailJoin       JoinDetail         `json:"detailJoin"`
}

type AnnualReportTarget struct {
	BestTargetName   string `json:"bestTargetName"`
	BestTargetAmount string `json:"bestTargetAmount"`
	FinishedTarget   string `json:"finishedTargetAmount"`
	OnProgressTarget string `json:"onProgressTarget"`
}

type TargetDetail struct {
	TargetTitle string `json:"targetTitle"`
	TargetImage string `json:"targetImage"`
}
type InvestmentDetail struct {
	InvestmentTitle string `json:"investmentTitle"`
	InvestmentImage string `json:"investmentImage"`
}
type JoinDetail struct {
	JoinTitle string `json:"joinTitle"`
	JoinImage string `json:"joinImage"`
}
