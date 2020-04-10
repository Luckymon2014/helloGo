package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*
	go没有数据库驱动，而是定义了标准接口，开发者可以根据接口来开发数据库驱动
	优点：数据库迁移时不用做任何修改
*/
func main() {
	// 打开一个注册过的数据库驱动，引用go-sql-driver注册的mysql驱动
	// 第二个参数是DSN(Data Source Name)，是参数一的驱动定义的数据库链接信息
	// go-sql-driver的DSN格式: user@unix(/path/to/socket)/dbname?charset=utf8
	db, err := sql.Open(
		"mysql",
		"user:password@tcp(ip:port)/dbname?charset=utf8")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	checkErr(err)
	res, err := stmt.Exec("test", "department", "2012-12-09")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)
	res, err = stmt.Exec("test", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)
	res, err = stmt.Exec(id)
	checkErr(err)
	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/*


标准接口：

sql.Register： 在init函数里调用 Register(name string, driver driver.Driver) 来完成驱动注册

type Driver interface {
	Open(name string) (Conn, error)
}

type Conn interface {
	Prepare(query string) (Stmt, error)
	Close() error
	Begin() (Tx, error)
}

type Stmt interface {
	Close() error
	NumInput() int
	Exec(args []Value) (Result, error)
	Query(args []Value) (Rows, error)
}

type Tx interface {
	Commit() error
	Rollback() error
}

type Execer interface {
	Exec(query string, args []Value) (Result, error)
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Rows interface {
	Columns() []string
	Close() error
	Next(dest []Value) error
}

type RowsAffected int64
func (RowsAffected) LastInsertId() (int64, error)
func (v RowsAffected) RowsAffected() (int64, error)

type Value interface{}

type ValueConverter interface {
	ConvertValue(v interface{}) (Value, error)
}

type Valuer interface {
	Value() (Value, error)
}

type DB struct {
	driver 	 driver.Driver
	dsn    	 string
	mu       sync.Mutex // protects freeConn and closed
	freeConn []driver.Conn
	closed   bool
}
*/
