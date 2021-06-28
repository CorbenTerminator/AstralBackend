package mysql

import (
	"database/sql"
	// "encoding/json"
)

const PostOrderSQL = `
insert into orders(user_id, createdAt, status_id) values
(?, NOW(), ?)
`

func (w *Worker) PostOrder(json []byte, userid string) (int64, error) {
	val := &order{}
	if err := val.UnmarshalJSON(json); err != nil {
		return 0, err
	}
	tx, err := w.DB.Begin()
	if err != nil {
		return 0, err
	}
	var res sql.Result
	res, err = tx.Exec(PostOrderSQL,
		userid,
		val.StatusID,
	)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

const PostProductSQL = `
insert into products(name, price, category_id) values 
(?, ?, ?)
`

func (w *Worker) PostProduct(json []byte) (int64, error) {
	val := &product{}
	if err := val.UnmarshalJSON(json); err != nil {
		return 0, err
	}
	tx, err := w.DB.Begin()
	if err != nil {
		return 0, err
	}
	var res sql.Result
	res, err = tx.Exec(PostProductSQL,
		val.Name,
		val.Price,
		val.CategoryID,
	)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

const PostOrderProductsSQL = `
	insert into order_products(order_id, product_id) values
	(?, ?)
`

func (w *Worker) PostOrderProducts(json []byte, orderid string) error {
	val := &orderProducts{}
	if err := val.UnmarshalJSON(json); err != nil {
		return err
	}
	tx, err := w.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(PostOrderProductsSQL,
		orderid,
		val.ProductID,
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil

}
