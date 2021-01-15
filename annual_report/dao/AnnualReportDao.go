package dao

import (
	"annual_report/dbconnection"
	"annual_report/entity/response"
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	_ "github.com/godror/godror"
)

type AnnualReportDao interface {
	GetAnnualReport(string) []response.AnnualReportOutputSchema
}

type annualReportDao struct {
	temp string
}

func NewAnnualReport() AnnualReportDao {
	return &annualReportDao{}
}

func (dao *annualReportDao) GetAnnualReport(bcaId string) []response.AnnualReportOutputSchema {
	var annualReport []response.AnnualReportOutputSchema
	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		single := response.AnnualReportOutputSchema{
			BcaId: "-1",
		}
		annualReport = append(annualReport, single)
		return annualReport
	}
	dat, err := ioutil.ReadFile(dir + "/query/getJoinedTime.query")
	if err != nil {
		fmt.Println(err)
		single := response.AnnualReportOutputSchema{
			BcaId: "-1",
		}
		annualReport = append(annualReport, single)
		return annualReport
	}
	query := string(dat)
	query = fmt.Sprintf(query, bcaId)
	fmt.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		single := response.AnnualReportOutputSchema{
			BcaId: "-1",
		}
		annualReport = append(annualReport, single)
		return annualReport
	}
	defer rows.Close()

	for rows.Next() {
		var BcaId string
		var Nama string
		var JoinedTime string
		var NoRek string
		var InvestmentTimes string
		var ReportTarget response.AnnualReportTarget
		var TopReksadana []string
		var DetailTarget response.TargetDetail
		var DetailInvestment response.InvestmentDetail
		var DetailJoin response.JoinDetail

		rows.Scan(&BcaId, &Nama, &JoinedTime, &NoRek)
		InvestmentTimes = getInvestmentTimes(db, dir, "/query/getInvestmentTimes.query", BcaId)
		ReportTarget = getReportTarget(db, dir, "/query/getPlannerReport.query", BcaId)
		TopReksadana = getTopReksadana(db, dir, "/query/getTopReksadanaReport.query", BcaId)

		var tempMod int
		n, err := strconv.Atoi(NoRek)
		if err == nil {
			tempMod = n % 3
			fmt.Println(n + 1)
		} else {
			fmt.Println(NoRek, "is not an integer.")
		}

		DetailTarget = getTargetDetail(ReportTarget, tempMod)
		DetailInvestment = getInvestDetail(InvestmentTimes, tempMod)
		DetailJoin = getJoinDetail(JoinedTime, tempMod)

		single := response.AnnualReportOutputSchema{
			BcaId:            BcaId,
			Nama:             Nama,
			JoinedTime:       JoinedTime,
			InvestmentTimes:  InvestmentTimes,
			ReportTarget:     ReportTarget,
			TopReksadana:     TopReksadana,
			DetailTarget:     DetailTarget,
			DetailInvestment: DetailInvestment,
			DetailJoin:       DetailJoin,
		}
		annualReport = append(annualReport, single)
	}

	return annualReport
}

func getInvestmentTimes(db *sql.DB, dir string, filename string, bcaId string) string {
	var investmentTimes string
	dat, err := ioutil.ReadFile(dir + filename)
	if err != nil {
		fmt.Println(err)
		return "-1"
	}
	query := string(dat)
	query = fmt.Sprintf(query, bcaId)
	fmt.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return "-1"
	}
	defer rows.Close()
	for rows.Next() {
		var temp1 string

		rows.Scan(&temp1)
		investmentTimes = temp1
	}
	return investmentTimes
}

