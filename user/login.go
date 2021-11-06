package user

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/samctang/socoro_api/database"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Login(c echo.Context) error {
	ctx := context.Background()

	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, "Cannot bind data")
	}

	user := DTO{
		Email: u.Email,
		Password: u.Password,
	}

	//check if email exists in db
	row := database.Conn.QueryRow(ctx,`select email, password from user_acc where email=$1;`, u.Email)
	err := row.Scan(&u.Email, &u.Password)
	if err != nil {
		log.Fatal("Email does not exist: ", err)
		return c.JSON(http.StatusUnauthorized, "Email does not exist")
	}

	//check hash
	if check := CheckPasswordHash(user.Password, u.Password); !check{
		log.Fatal("Username or password incorrect: ", err)
		return c.JSON(http.StatusUnauthorized, "Username or password incorrect")
	}

	return c.JSON(http.StatusOK, "Authentication successful")
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
