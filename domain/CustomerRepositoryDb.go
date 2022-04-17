package domain

import (
	"database/sql"
	"github.com/fprogress17/banking/errs"
	"github.com/fprogress17/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	//var rows *sql.Rows
	//var err error
	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		rows, err := d.client.Query(findAllSql)
		if err != nil {
			log.Println(" error while querying" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected db error")
		}

		customers := make([]Customer, 0)
		for rows.Next() {
			var c Customer
			err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
			if err != nil {
				logger.Error("error while scanning" + err.Error())
				return nil, errs.NewUnexpectedError("unexpected db error")
			}
			customers = append(customers, c)
		}

		return customers, nil
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		rows, err := d.client.Query(findAllSql, status)
		if err != nil {
			logger.Error(" error while querying" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected db error")
		}

		customers := make([]Customer, 0)
		for rows.Next() {
			var c Customer
			err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
			if err != nil {
				//log.Println("error while scanning" + err.Error())
				logger.Error("error while scanning" + err.Error())
				return nil, errs.NewUnexpectedError("unexpected db error")
			}
			customers = append(customers, c)
		}

		return customers, nil
	}

}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			log.Println("error while scanning customer" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected db error")
		}
	}

	return &c, nil
}
func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:Donsou1234@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}
