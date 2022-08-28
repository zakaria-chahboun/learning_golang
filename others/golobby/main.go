package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golobby/orm"
	_ "github.com/mutecomm/go-sqlcipher"
	//_ "github.com/mattn/go-sqlite3"
)

// sqlite connection string
const (
	key    = "123456"
	dbname = "./database/database_e.db"
	dsn    = dbname + "?_pragma_key=" + key + "&_pragma_cipher_page_size=4096"
)

type User struct {
	ID   int64
	Name string
	orm.Timestamps
}

type Laptop struct {
	ID     int64
	Name   string
	Model  string
	UserID int64
	orm.Timestamps
}

// It will be called by ORM to setup entity.
func (u User) ConfigureEntity(e *orm.EntityConfigurator) {
	// Specify related database table for the entity.
	e.Table("users").HasMany(&Laptop{}, orm.HasManyConfig{})

}
func (l Laptop) ConfigureEntity(e *orm.EntityConfigurator) {
	// Specify related database table for the entity.
	e.Table("laptops").BelongsTo(&User{}, orm.BelongsToConfig{})

}

func main() {

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatalln("[1]", err)
	}
	defer db.Close()

	err = createTables(db)
	if err != nil {
		log.Fatalln("cannot -- ", err)
	}

	err = orm.SetupConnections(orm.ConnectionConfig{
		Name:                "default",
		DB:                  db,
		Dialect:             orm.Dialects.SQLite3,
		Entities:            []orm.Entity{&User{}, Laptop{}},
		DatabaseValidations: true,
	})

	if err != nil {
		log.Fatalln("[2]", err)
	}

	//user, err := orm.Query[User]().WherePK(2).Get()

	//laptops, err := orm.HasMany[Laptop](User{ID: 4}).All()
	//user, err := orm.BelongsTo[User](Laptop{ID: 2}).Get()

	user := User{Name: "Sarah"}

	err = orm.Save(&user)
	if err != nil {
		log.Fatalln("[3]", err)
	}

	fmt.Println("new user:", user.ID, user.Name)

	laptops := []orm.Entity{Laptop{Name: "Elite book", Model: "Hp"}, Laptop{Name: "Sandisk", Model: "Toshiba"}}
	err = orm.Add(user, laptops...)
	if err != nil {
		panic(err)
	}

	fmt.Println("Okay!")
}

func createTables(db *sql.DB) (err error) {
	_, err = db.Exec(`

	DROP users;
	CREATE TABLE users (
    id         INTEGER  PRIMARY KEY AUTOINCREMENT
                        NOT NULL,
    name       VARCHAR,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME
);

DROP laptops;
CREATE TABLE laptops (
    id         INTEGER  PRIMARY KEY AUTOINCREMENT
                        NOT NULL,
    name       VARCHAR,
    model      VARCHAR,
    user_id    INTEGER  REFERENCES users (id) ON DELETE CASCADE
                                              ON UPDATE CASCADE,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME
);
`)
	return
}
