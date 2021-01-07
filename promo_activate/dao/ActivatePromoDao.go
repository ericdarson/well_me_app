package dao

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"promo_activate/dbconnection"
	"promo_activate/entity/request"
	"promo_activate/entity/response"
	"time"

	_ "github.com/godror/godror"
)

type PromoActivateDao interface {
	ActivatePromo(request.ActivatePromoRequest) string
	InquiryPromo(string) response.ListPromoOutputSchema
	ClaimPromo(request.ActivatePromoRequest) string
}

type promoActivateDao struct {
}

func NewPromoActivate() PromoActivateDao {
	return &promoActivateDao{}
}

func (dao *promoActivateDao) ActivatePromo(req request.ActivatePromoRequest) string {

	conn := dbconnection.New()
	db := conn.GetConnection()
	var result string

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `BEGIN SP_ACTIVATE_PROMO_AKUMULASI('%s','%s','SYSTEM API',:1 ); END;`
	query = fmt.Sprintf(query, req.BcaID, req.KodePromo)
	if _, err := db.ExecContext(ctx, query, sql.Out{Dest: &result}); err != nil {
		log.Printf("Error running %q: %+v", query, err)
		return "TIME OUT"
	}

	return result
}

func returnErrListPromo() response.ListPromoOutputSchema {
	return response.ListPromoOutputSchema{
		Objectives: []response.Objectives{response.Objectives{
			KodePromo: "-1",
		}},
		Promotion: nil,
	}
}

func (dao *promoActivateDao) InquiryPromo(BcaID string) response.ListPromoOutputSchema {
	var outputSchema response.ListPromoOutputSchema
	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return returnErrListPromo()
	}

	var listPromotions []response.Promotion
	dat, err := ioutil.ReadFile(dir + "/query/InquiryPromo.query")
	if err != nil {
		fmt.Println(err)
		return returnErrListPromo()
	}
	query := string(dat)
	query = fmt.Sprintf(query, BcaID)
	//fmt.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return returnErrListPromo()
	}
	defer rows.Close()

	for rows.Next() {
		single := response.Promotion{}
		rows.Scan(&single.KodePromo, &single.Title, &single.Subtitle, &single.Description, &single.Minimum, &single.Cashback, &single.DateAvailable)
		listPromotions = append(listPromotions, single)
	}

	var listObjectives []response.Objectives
	dat, err = ioutil.ReadFile(dir + "/query/InquiryObjectives.query")
	if err != nil {
		fmt.Println(err)
		return returnErrListPromo()
	}
	query = string(dat)
	query = fmt.Sprintf(query, BcaID)
	//fmt.Println(query)
	rows, err = db.Query(query)
	if err != nil {
		fmt.Println(err)
		return returnErrListPromo()
	}
	defer rows.Close()

	for rows.Next() {
		single := response.Objectives{}
		rows.Scan(&single.KodePromo, &single.Title, &single.Subtitle, &single.Description, &single.Current, single.Target, &single.Cashback, &single.DateAvailable)
		listObjectives = append(listObjectives, single)
	}
	outputSchema.Objectives = listObjectives
	outputSchema.Promotion = listPromotions
	return outputSchema
}

func (dao *promoActivateDao) ClaimPromo(req request.ActivatePromoRequest) string {

	conn := dbconnection.New()
	db := conn.GetConnection()
	var result string

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `BEGIN SP_CLAIM_PROMO_AKUMULASI('%s','%s','SYSTEM API',:1 ); END;`
	query = fmt.Sprintf(query, req.BcaID, req.KodePromo)
	if _, err := db.ExecContext(ctx, query, sql.Out{Dest: &result}); err != nil {
		log.Printf("Error running %q: %+v", query, err)
		return "TIME OUT"
	}

	return result
}
