package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

// 定义一个全局对象db
var db *sql.DB

func executeSql(conn *sql.DB, executeSql string) error {
	res, err := conn.Exec(executeSql)
	if err != nil {
		return err
	}
	affectNum, err := res.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Printf("affected result: %d \n ", affectNum)
	return nil
}

func usage() {
	fmt.Println()
	fmt.Println("crud")
	fmt.Println("  a complete code of create table insert, update, query and delete then drop table.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println(`  curd -server server_url -sql sql`)
	flag.PrintDefaults()
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println(`  crud -server "ecommapp:ecomm123@tcp(192.168.127.43:3306)/ecomm?charset=utf8mb4&parseTime=True" -sql "select * from dual"`)
	fmt.Println()
}


func main() {

	var (
		server  string
		execSql string
	)

	flag.StringVar(&server, "server", "", "Server's URL, ecommapp:ecomm123@tcp(192.168.127.43:3306)/ecomm?charset=utf8mb4&parseTime=True")
	flag.StringVar(&execSql, "sql", "select * from dual", "execute sql")
	flag.Parse()

	connStr := os.ExpandEnv(server)
	if connStr == "" {
		fmt.Println("Missing -server option")
		usage()
		os.Exit(1)
	}
	fmt.Println("Connection string: ", connStr)
	fmt.Println("execute sql: ", execSql)
	conn, err := sql.Open("mysql", server)
	if err != nil {
		fmt.Println("Can't open the driver: ", err)
		return
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			fmt.Println("Can't close connection: ", err)
		}
	}()

	err = conn.Ping()
	if err != nil {
		fmt.Println("Can't ping connection: ", err)
		return
	}

	err = executeSql(conn, execSql)
	if err != nil {
		fmt.Println("execute sql error: ", err)
	}
}


// https://zhuanlan.zhihu.com/p/442534091