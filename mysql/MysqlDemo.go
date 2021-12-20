package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	constr := "qms:Myegoo@3466@tcp(192.168.25.191:3306)/smart_ant"
	//mysql.DeregisterReaderHandler(constr)

	db, err := sql.Open("mysql", constr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	// 创建表 _不需要他的值
	//_, err = db.Exec("create table person (id int auto_increment primary key ,name varchar(12) not null,age int default 1 );")

	// 新增一条数据
	//_, err = db.Exec("insert into person (name,age) "+" values (?,?);", "TOM", 18)

	// 查询数据
	rows, err := db.Query("select id,age,name from person")

scan6:
	next := rows.Next()
	if next {
		person := new(Person)
		err := rows.Scan(&person.Id, &person.Age, &person.Name)
		if err !=nil {
			log.Println(err.Error())
		}
		fmt.Println(person.Id,person.Age,person.Name)
		goto scan6
	}


	//result, err := db.Exec("select * from t_role")

	if err != nil {
		log.Fatalln(err.Error())
		return
	} else {
		fmt.Print("新增成功！")
	}

}

type Person struct {
	Id int
	Name string
	Age int

}
