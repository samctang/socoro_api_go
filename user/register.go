package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/samctang/socoro_api/database"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register(c echo.Context) error {
	ctx := context.Background()

	//begin transaction
	tx, err := database.Conn.Begin(ctx)
	if err != nil{
		log.Fatal("Begin transaction failed: ", err)
	}

	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, "Cannot bind data")
	}

	user := DTO{
		Name: u.Name,
		Email: u.Email,
		Role: u.Role,
		isAdmin: true,
	}

	//check if email exists in db
	row := tx.QueryRow(ctx,`select id from user_acc where email=$1;`, u.Email)
	err = row.Scan(&u.Id)
	if err == nil {
		log.Fatal("Email already exists: ", err)
		return c.JSON(http.StatusBadRequest, "Email already exists")
	}

	user.Uuid = uuid.New().String()

	user.Password, err = GeneratePasswordHash(u.Password)
	if err != nil {
		log.Fatal("Password hash failed: ", err)
	}

	//insert into user_acc
	_, err = tx.Exec(ctx, `insert into user_acc (uuid,name,email,password,role,isAdmin) values($1,$2,$3,$4,$5,$6) returning email`, user.Uuid, user.Name, user.Email, user.Password, user.Role, user.isAdmin)
	if err != nil {
		log.Fatalln("Account insert failed: ", err)
	}

	//commit transaction
	if err := tx.Commit(ctx); err != nil {
		log.Fatalln("Commit tx failed: ", err)
	}

	return c.JSON(http.StatusOK, "Account created")
}

func GeneratePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}