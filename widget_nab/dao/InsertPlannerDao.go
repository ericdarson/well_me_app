package dao

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
	"widget_nab_service/dbconnection"

	_ "github.com/godror/godror"
)

type InsertPlannerDao interface {
	InsertPlanner(bcaId string, namaPlan string, goalAmount string, currentAmount string, periodic string, dueDate string, kategori string) string
}

type insertPlannerDao struct {
	temp string
}

func NewInsertPlannerDao() InsertPlannerDao {
	return &insertPlannerDao{}
}
func (dao *insertPlannerDao) InsertPlanner(bcaId string, namaPlan string, goalAmount string, currentAmount string, periodic string, dueDate string, kategori string) string {
	var detailInsertPlanner string
	conn := dbconnection.New()
	db := conn.GetConnection()

	var status string
	fmt.Println(goalAmount)
	puzzle_randomize := CreatePuzzleRandomize()
	query := `BEGIN SP_INSERT_PLANNER('` + bcaId + `','` + namaPlan + `','` + goalAmount + `','` + currentAmount + `','` + periodic + `','` + dueDate + `','` + bcaId + `','` + kategori + `','` + puzzle_randomize + `',:1); END;`
	if _, err := db.Exec(query, sql.Out{Dest: &status}); err != nil {
		log.Printf("Error running %q: %+v", query, err)
		return detailInsertPlanner
	} else {
		if status == "1" {
			detailInsertPlanner = "Sukses"
		} else {
			detailInsertPlanner = "Gagal"
		}
	}
	return detailInsertPlanner

}

func CreatePuzzleRandomize() string {
	const N = 9
	rng := NewUniqueRand(N)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = (rng.Int()) << 1
		a[i] = (a[i] / 2) + 1
	}
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	kategori := r.Intn(2)
	a = append([]int{kategori}, a...)
	puzzle_randomize := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), ""), "[]")
	return puzzle_randomize
}

type UniqueRand struct {
	generated map[int]bool
	rng       *rand.Rand
	scope     int
}

func NewUniqueRand(N int) *UniqueRand {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return &UniqueRand{
		generated: map[int]bool{},
		rng:       r1,
		scope:     N,
	}
}

func (u *UniqueRand) Int() int {
	if u.scope > 0 && len(u.generated) >= u.scope {
		return -1
	}
	for {
		var i int
		if u.scope > 0 {
			i = u.rng.Int() % u.scope
		} else {
			i = u.rng.Int()
		}
		if !u.generated[i] {
			u.generated[i] = true
			return i
		}
	}
}
