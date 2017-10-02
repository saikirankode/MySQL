package main

import (
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

func main() {
	db, err := gorm.Open("mysql", "root:kiran@/kode")
	if err != nil {
		panic(err.Error())
	}
	//db.DB().Ping()
	//fmt.Println(err)

	defer db.Close()
	//db.SingularTable(true)
	db.DropTableIfExists(&Owner{}, &Book{}, &Author{})
	db.CreateTable(&Owner{}, &Book{})

	/*owner := Owner{
		FirstName : "sai kiran",
		LastName :"kode",
	}*/

	/*db.Create(&owner)    // creating
	owner.FirstName = "sairam"*/

	/*db.Save(&owner)  // updating
	var o Owner
	db.Debug().First(&o)
	fmt.Println(o)
	db.Delete(&owner)   // Deleting
	o = Owner{}
	db.Debug().First(&o)
	fmt.Println(o)*/
}


type Owner struct {
	gorm.Model
	FirstName string
	LastName string
	Books [] Book
}

type Book struct {
	gorm.Model
	Name string
	PublishDate time.Time
	OwnerID uint `sql:"index"`  // creating index
	Authors []Author `gorm:"many2many:books_authors"`  // relating books and author
}

type Author struct {
	gorm.Model
	FirstName string
	LastName string
}
/*func  ( o *Owner) TableName() string{    // changing the table name
	return "foo"
}*/
