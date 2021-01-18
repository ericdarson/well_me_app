package dao

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"
	"widget_nab_service/dbconnection"
	"widget_nab_service/entity/response"

	_ "github.com/godror/godror"
)

type TargetSimulationPlannerDao interface {
	Simulation(bcaId string, target float64, dueDate string, periodic string) ([]response.TargetSimulationDetail, string)
}

type targetSimulationPlannerDao struct {
	temp string
}

func NewTargetSimulationPlanner() TargetSimulationPlannerDao {
	return &targetSimulationPlannerDao{}
}

//bcaid ,target,targettujuan, dan return total tabungan perbulan
func (dao *targetSimulationPlannerDao) Simulation(bcaId string, target float64, dueDate string, periodic string) ([]response.TargetSimulationDetail, string) {
	var targetSimulation []response.TargetSimulationDetail

	conn := dbconnection.New()
	db := conn.GetConnection()
	layout := "02-01-2006"
	t, err := time.Parse(layout, dueDate)
	dtstr2 := t.Format("02-Jan-2006")
	if err != nil {
		fmt.Println(err)
		targetSimulation = append(targetSimulation, response.TargetSimulationDetail{
			NominalInvestasi: "-1",
			Date:             "-1",
		})
		return targetSimulation, "-1"
	}
	var result string
	datediff := t.Sub(time.Now()).Hours() / 24
	repeatInvestment := 0.0
	if periodic == "Monthly" {
		repeatInvestment = math.Ceil(datediff / 30)
	} else if periodic == "Weekly" {
		repeatInvestment = math.Ceil(datediff / 7)
	} else if periodic == "Yearly" {
		repeatInvestment = math.Ceil(datediff / 365)
	}

	query := `BEGIN SP_PERKIRAAN_KEUNTUNGAN_SIMULATION('` + bcaId + `','` + dtstr2 + `',:1); END;`
	if _, err := db.Exec(query, sql.Out{Dest: &result}); err != nil || result == "-1" {
		log.Printf("Error running %q: %+v", query, err)
		targetSimulation = append(targetSimulation, response.TargetSimulationDetail{
			NominalInvestasi: "-1",
			Date:             "-1",
		})
		return targetSimulation, "-1"
	} else {
		y := int(repeatInvestment)
		res, err := strconv.ParseFloat(result, 64)
		if err != nil {
			fmt.Println(err)
			targetSimulation = append([]response.TargetSimulationDetail{}, response.TargetSimulationDetail{
				NominalInvestasi: "-1",
				Date:             "-1",
			})
			return targetSimulation, "-1"
		}
		oneTimeProfit := res / repeatInvestment

		targetBefore := 0.0

		targetBase := target / (repeatInvestment)

		for i := 0; i < y; i++ {
			targetPiece := 0.0
			if i == y-1 {
				targetPiece = targetBase + (targetBase * oneTimeProfit / 100) + targetBefore
			} else {
				targetPiece = targetBase - (targetBase * oneTimeProfit / 100) + targetBefore
			}

			//	fmt.Println(i, "Targetpiece: ", targetPiece, "target: ", target, "repeatInvestment", repeatInvestment)
			if periodic == "Monthly" {
				targetSimulation = append(targetSimulation, response.TargetSimulationDetail{
					NominalTanpaInvestasi: strconv.FormatFloat(targetBase*(float64(i)+1), 'f', 2, 64),
					NominalInvestasi:      strconv.FormatFloat(targetPiece, 'f', 2, 64),
					Date:                  time.Now().AddDate(0, i, 0).Month().String()[:3] + " " + strconv.Itoa(time.Now().AddDate(int(math.Floor(float64(i)/12)), 0, 0).Year())[2:],
				})

				fmt.Println(time.Now().AddDate(0, i, 0).Month().String()[:3])
			} else if periodic == "Weekly" {
				targetSimulation = append(targetSimulation, response.TargetSimulationDetail{
					NominalTanpaInvestasi: strconv.FormatFloat(targetBase*(float64(i)+1), 'f', 2, 64),
					NominalInvestasi:      strconv.FormatFloat(targetPiece, 'f', 2, 64),
					Date:                  time.Now().AddDate(0, 0, i*7).Format("02 Jan"),
				})
			} else if periodic == "Yearly" {
				targetSimulation = append(targetSimulation, response.TargetSimulationDetail{
					NominalTanpaInvestasi: strconv.FormatFloat(targetBase*(float64(i)+1), 'f', 2, 64),
					NominalInvestasi:      strconv.FormatFloat(targetPiece, 'f', 2, 64),
					Date:                  strconv.Itoa(time.Now().AddDate(i, 0, 0).Year()),
				})
			}

			targetBefore = targetPiece + (targetPiece * oneTimeProfit / 100)

		}

		return targetSimulation, strconv.FormatFloat(targetBase-(targetBase*oneTimeProfit/100), 'f', 2, 64)
	}
	//listNAB = append(listNAB, single)

}
