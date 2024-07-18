package service

import (
	"ExcelToStruct/model"
	"ExcelToStruct/utils"
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"mime/multipart"
	"reflect"
	"strings"
)

func ConvertExcelToStruct(file multipart.File) ([]model.Item, error) {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, errors.New("error_while_opening_excel_file")
	}

	var sheetData [][]string
	if strings.EqualFold(strings.TrimSpace(f.GetSheetName(1)), model.ItemSheet) {
		sheetData = f.GetRows(f.GetSheetName(1))
	} else {
		return nil, errors.New(fmt.Sprintf("sheet with name %s not found", model.ItemSheet))
	}

	dataRowOffset := 2 // this represents that data in Excel is starting from row 2
	itemList, err := getStructFromExcelRows(dataRowOffset, sheetData)
	if err != nil {
		return nil, err
	}
	return itemList, nil
}

func getStructFromExcelRows(dataOffset int, sheetData [][]string) (itemList []model.Item, err error) {
	headers := sheetData[0]
	columnNameByColumnNo, err := utils.BuildHeaderMap(headers)
	if err != nil {
		return nil, err
	}

	FieldNameByExcelTag := utils.BuildFieldNameByTagMap(&model.Item{})

	for i, row := range sheetData[1:] {
		if utils.IsRowEmpty(row) {
			continue
		}
		var item model.Item
		element := reflect.ValueOf(&item).Elem()

		for j, columnValue := range row {
			if fieldName, ok := FieldNameByExcelTag[columnNameByColumnNo[j]]; ok {
				field := element.FieldByName(fieldName)
				err := utils.SetValueToField(field, columnValue)
				if err != nil {
					return nil, errors.New(fmt.Sprintf("error setting %s = %s, rowNo: %d columnNo:%d\n", fieldName, columnValue, i+dataOffset, j))
				}
			}
		}
		itemList = append(itemList, item)
	}
	return itemList, nil
}
