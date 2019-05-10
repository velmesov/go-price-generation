package main

import (
	"GoProjects/price/models/product"
	"GoProjects/price/xlsx"
	"flag"

	_ "github.com/go-sql-driver/mysql"
)

var (
	pathSaveFile string
	pathTemplate string
)

func init() {
	// Устанавливаем флаги
	flag.StringVar(&pathSaveFile, "s", "Document.xlsx", "Укажите путь к сохраняемому файлу включая название")
	flag.StringVar(&pathTemplate, "t", "Template.xlsx", "Укажите путь к шаблону включая название файла")
}

func main() {
	// Парсим флаги
	flag.Parse()

	// Подключаемся к базе
	db := Db()

	// Закрываем соединение по завершению текущей функции
	defer db.Close()

	// Получаем список товаров
	products := product.List(db)

	// Пишем в шаблон список товаров
	xlsx.WriteToTemplate(products, pathSaveFile, pathTemplate)
}
