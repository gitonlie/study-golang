package scripts

import (
	"fmt"
	"gorm.io/gorm"
)

type Student struct {
	ID    int
	Name  string
	Age   int
	Grade string
}

func RunStudent(db *gorm.DB) {
	db.AutoMigrate(&Student{})

	db.Debug().Create(&Student{Name: "张三", Age: 20, Grade: "三年级"})

	var student []Student
	db.Debug().Where("age > ?", 18).Find(&student)
	fmt.Println(student)

	var student2 Student
	student2.Name = "张三"
	db.Debug().Where(&student2).Updates(&Student{Name: "李四"})

	db.Debug().Where("age < ?", 15).Delete(&Student{})
}
