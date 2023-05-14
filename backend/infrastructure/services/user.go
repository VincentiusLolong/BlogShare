package services

import (
	"fmt"
	"mailinglist/configs/db"
	"mailinglist/infrastructure/models"
	"mailinglist/infrastructure/security"
	"regexp"

	"github.com/google/uuid"
)

type Services interface {
	SignUp(user models.User) (string, error)
	Signin(users models.Login) (uuid.UUID, error)
	HomepageUsers(ids string) (map[string]interface{}, error)

	CreateContent(contents models.Contents, userid string) (string, error)
	GetUserContent(userid string) ([]models.Userconten, error)
	EditContent(contents models.GetContent, userid string) (string, error)
	// Logout() (err error)
}

type services struct {
	psql db.Postgre
}

func New(psql db.Postgre) Services {
	return &services{psql: psql}
}

// func formats(data string) (string, error) {
// 	layout := "02-01-2006"
// 	t, err := time.Parse(layout, data)
// 	if err != nil {
// 		return "", err
// 	}
// 	return t.Format("2006-01-02"), nil
// }

func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}

func (s *services) SignUp(user models.User) (string, error) {
	// birthdate, err := formats(user.Birth_date)
	// if err != nil {
	// 	return "", err
	// }
	hash, errs := security.HashPassword(user.Password)
	if errs != nil {
		return "", errs
	}
	data := fmt.Sprintf("INSERT INTO Account(email, username, password) VALUES ('%v','%v','%v')", user.Email, user.Username, hash)
	row, errdata := s.psql.PGExecQuery(data)
	if errs != nil {
		return "", errdata
	}
	return row, nil
}

func (s *services) Signin(users models.Login) (uuid.UUID, error) {
	var pass string
	var Account_id uuid.UUID
	var getppasserr error
	var getuseriderr error

	if isValidEmail(users.Email) {
		data := fmt.Sprintf("select password from account WHERE email= '%v'", users.Email)
		user, _ := s.psql.PGRowQuery(data, true)
		getppasserr = user.Scan(&pass)
	} else {
		data := fmt.Sprintf("select password from account WHERE username = '%v'", users.Username)
		user, _ := s.psql.PGRowQuery(data, true)
		getppasserr = user.Scan(&pass)
	}

	if getppasserr != nil {
		return uuid.Nil, getppasserr
	}

	err := security.VerifyPassword(pass, users.Password)
	if err != nil {
		return uuid.Nil, err
	} else {
		if isValidEmail(users.Email) {
			data2 := fmt.Sprintf("select account_id from account WHERE email = '%v'", users.Email)
			user, _ := s.psql.PGRowQuery(data2, true)
			getuseriderr = user.Scan(&Account_id)
		} else {
			data2 := fmt.Sprintf("select account_id from account WHERE username = '%v'", users.Username)
			user, _ := s.psql.PGRowQuery(data2, true)
			getuseriderr = user.Scan(&Account_id)
		}
	}

	if getuseriderr != nil {
		return uuid.Nil, err
	}

	return Account_id, nil
}

func (s *services) HomepageUsers(ids string) (map[string]interface{}, error) {
	var public models.PublicUser

	data := fmt.Sprintf("select account_create, username from account where account_id ='%v'", ids)
	user, _ := s.psql.PGRowQuery(data, true)
	getusererr := user.Scan(&public.Account_create, &public.Username)
	if getusererr != nil {
		return nil, getusererr
	}

	converts := fmt.Sprintf("%v", public.Account_create.Format("02-01-2006"))
	alldata := map[string]interface{}{
		"account_created": converts,
		"username":        public.Username,
	}

	return alldata, nil
}

// func (s *services) Logout() (err error) {
// 	err = s.psql.StopPsqlConnection()
// 	return
// }
