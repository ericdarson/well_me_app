package dao

import (
	"cms_api/dbconnection"
	"cms_api/entity/response"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/godror/godror"
)

type TransactionCMSDao interface {
	InquiryHistoryPembelian(string) []response.TransactionHistoryOutputSchema
	InquiryHistoryPenjualan(string) []response.SellTransactionHistoryOutputSchema
}

type transactionCMSDao struct {
}

func NewTransactionCMS() TransactionCMSDao {
	return &transactionCMSDao{}
}

func returnErrTransactionHistory() []response.TransactionHistoryOutputSchema {
	return []response.TransactionHistoryOutputSchema{response.TransactionHistoryOutputSchema{TanggalTransaksi: "-1"}}
}

func returnErrSellTransactionHistory() []response.SellTransactionHistoryOutputSchema {
	return []response.SellTransactionHistoryOutputSchema{response.SellTransactionHistoryOutputSchema{TanggalTransaksi: "-1"}}
}

func (dao *transactionCMSDao) InquiryHistoryPembelian(bcaid string) []response.TransactionHistoryOutputSchema {
	var outputSchema []response.TransactionHistoryOutputSchema

	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return returnErrTransactionHistory()
	}
	dat, err := ioutil.ReadFile(dir + "/query/InquiryHistoryPembelian.query")
	if err != nil {
		fmt.Println(err)
		return returnErrTransactionHistory()
	}
	query := string(dat)
	query = fmt.Sprintf(query, bcaid)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return returnErrTransactionHistory()
	}
	defer rows.Close()

	for rows.Next() {
		single := response.TransactionHistoryOutputSchema{}
		rows.Scan(&single.TransactionID, &single.BCAID, &single.IDProduk, &single.NamaProduk, &single.IDPlan, &single.NamaPlan, &single.Status, &single.KodePromo, &single.Nab, &single.JumlahUnit, &single.TotalNominal, &single.TanggalTransaksi, &single.TanggalVerifikasiBank, &single.TanggalVerifikasiPembelian)
		outputSchema = append(outputSchema, single)
	}

	return outputSchema
}

func (dao *transactionCMSDao) InquiryHistoryPenjualan(bcaid string) []response.SellTransactionHistoryOutputSchema {
	var outputSchema []response.SellTransactionHistoryOutputSchema

	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return returnErrSellTransactionHistory()
	}
	dat, err := ioutil.ReadFile(dir + "/query/InquiryHistoryPenjualan.query")
	if err != nil {
		fmt.Println(err)
		return returnErrSellTransactionHistory()
	}
	query := string(dat)
	query = fmt.Sprintf(query, bcaid)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return returnErrSellTransactionHistory()
	}
	defer rows.Close()

	for rows.Next() {
		single := response.SellTransactionHistoryOutputSchema{}
		rows.Scan(&single.TransactionID, &single.BCAID, &single.IDProduk, &single.NamaProduk, &single.IDPlan, &single.NamaPlan, &single.Status, &single.Nab, &single.JumlahUnit, &single.TotalNominal, &single.TanggalTransaksi, &single.TanggalVerifikasiBank, &single.TanggalVerifikasiPembelian)
		outputSchema = append(outputSchema, single)
	}

	return outputSchema
}
