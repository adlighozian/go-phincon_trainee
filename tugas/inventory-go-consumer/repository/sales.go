package repository

import (
	"inventory/db"
	"inventory/model"
)

func TambahSales(req []model.SendSales) {
	db := db.GetConnection()

	// Start a transaction
	tx := db.Begin()

	// Execute the native SQL insert statement within the transaction
	updateProduct := `update products set stock = $1 where name = $2`
	selectProduct := `select stock from products where name = $1`
	insertSales := `insert into sales (order_number,orang,total) values ($1, $2, $3) returning id`
	insertDetail := `insert into sales_detail (sales_id, item,price,quantity,total) values ($1,$2,$3,$4,$5) returning id`
	checkProduct := `select id from products where name = $1`

	for _, d := range req {
		var idSales uint
		var idSalesDetail uint
		var err error
		var idProduct uint
		var stock int

		err = tx.Raw(checkProduct, d.Item).Scan(&idProduct).Error
		if err != nil {
			tx.Rollback()
		}

		if idProduct == 0 {
			continue
		}

		// get stock
		err = tx.Raw(selectProduct, d.Item).Scan(&stock).Error
		if err != nil {
			tx.Rollback()
		}

		if stock < d.Total {
			continue
		}

		jumlahStock := stock - d.Total

		// update product
		err = db.Exec(updateProduct, jumlahStock, d.Item).Error
		if err != nil {
			tx.Rollback()
		}

		//insert sales
		err = tx.Raw(insertSales, d.OrderNumber, d.From, d.Total).Scan(&idSales).Error
		if err != nil {
			tx.Rollback()
		}

		// insert Sales detail
		err = tx.Raw(insertDetail, idSales, d.Item, d.Price, d.Total, jumlahStock).Scan(&idSalesDetail).Error
		if err != nil {
			tx.Rollback()
		}

	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
	}

}
