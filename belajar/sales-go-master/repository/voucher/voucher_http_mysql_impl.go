package voucher

import (
	"fmt"
	"context"
	"sales-go/db"
	"sales-go/model"
	"time"
)

type repositoryhttpmysql struct {}

func NewMySQLHTTPRepository () *repositoryhttpmysql {
	return &repositoryhttpmysql{}
}

func (repo *repositoryhttpmysql) GetList() (listVoucher []model.Voucher, err error) {
	db := client.NewConnection(client.Database).GetMysqlConnection()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `SELECT id, code, persen FROM voucher`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.QueryContext(ctx)
	if err != nil {
		return
	}

	for res.Next() {
		var voucher model.Voucher
		res.Scan(&voucher.Id, &voucher.Code, &voucher.Persen)
		fmt.Println(voucher)
		listVoucher = append(listVoucher, voucher)
	}

	return
}

func (repo *repositoryhttpmysql) GetVoucherByCode(code string) (voucherData model.Voucher, err error) {
	db := client.NewConnection(client.Database).GetMysqlConnection()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `SELECT id, code, persen FROM voucher WHERE code = ?`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.QueryContext(ctx, code)
	if err != nil {
		return
	}

	for res.Next() {
		res.Scan(&voucherData.Id, &voucherData.Code, &voucherData.Persen)
	}
	return
}

func (repo *repositoryhttpmysql) Create(req []model.VoucherRequest) (response []model.Voucher, err error) {
	db := client.NewConnection(client.Database).GetMysqlConnection()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	trx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	query := `INSERT INTO voucher (code, persen) values (?, ?)`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	for _, v := range req {
		res, err := stmt.ExecContext(ctx, v.Code, v.Persen)
		if err != nil {
			trx.Rollback()
			return []model.Voucher{}, err
		}

		lastID, err := res.LastInsertId()
		if err != nil {
			return []model.Voucher{}, err
		}

		response = append(response, model.Voucher{
			Id:     int(lastID),
			Code:   v.Code,
			Persen: v.Persen,
		})
	}

	trx.Commit()

	return
}
