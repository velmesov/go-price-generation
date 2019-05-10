package xlsx

import (
	"GoProjects/price/error"
	"GoProjects/price/models"
	"fmt"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// WriteToTemplate создание файла
func WriteToTemplate(products map[int]models.ProductsFullData, pathSaveFile string, pathTemplate string) {
	xlsx, err := excelize.OpenFile(pathTemplate)
	error.CheckNil(err)

	// Стили выравнивания
	styleAlignment, err := xlsx.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"}}`)
	error.CheckNil(err)

	// Стили переноса
	styleWrap, err := xlsx.NewStyle(`{"alignment":{"wrap_text":true,"vertical":"center"}}`)
	error.CheckNil(err)

	// Записываем дату и время создания документа
	currTime := time.Now()

	xlsx.SetCellValue("Лист1", "E3", currTime.Format("02.01.2006"))
	xlsx.SetCellValue("Лист1", "F3", currTime.Format("15:04"))

	// Номер строки с которой начать вывод товаров
	indexStartRow := 6

	// Формируем список товаров
	// TODO: Не использовать итератор range по карте, теряется сортировка
	// Данные в карте хранятся неотсортированными по ключам
	for i := 0; i < len(products); i++ {
		positionCode := fmt.Sprintf("%s%d", "A", indexStartRow)
		positionName := fmt.Sprintf("%s%d", "B", indexStartRow)
		positionInStock := fmt.Sprintf("%s%d", "C", indexStartRow)
		positionPrice1 := fmt.Sprintf("%s%d", "D", indexStartRow)
		positionPrice2 := fmt.Sprintf("%s%d", "E", indexStartRow)

		xlsx.SetCellValue("Лист1", positionCode, products[i].Code)
		xlsx.SetCellValue("Лист1", positionName, products[i].Name)
		xlsx.SetCellValue("Лист1", positionInStock, products[i].InStock)
		xlsx.SetCellValue("Лист1", positionPrice1, products[i].Price1)
		xlsx.SetCellValue("Лист1", positionPrice2, products[i].Price2)

		xlsx.SetCellStyle("Лист1", positionCode, positionCode, styleAlignment)
		xlsx.SetCellStyle("Лист1", positionName, positionName, styleWrap)
		xlsx.SetCellStyle("Лист1", positionInStock, positionInStock, styleAlignment)
		xlsx.SetCellStyle("Лист1", positionPrice1, positionPrice1, styleAlignment)
		xlsx.SetCellStyle("Лист1", positionPrice2, positionPrice2, styleAlignment)

		xlsx.SetRowHeight("Лист1", indexStartRow, recalcHeight(products[i].Name))

		indexStartRow++
	}

	// Сохраняем файл
	err = xlsx.SaveAs(pathSaveFile)
	error.CheckNil(err)
}

/*
	Пересчет высоты строк в зависимости
	от количества символов в строке
*/
func recalcHeight(name string) float64 {
	var defaultHeight = 14
	var maxSimbols = 65
	var height = defaultHeight
	var countSimbols = len([]rune(name))

	if countSimbols > maxSimbols {
		height = ((countSimbols / maxSimbols) + 1) * defaultHeight
	}

	return float64(height)
}
