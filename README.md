# oracle client
https://zhuanlan.zhihu.com/p/371544255  
https://github.com/mattn/go-oci8
https://github.com/godror/godror

试试看go-ora？ 免客户端的
https://github.com/sijms/go-ora

本文採用go-ora

## 1 dev
```shell
 go run .\cmd\main.go -server oracle://cdrdev:cdr@192.168.100.177:1522/cdrtest -sql "select * from dual"
 go run .\cmd\main.go -server oracle://cdrdev:cdr@192.168.100.177:1522/cdrtest -sql "CREATE TABLE VIN_DEPT_ZW(EPTNO NUMBER(2) CONSTRAINT PK_DEPT PRIMARY KEY, DNAME VARCHAR2(14))"
 
 PS D:\code\go\go-oraclient> go run .\cmd\main.go -server oracle://cdrdev:cdr@192.168.100.177:1522/cdrtest -sql "insert into VIN_DEPT_ZW(EPTNO,Dname) values (3,'vincentzou')"
Connection string:  oracle://cdrdev:cdr@192.168.100.177:1522/cdrtest
execute sql:  insert into VIN_DEPT_ZW(EPTNO,Dname) values (3,'vincentzou')
affected result: 1

PS D:\code\go\go-oraclient> go run .\cmd\main.go -server oracle://cdrdev:cdr@192.168.100.177:1522/cdrtest -sql "delete from VIN_DEPT_ZW where eptno = 3"
Connection string:  oracle://cdrdev:cdr@192.168.100.177:1522/cdrtest
execute sql:  delete from VIN_DEPT_ZW where eptno = 3
affected result: 1

PS D:\code\go\go-oraclient> go run .\cmd\main.go -server oracle://cdrdev:cdr@192.168.100.177:1522/cdrtest -sql "select * from VIN_DEPT_ZW"
Connection string:  oracle://cdrdev:cdr@192.168.100.177:1522/cdrtest
execute sql:  select * from VIN_DEPT_ZW
affected result: 0

PS D:\code\go\go-oraclient> go run .\cmd\main.go -server oracle://cdrdev:cdr@192.168.100.177:1522/cdrtest -sql "DROP TABLE VIN_DEPT_ZW"   
Connection string:  oracle://cdrdev:cdr@192.168.100.177:1522/cdrtest
execute sql:  DROP TABLE VIN_DEPT_ZW
affected result: 0 
```

## 2. compile
```shell
env GOOS=linux GOARCH=amd64 go build -o oractl main.go 
go build -o oractl main.go 
oractl -uname cdrdev -password cdr -host 192.168.100.177 -port 1522 -schema cdrtest
```

## 3. make image
```shell

docker build --target oraclectl -t blueskyxi3/oraclectl:v0.0.1 -f Dockerfile .
docker build --target mysqlctl -t blueskyxi3/mysqlctl:v0.0.1 -f Dockerfile .
docker push blueskyxi3/oraclectl:v0.0.1
docker push  blueskyxi3/mysqlctl:v0.0.1
date

```

# mysql
```shell
go run .\mysql\main.go -server "ecommapp:ecomm123@tcp(192.168.127.43:3306)/ecomm?charset=utf8mb4&parseTime=True" -sql "CREATE TABLE `testzw` (`T_ID` int(8) ,`T_NAME` varchar(8) ) ENGINE=Inn
oDB DEFAULT CHARSET=utf8 "  

go run .\mysql\main.go -server "ecommapp:ecomm123@tcp(192.168.127.43:3306)/ecomm?charset=utf8mb4&parseTime=True" -sql "insert into `testzw`(T_ID,T_NAME) values (3,'vincentzou')"
```
