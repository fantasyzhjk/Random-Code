package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type res struct {
	Hitokoto string `json:"hitokoto"`
	From     string `json:"from"`
}

func init() {
	sqlInit()
}

func clear() {
	clear := make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear[runtime.GOOS]()
}

func get(c string) {
	clear()
	params := url.Values{}
	URL, _ := url.Parse("https://v1.hitokoto.cn/")
	params.Set("c", c)
	URL.RawQuery = params.Encode()
	// fmt.Println(URL.String())
	hitokoto, err := http.Get(URL.String())
	if err != nil {
		fmt.Println("获取失败惹qwq...")
		return
	}
	defer hitokoto.Body.Close()
	body, _ := ioutil.ReadAll(hitokoto.Body)
	fmt.Println(string(body))
	var result res
	_ = json.Unmarshal(body, &result)
	fmt.Printf("[%v] - %v\n", result.Hitokoto, result.From)
}

func sqlInit() {
	db, err := sql.Open("sqlite3", "./hitokoto.db")
	if err != nil {
		panic(err)
	}
	sqlTable := `
    CREATE TABLE IF NOT EXISTS Hitokoto(
        uid INTEGER PRIMARY KEY AUTOINCREMENT,
        hitokoto VARCHAR(64) NULL,
        author VARCHAR(64) NULL,
        created VARCHAR(64) NULL
    );
    `
	db.Exec(sqlTable)
}

func sqlInsert(hitokoto, author string) {
	created := time.Now().Format("2006-01-02")
	// fmt.Println(created)
	db, err := sql.Open("sqlite3", "./hitokoto.db")
	if err != nil {
		fmt.Println("导入失败惹qwq...")
		return
	}
	stmt, err := db.Prepare("INSERT INTO Hitokoto(hitokoto, author, created) values(?,?,?)")
	if err != nil {
		fmt.Println("导入失败惹qwq...")
		return
	}
	res, err := stmt.Exec(hitokoto, author, created)
	if err != nil {
		fmt.Println("导入失败惹qwq...")
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("导入失败惹qwq...")
		return
	}
	fmt.Println("导入完成喵~ uid为：", id)
}

func sqlQuery() {
	clear()
	db, err := sql.Open("sqlite3", "./hitokoto.db")
	rows, err := db.Query("SELECT * FROM Hitokoto")
	if err != nil {
		fmt.Println("获取失败惹qwq...")
		return
	}
	var uid int
	var hitokoto string
	var author string
	var created string
	for rows.Next() {
		err = rows.Scan(&uid, &hitokoto, &author, &created)
		if err != nil {
			fmt.Println("获取失败惹qwq...")
			return
		}
		fmt.Printf("[%v] - %v (%v/%v)\n", hitokoto, author, created, uid)
	}
	rows.Close()
	if err != nil {
		fmt.Println("导入失败惹qwq...")
		return
	}
}

func sqlRead(id int) {
	clear()
	db, err := sql.Open("sqlite3", "./hitokoto.db")
	var hitokoto string
	var author string
	if id == 0 {
		err = db.QueryRow("SELECT hitokoto,author FROM Hitokoto ORDER BY RANDOM() limit 1").Scan(&hitokoto, &author)
	} else {
		err = db.QueryRow("SELECT hitokoto,author FROM Hitokoto WHERE uid=?", id).Scan(&hitokoto, &author)
	}
	if err != nil {
		fmt.Println("获取失败惹qwq...")
		return
	}
	fmt.Printf("[%v] - %v\n", hitokoto, author)
}

func randString() string {
	rand.Seed(time.Now().UnixNano())
	return string(97 + rand.Intn(109-97))
}

func main() {
	for {
		clear()
		fmt.Println(`
-------------------
      食用方法
    1随机获取一言
    2指定类型一言
   3导入自定义一言
   4查询所有自定义
     （默认为1）
-------------------`)
		var i uint8
		fmt.Println("请输入:")
		fmt.Scanln(&i)
		switch i {
		case 1:
			r := randString()
			if r == "m" {
				sqlRead(0)
			} else {
				get(r)
			}
			break
		case 2:
			var r string
			fmt.Println("请输入一言类型:")
			fmt.Scanln(&r)
			if r == "m" {
				sqlRead(0)
			} else {
				get(r)
			}
			break
		case 3:
			var hitokoto, author string
			fmt.Println("请输入自定义一言:")
			fmt.Scanln(&hitokoto)
			fmt.Println("请输入作者/来源:")
			fmt.Scanln(&author)
			sqlInsert(hitokoto, author)
			break
		case 4:
			sqlQuery()
			break
		default:
			r := randString()
			if r == "m" {
				sqlRead(0)
			} else {
				get(r)
			}
			break
		}
		fmt.Println("按回车键继续")
		fmt.Scanln()
	}
}
