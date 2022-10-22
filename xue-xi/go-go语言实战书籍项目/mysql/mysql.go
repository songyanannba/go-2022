package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"strconv"
)


//CREATE TABLE `syn` (
//`id` int unsigned NOT NULL AUTO_INCREMENT,
//`name` varchar(20) NOT NULL,
//`age` int unsigned DEFAULT NULL,
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3;


type User struct {
	Id int
	Name string
	Age int
}

var db *sql.DB
var err error

func init() {
	db, err = sql.Open("mysql", "root:003416nba@tcp(127.0.0.1:3306)/syn?charset=utf8")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}



var tpl = `<html>
<head>
<title></title>
</head>
<body>
<form action="/info" method="post">
	用户名:<input type="text" name="name">
	年龄:<input type="text" name="age">
	<input type="submit" value="提交">
</form>
</body>
</html>`

func submitForm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	var t *template.Template
	t = template.New("Products") //创建一个模板
	t, _ = t.Parse(tpl)
	log.Println(t.Execute(w, nil))
}


func store(user User) {
	//插入数据
	stmt, err := db.Prepare("INSERT INTO user SET name=?,age=?")
	res, err := stmt.Exec(user.Name, user.Age)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Printf("last insert id is: %d\n", id)
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	//请求的是登录数据，那么执行登录的逻辑判断
	_ = r.ParseForm()
	if r.Method == "POST" {
		age, _ := strconv.Atoi(r.Form.Get("age"))
		user1 := User{Name: r.Form.Get("name"), Age: age}
		store(user1)
		fmt.Fprintf(w, " %v", queryByName("syn")) //这个写入到w的是输出到客户端的
	}
}

func queryByName(name string) User {
	user := User{}
	stmt, err := db.Prepare("select * from user where name=?")
	checkErr(err)

	rows, _ := stmt.Query(name)

	fmt.Println("\nafter deleting records: ")
	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		checkErr(err)
		fmt.Printf("[%d, %s, %s, %s]\n", id, name, age)
		user = User{id, name, age}
		break
	}
	return user
}

func main() {

	http.HandleFunc("/form", submitForm)     //设置访问的路由
	http.HandleFunc("/info", userInfo)       //设置访问的路由

	http.ListenAndServe(":8008", nil)
}


