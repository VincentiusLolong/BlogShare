package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"mailinglist/infrastructure/models"

	"github.com/jackc/pgx/v5"
)

type Postgre interface {
	Validatorpsql()
	PGRowQuery(query string, choice bool) (pgx.Row, pgx.Rows)
	PGExecQuery(query string) (string, error)
	contexts() (context.Context, context.CancelFunc)
	StopPsqlConnection() error
}
type postgresql struct {
	pgsql *pgx.Conn
}

func New() Postgre {
	return &postgresql{}
}

// var pgsql *pgx.Conn

func custompostgreurl() (result string) {
	result = fmt.Sprintf("%s/%s?sslmode=disable", models.PostgreSetting.Connstr, models.PostgreSetting.DBname)
	return
}

func (p *postgresql) contexts() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx, cancel
}

func (p *postgresql) Validatorpsql() {
	ctx, cancel := p.contexts()
	defer cancel()
	config, err := pgx.ParseConfig(custompostgreurl())
	if err != nil {
		fmt.Println("Error parsing database URL:", err)
		return
	}

	conn, err := pgx.ConnectConfig(ctx, config)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	p.pgsql = conn
	fmt.Println("connected")
}

func (p *postgresql) PGRowQuery(query string, choice bool) (pgx.Row, pgx.Rows) {
	ctx, cancel := p.contexts()
	defer cancel()
	if !choice {
		rows, err := p.pgsql.Query(ctx, query)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", rows)
			os.Exit(1)
		} else {
			return nil, rows
		}
	}
	row := p.pgsql.QueryRow(ctx, query)
	return row, nil

	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", rows)
	// 	os.Exit(1)
	// }
}

func (p *postgresql) PGExecQuery(query string) (string, error) {
	ctx, cancel := p.contexts()
	defer cancel()
	conn, err := p.pgsql.Exec(ctx, query)
	numAffectedRows := fmt.Sprintf("Number of RowsAffected %v", conn.RowsAffected())
	return numAffectedRows, err
}

func (p *postgresql) StopPsqlConnection() error {
	ctx, cancel := p.contexts()
	defer cancel()
	return p.pgsql.Close(ctx)
}

// func TableQuery(query string) {
// 	rows, err := pgsql.Query(context.Background(), query)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", rows)
// 		os.Exit(1)
// 	}

// 	var emailEntries []models.User
// 	for rows.Next() {
// 		var email models.User
// 		err := rows.Scan(&email.Account_id, &email.Account_create, &email.Email, &email.Username, &email.Password, &email.Account_create)
// 		if err != nil {
// 			fmt.Println("Error scanning row:", err)
// 			continue
// 		}
// 		emailEntries = append(emailEntries, email)
// 	}

// 	if rows.Err() != nil {
// 		fmt.Println("Error iterating rows:", rows.Err())
// 		return
// 	}

// 	fmt.Println(emailEntries)
// 	defer pgsql.Close(context.Background())
// 	defer rows.Close()
// }
