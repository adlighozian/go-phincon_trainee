package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type contacts struct {
	Id      int    `gorm:"primaryKey;column:id"`
	Name    string `gorm:"column:name"`
	No_telp string `gorm:"column:no_telp"`
}

// one to one
type Books struct {
	gorm.Model
	Name         string `gorm:"column:book_name"`
	NumberOfPage string `gorm:"column:pages"`
	PublisherId  uint
	Publishers   Publisher `gorm:"foreignKey:PublisherId; constraint:OnUpdate:CASCADE, OnDelete:CASECADE"`
}

type Publisher struct {
	gorm.Model
	PublisherName string `gorm:"column:name"`
	PublisherAddr string `gorm:"column:address"`
}

// one to many
type Booksm struct {
	gorm.Model
	Name         string `gorm:"column:book_name"`
	NumberOfPage string `gorm:"column:pages"`
	PublisherId  uint
}

type Publisherm struct {
	gorm.Model
	Np            string
	PublisherName string  `gorm:"column:name"`
	PublisherAddr string  `gorm:"column:address"`
	Books         []Books `gorm:"foreignKey:PublisherId; constraint:OnUpdate:CASCADE, OnDelete:CASECADE; Reference: Np"`
}

// many to many
type Songs struct {
	gorm.Model
	Title  string    `gorm:"column:judul"`
	Genres []*Genres `gorm:"many2many:song_genres;"`
}

type Genres struct {
	gorm.Model
	Name  string   `gorm:"column:nama"`
	Songs []*Songs `gorm:"many2many:song_genres;"`
}

func main() {

	dsn := "root:@tcp(localhost:3306)/bootcamp?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		PrepareStmt:            false,
	})
	if err != nil {
		panic(err)
	}
	db.Begin()
	log.Println("server start")
	// var temp *[]contacts

	// var contact contacts

	// var temp *Books
	// db.InnerJoins("Publishers").Find(&temp)

	// var tempm *Publisherm
	// db.InnerJoins("Books").Find(&tempm)

	res := db.Model(&Songs{}).Preload("Genres").Find([]Songs)

	// dbConn.Where(contacts{Id: 1}).Find(&temp)

	// log.Println(temp)
	// contact := contacts{
	// 	Id:      2,
	// 	Name:    "asd",
	// 	No_telp: "535351",
	// }

	// res := db.Select("name", "no_telp").Create(contact)
	// err = res.Error
	// if err != nil {
	// 	panic(err)
	// }

	// res := db.Delete(&contact)
	// err = res.Error
	// if err != nil {
	// 	panic(err)
	// }
}
