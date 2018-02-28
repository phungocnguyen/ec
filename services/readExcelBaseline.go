package services

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"path/filepath"
	"platformOps-EC/models"
)

func LoadFromExcel(file string) (b models.Baseline, controls []models.Control) {
	return loadBaseline(file), loadControl(file)
}

func loadBaseline(file string) (b models.Baseline) {
	name := filepath.Base(file)
	return models.Baseline{Name: name}
}

func loadControl(file string) (controls []models.Control) {
	xlFile, err := xlsx.OpenFile(file)
	if err != nil {
		fmt.Println("error reading")
	}
	sheet := xlFile.Sheets[0]
	length := len(sheet.Rows)

	// Removing header in excel sheet
	rows := sheet.Rows[1 : length-1]

	for _, row := range rows {

		cells := row.Cells

		reqId, err := cells[0].Int()
		if err != nil {
			fmt.Println("error reading reqId")
		}

		control := models.Control{ReqId: reqId, CisId: cells[1].String(), Category: cells[2].String(),
			Requirement: cells[3].String(), Discussion: cells[4].String(),
			CheckText: cells[5].String(), FixText: cells[6].String(),
			RowDesc: cells[0].String()}

		controls = append(controls, control)
	}

	return controls
}