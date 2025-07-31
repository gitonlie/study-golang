package scripts

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     int
}

// 创建表
func createEmployeeTable(db *sqlx.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS employees (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(80) NOT NULL,
		department VARCHAR(100),
		salary INT
	)`
	_, err := db.Exec(schema)
	return err
}

// 创建数据 (Create)
func createEmployee(db *sqlx.DB, emp *Employee) (int64, error) {
	res, err := db.NamedExec(
		`INSERT INTO employees (name, department, salary) VALUES (:name, :department, :salary)`, emp)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func RunExtend() {
	db, err := sqlx.Open("mysql", "root:admin123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True")
	if err != nil {
		fmt.Println("连接数据库异常:", err)
		return
	}
	defer db.Close()

	//先建表
	createEmployeeTable(db)

	//插入数据
	data := []Employee{
		{Name: "小李", Department: "技术部", Salary: 6000},
		{Name: "小王", Department: "技术部", Salary: 8000},
		{Name: "小张", Department: "技术部", Salary: 9000}}

	for _, v := range data {
		idKey, _ := createEmployee(db, &v)
		fmt.Println("LastInsertId:", idKey)
	}

	var emp []Employee
	sql1 := "select id, name, department,salary from employees where department = ?"
	db.Select(&emp, sql1, "技术部")
	fmt.Println(emp)

	var empMax Employee
	sql2 := "select id, name, department,salary from employees where salary = (select max(salary) from employees) limit 1"
	db.Get(&empMax, sql2)
	fmt.Println(empMax)

}