func getReportTarget(db *sql.DB, dir string, filename string, bcaId string) response.AnnualReportTarget {
	var reportTarget response.AnnualReportTarget
	dat, err := ioutil.ReadFile(dir + filename)
	if err != nil {
		fmt.Println(err)
		return reportTarget
	}
	query := string(dat)
	query = fmt.Sprintf(query, bcaId, bcaId, bcaId)
	fmt.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return reportTarget
	}
	defer rows.Close()
	for rows.Next() {
		var BestTargetName string
		var BestTargetAmount string
		var FinishedTarget string
		var OnProgressTarget string

		rows.Scan(&BestTargetName, &BestTargetAmount, &FinishedTarget, &OnProgressTarget)
		reportTarget = response.AnnualReportTarget{
			BestTargetName:   BestTargetName,
			BestTargetAmount: BestTargetAmount,
			FinishedTarget:   FinishedTarget,
			OnProgressTarget: OnProgressTarget,
		}
	}
	return reportTarget
}

func getTopReksadana(db *sql.DB, dir string, filename string, bcaId string) []string {
	var listReksadana []string
	dat, err := ioutil.ReadFile(dir + filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	query := string(dat)
	query = fmt.Sprintf(query, bcaId)
	fmt.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var NamaReksadana string

		rows.Scan(&NamaReksadana)

		listReksadana = append(listReksadana, NamaReksadana)

	}
	return listReksadana
}

func getTargetDetail(report response.AnnualReportTarget, tempMod int) response.TargetDetail {
	var targetTitle string
	var targetImage string

	var listA = [3]string{"Target Achiever", "On Target", "Completionist!"}
	n, err := strconv.Atoi(report.FinishedTarget)
	if err == nil {
		if n >= 1 {
			targetTitle = listA[tempMod]
			targetImage = fmt.Sprint("Target-", n)
		} else {
			targetTitle = ""
			targetImage = ""
		}
	} else {
		fmt.Println(report.FinishedTarget, "is not an integer.")
		targetTitle = ""
		targetImage = ""
	}

	single := response.TargetDetail{
		TargetTitle: targetTitle,
		TargetImage: targetImage,
	}

	return single
}

func getInvestDetail(investTimes string, tempMod int) response.InvestmentDetail {
	var investTitle string
	var investImage string

	n, err := strconv.Atoi(investTimes)
	if err == nil {
		var listA = [3]string{"Investor Pemula", "Newbie Investor", "Investor Awal"}
		var listB = [3]string{"Investor Menengah", "Intermediate Investor", "Rajin Investasi"}
		var listC = [3]string{"Consistent Invenstor", "Invenstor Disiplin", "Rajin Berinvest"}

		if n < 1 {
			investTitle = ""
			investImage = ""
		} else if n >= 1 && n <= 5 {
			investTitle = listA[tempMod]
			investImage = fmt.Sprint("Invest-", 1)
		} else if n > 5 && n <= 10 {
			investTitle = listB[tempMod]
			investImage = fmt.Sprint("Invest-", 2)
		} else if n > 10 {
			investTitle = listC[tempMod]
			investImage = fmt.Sprint("Invest-", 3)
		}

	} else {
		fmt.Println(investTimes, "is not an integer.")
		investTitle = ""
		investImage = ""
	}

	single := response.InvestmentDetail{
		InvestmentTitle: investTitle,
		InvestmentImage: investImage,
	}
	return single
}

func getJoinDetail(joinedTime string, tempMod int) response.JoinDetail {
	var joinedTitle string
	var joinedImage string

	n, err := strconv.Atoi(joinedTime)
	if err == nil {
		var listA = [3]string{"Newcomer", "Freshman", "Apprentices"}
		var listB = [3]string{" Sudah Betah", "Sudah Nyaman", "Mulai Jago"}
		var listC = [3]string{"Loyalist!", "Expert!", "Veteran!"}

		if n < 90 {
			joinedTitle = listA[tempMod]
			joinedImage = fmt.Sprint("Join-", 1)
		} else if n >= 90 && n <= 180 {
			joinedTitle = listB[tempMod]
			joinedImage = fmt.Sprint("Join-", 2)
		} else if n > 180 {
			joinedTitle = listC[tempMod]
			joinedImage = fmt.Sprint("Join-", 3)
		}

	} else {
		fmt.Println(joinedTime, "is not an integer.")
		joinedTitle = ""
		joinedImage = ""
	}
	single := response.JoinDetail{
		JoinTitle: joinedTitle,
		JoinImage: joinedImage,
	}
	return single
}
