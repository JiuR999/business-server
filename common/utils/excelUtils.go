package utils

import (
	"BusinessServer/common"
	"github.com/xuri/excelize/v2"
	"os"
	"runtime"
	"sync"
)

const (
	Name           = "A"
	AssetCode      = "B"
	AssetHeader    = "C"
	ProductionTime = "D"
	ServiceLength  = "E"
	Price          = "F"
	AssetStatus    = "G"
	AssetType      = "H"
	Producer       = "I"
	Comment        = "J"
	CELL_NOT_VALUE = ""
)

var (
	str2IntMap map[string]int
	once       sync.Once
)

func GetTemplate(fileName string) (*excelize.File, common.SwustError) {
	goos := runtime.GOOS
	var filePath string
	switch goos {
	case "linux":
		filePath = "/etc/business/"
	case "windows":
		filePath, _ = os.Getwd()
		filePath += "/"
	}
	filePath += fileName
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, common.NewServiceError(err.Error())
	}
	return file, nil
}

func SetExcelDropList(file *excelize.File, sheet string, allowBlank bool, sqref string, key []string) {
	validation := excelize.NewDataValidation(allowBlank)
	validation.Sqref = sqref
	validation.SetDropList(key)
	file.AddDataValidation(sheet, validation)
}

func AddAllowBlank(file *excelize.File, sheet string, allowBlank bool, sqref string) {
	validation := excelize.NewDataValidation(allowBlank)
	validation.Sqref = sqref
	file.AddDataValidation(sheet, validation)
}

func GetCellValue(excel *excelize.File, sheet string, cell string) string {
	name, err := excel.GetCellValue(sheet, cell)
	if err != nil {
		return CELL_NOT_VALUE
	}
	return name
}

func GetStr2IntMap() map[string]int {
	once.Do(func() {
		str2IntMap = map[string]int{
			"在用":           0,
			"故障":           1,
			"维修":           2,
			"报废":           3,
			"闲置":           4,
			CELL_NOT_VALUE: 4,
		}
	})
	return str2IntMap
}
