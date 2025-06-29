package handler

import (
	"errors"
	"log"
	"net"
	"query-service/internal/errs"

	"github.com/go-sql-driver/mysql"
)

func DBErrHandler(err error) error {
	var (
		opErr     *net.OpError
		driverErr *mysql.MySQLError
	)
	if errors.As(err, &opErr) {
		log.Println(err.Error())
		return errs.NewInternalError(opErr.Error())
	}
	if errors.As(err, &driverErr) {
		log.Printf("Code:%d Message:%s \n", driverErr.Number, driverErr.Message)
		return errs.NewInternalError(driverErr.Message)
	}
	log.Println(err.Error())
	return errs.NewInternalError(err.Error())
}
