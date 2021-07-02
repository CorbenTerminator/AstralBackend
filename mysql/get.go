package mysql

import (
	"bytes"
	"fmt"
	"log"
)

const GetOrdersSQL = `
select json_object(
	'order_id', t2.order_id,
	'status_id', t2.status_id,
	'status', t3.name,
	'createdAt', UNIX_TIMESTAMP(t2.createdAt))
from orders as t2 
inner join status as t3 on t2.status_id=t3.status_id where t2.user_id=? order by t2.createdAt desc
`

func (w *Worker) GetOrders(params []interface{}) (string, error) {
	rows, err := w.DB.Query(GetOrdersSQL, params...)
	if err != nil {
		log.Printf("GetOrders: %+v", err)
		return "", err
	}
	data := ""
	ifFirst := true

	var buf bytes.Buffer
	fmt.Fprint(&buf, "[")
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&data)
		if err != nil {
			log.Printf("GetOrders: %+v", err)
			return "", err
		}
		if ifFirst {
			ifFirst = false
			fmt.Fprint(&buf, data)
		} else {
			fmt.Fprintf(&buf, ",%s", data)
		}
	}
	fmt.Fprint(&buf, "]")
	return buf.String(), nil
}

const GetProductsSQL = `
select json_object(
	'product_id', t1.product_id,
	'product_name', t1.name,
	'price', t1.price,
	'category_id', t2.category_id,
	'category', t2.name
) from products t1
inner join categories t2 on t1.category_id=t2.category_id order by t1.product_id
`

func (w *Worker) GetProducts(params []interface{}) (string, error) {
	rows, err := w.DB.Query(GetProductsSQL, params...)
	if err != nil {
		log.Printf("GetProducts: %+v", err)
		return "", err
	}
	data := ""
	ifFirst := true

	var buf bytes.Buffer
	fmt.Fprint(&buf, "[")
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&data)
		if err != nil {
			log.Printf("GetProducts: %+v", err)
			return "", err
		}
		if ifFirst {
			ifFirst = false
			fmt.Fprint(&buf, data)
		} else {
			fmt.Fprintf(&buf, ",%s", data)
		}
	}
	fmt.Fprint(&buf, "]")
	return buf.String(), nil
}

const GetOrderProductsSQL = `
select json_object(
	'product_name', t1.name,
	'price', t1.price,
	'category', t2.name
) from products t1
inner join categories t2 on t1.category_id=t2.category_id
inner join order_products t3 on t3.product_id=t1.product_id
inner join orders t4 on t4.order_id=t3.order_id where t3.order_id=? and t4.user_id=? order by t1.product_id
`

func (w *Worker) GetOrderProducts(params []interface{}) (string, error) {
	rows, err := w.DB.Query(GetOrderProductsSQL, params...)
	if err != nil {
		log.Printf("GetOrderProducts: %+v", err)
		return "", err
	}
	data := ""
	ifFirst := true

	var buf bytes.Buffer
	fmt.Fprint(&buf, "[")
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&data)
		if err != nil {
			log.Printf("GetOrderProducts: %+v", err)
			return "", err
		}
		if ifFirst {
			ifFirst = false
			fmt.Fprint(&buf, data)
		} else {
			fmt.Fprintf(&buf, ",%s", data)
		}
	}
	fmt.Fprint(&buf, "]")
	return buf.String(), nil
}

// func (w *Worker) GetOrderByID(params []interface{}) (string, error) {
// 	sqlStr := fmt.Sprintf("%s WHERE order_id=?", GetOrdersSQL)
// 	rows, err := w.DB.Query(sqlStr, params...)
// 	if err != nil {
// 		return "", err
// 	}
// 	data := ""
// 	defer rows.Close()
// 	if rows.Next() {
// 		err = rows.Scan(&data)
// 		if err != nil {
// 			return "", err
// 		}
// 		return data, nil
// 	}
// 	return "", nil
// }

const GetUserByLogin = `
	select user_id, password from users where login = ?
`

func (w *Worker) GetUserByLogin(params []interface{}) (string, string, error) {
	rows, err := w.DB.Query(GetUserByLogin, params...)
	if err != nil {
		log.Printf("GetUserByLogin: %+v", err)
		return "", "", err
	}
	user_id := ""
	password := ""
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&user_id, &password)
		if err != nil {
			log.Printf("GetUserByLogin: %+v", err)
			return "", "", err
		}
		return user_id, password, nil
	}
	return "", "", nil
}

const GetCategoriesSQL = `
	select json_object(
		'category_id', category_id,
		'name', name
	) from categories
`

func (w *Worker) GetCategories(params []interface{}) (string, error) {
	rows, err := w.DB.Query(GetCategoriesSQL, params...)
	if err != nil {
		log.Printf("GetCategories: %+v", err)
		return "", err
	}
	data := ""
	ifFirst := true

	var buf bytes.Buffer
	fmt.Fprint(&buf, "[")
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&data)
		if err != nil {
			log.Printf("GetCategories: %+v", err)
			return "", err
		}
		if ifFirst {
			ifFirst = false
			fmt.Fprint(&buf, data)
		} else {
			fmt.Fprintf(&buf, ",%s", data)
		}
	}
	fmt.Fprint(&buf, "]")
	return buf.String(), nil
}
