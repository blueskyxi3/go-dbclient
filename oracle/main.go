package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/sijms/go-ora/v2"
	"os"
)

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
	fmt.Println(`  crud -server "oracle://user:pass@server:1521/service_name" -sql "select * from dual"`)
	fmt.Println()
}

func main() {
	var (
		server  string
		execSql string
	)

	flag.StringVar(&server, "server", "", "Server's URL, oracle://user:pass@server:port/service_name")
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
	conn, err := sql.Open("oracle", server)
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
