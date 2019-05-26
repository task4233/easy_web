 package router

import (
	"net/http"

	// mysql
  _ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"

	
	"github.com/labstack/echo"
)

type (
	MakeProblem struct {
		Id int `db: "id" json:"id"`
		Ans string `db:"ans" json:"ans"`
		Score int `db:"score" json:"score"`
		Silverbullet string `db:"silverbullet" json:"silverbullet`
	}

	UserAns struct {
		Id int `db: "id" json:"id"`
		Ans string `db:"ans" json:"ans"`
	}
	
	Problem struct {
		Id int `db: "id" json:"id"`
		Ans string `db:"ans" json:"ans"`
		Score int `db:"score" json:"score"`
	}

	ResponseProblem struct {
		Id int `json:"id"`
		Result bool `json:"result"`
		Score int `json:"score"`
	}

	FlagData struct {
		Id int `db: "id" json:"id"`
		Flag string `db:"flag" json:"flag"`
		Silverbullet string `db:"silverbullet" json:"pass"`
	}
)

var (
	problemtable = "probleminfo"
	flagtable = "flaginfo"
	seq = 1
	// dbr.Open("mysql","[user]:[pass]&tcp([hostname])/[dbname]")
	conn, err = dbr.Open("mysql", "root:R;5aritz@tcp(127.0.0.1:3306)/workout", nil)
	sess = conn.NewSession(nil)
)

/*
Routers handler
args
e: pointer of echo structure
return
none
*/
func Routes(e *echo.Echo) {
	e.POST("/api/v1/problems/check", checkAns)
	e.POST("/api/v1/problems/check/new", makeAns)
	e.POST("/api/v1/problems/getFlag", getFlag)
	e.POST("/api/v1/problems/getFlag/new", makeFlag)
}

func makeFlag(c echo.Context) error {
	flag := new(FlagData)
	if err := c.Bind(flag); err != nil {
		return err
	}
	
	if flag.Silverbullet != "S1LVERBULLET" {
		return c.NoContent(http.StatusBadRequest) 
	}
	sess.InsertInto(flagtable).Columns("id","flag").Values(flag.Id, flag.Flag).Exec()
	
	return c.NoContent(http.StatusOK)
}


func getFlag(c echo.Context) error {
	flag := new(FlagData)
	if err := c.Bind(flag); err != nil {
		return err
	}
	
	if flag.Silverbullet != "S1LVERBULLET" {
		return c.NoContent(http.StatusBadRequest) 
	}

	response := new(FlagData)
	sess.Select("*").From(flagtable).Where("id=?", flag.Id).Load(&response)
	return c.JSON(http.StatusOK, response)
}

func makeAns(c echo.Context) error {
	problem := new(MakeProblem)
	if err := c.Bind(problem); err != nil {
		return err
	}
	if problem.Silverbullet != "R;5aritz" {
		return c.NoContent(http.StatusBadRequest)
	}

	sess.InsertInto(problemtable).Columns("id","ans","score").Values(problem.Id, problem.Ans, problem.Score).Exec()
	return c.NoContent(http.StatusOK)
}

func checkAns(c echo.Context) error {
	user_ans := new(UserAns)
	if err := c.Bind(user_ans); err != nil {
		return err
	}
	
	id := user_ans.Id
	problem := new(Problem)
	sess.Select("*").From(problemtable).Where("id=?", id).Load(&problem)

	res := problem.Ans == user_ans.Ans
	response := ResponseProblem {
		Id: problem.Id,
		Result: res,
		Score: problem.Score,
	}
	
	return c.JSON(http.StatusOK, response)
}
