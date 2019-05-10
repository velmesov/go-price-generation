package product

import (
	"GoProjects/price/error"
	"GoProjects/price/models"
	"database/sql"
)

// Структура карты цен
type prices struct {
	Price1 float64
	Price2 float64
}

// Создаем карту товаров с типом int
var productsFullData map[int]models.ProductsFullData

// List получение списка товаров
func List(db *sql.DB) map[int]models.ProductsFullData {
	// Получаем основные данные товаров
	rows1, err := db.Query("SELECT `sshop`.`skod`  AS `code`, SUM(`sshop`.`tost`)  AS `in_stock`, `cshop`.`ctov`  AS `name` FROM `sshop` INNER JOIN `cshop` ON `sshop`.`idcat` = `cshop`.`idn` WHERE `tost` > '0' GROUP BY `sshop`.`skod` ORDER BY `cshop`.`ctov`")
	error.CheckNil(err)
	defer rows1.Close()

	// Если нет ошибок
	if err = rows1.Err(); err == nil {
		// Создаем срез для основных данных товаров
		products := make([]*models.ProductsMainData, 0)

		// Пробегаемся по строкам
		for rows1.Next() {
			pointer := new(models.ProductsMainData)
			err := rows1.Scan(&pointer.Code, &pointer.InStock, &pointer.Name)
			error.CheckNil(err)

			products = append(products, pointer)
		}

		// Получаем остальные данные товаров
		// В данном случае сортируем партии по максимальному остатку
		rows2, err := db.Query("SELECT `rcen` AS `price_1`, `rcen1` AS `price_2`, `skod`  AS `code` FROM `sshop` WHERE `tost` > '0' ORDER BY `skod`, `tost` DESC, `idn` ASC")
		error.CheckNil(err)
		defer rows2.Close()

		// Если нет ошибок
		if err = rows2.Err(); err == nil {
			// Создаем карту цен с типом int
			var dataPrices map[int]prices

			// Инициализируем карту цен
			dataPrices = make(map[int]prices)

			// Предыдущий код
			prevCode := 0

			// Пробегаемся по строкам
			for rows2.Next() {
				pointer := new(models.ProductsOtherData)
				err := rows2.Scan(&pointer.Price1, &pointer.Price2, &pointer.Code)
				error.CheckNil(err)

				// Выбираем по первой в результате партии по каждому коду
				if pointer.Code != prevCode {
					dataPrices[pointer.Code] = prices{
						Price1: pointer.Price1,
						Price2: pointer.Price2,
					}

					// Временно сохраняем код
					prevCode = pointer.Code
				}
			}

			// Инициализируем карту товаров
			productsFullData = make(map[int]models.ProductsFullData)

			// Создаем список товаров со всеми данными
			for i, v := range products {
				productsFullData[i] = models.ProductsFullData{
					Code:    v.Code,
					InStock: v.InStock,
					Name:    v.Name,
					Price1:  dataPrices[v.Code].Price1,
					Price2:  dataPrices[v.Code].Price2,
				}
			}
		} else {
			error.CheckNil(err)
		}
	} else {
		error.CheckNil(err)
	}

	// Возвращаем список товаров
	return productsFullData
}
