// files.go contains any data access layer representations that access
// data from specific file types, such as XLSX or CSV.

package dal

import (
	"encoding/csv"
	"github.com/n30w/Darkspace/internal/models"
	"github.com/xuri/excelize/v2"
	"os"
)

type ExcelStore struct {
	excelTemplatePath, excelTemplateSheetName, excelTemplateName string
}

func NewExcelStore() *ExcelStore {
	return &ExcelStore{
		excelTemplateSheetName: "Sheet1",
	}
}

// Get retrieves all the data in a file. It takes optional
// arguments. It is a slice, where index 0 is the path and
// index 1 is the sheet name. Defaults to struct initials
// if left blank.
func (es *ExcelStore) Get(path ...string) (
	[][]string, error,
) {
	var p, n string
	p = es.excelTemplatePath
	n = es.excelTemplateSheetName

	if len(path) == 1 {
		p = path[0]
	} else if len(path) > 1 {
		p = path[0]
		n = path[1]
	}

	// Open the file
	f, err := excelize.OpenFile(p)
	if err != nil {
		return nil, err
	}

	f.Close()

	// Get all the rows in a sheet.
	rows, err := f.GetRows(n)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// Save saves the Excel file to a place on disk, given a path.
func (es *ExcelStore) Save(file *excelize.File, to string) (string, error) {
	err := file.SaveAs(to)
	if err != nil {
		return "", nil
	}

	return "", nil
}

// Open opens an Excel file at a specified path. Uses variadric
// parameters to accept an optional value. If the optional value
// is not set, uses the struct default templatePath.
func (es *ExcelStore) Open(path ...string) (*excelize.File, error) {
	var p string
	if len(path) >= 1 {
		p = path[0]
	} else {
		p = es.excelTemplatePath
	}

	f, err := excelize.OpenFile(p)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (es *ExcelStore) AddRow(row []string) error {
	return nil
}

// ========================================================================== //
// CSV defines access operations for accessing data from a CSV file.
// This exists because we currently do not have a functioning database just yet.
// General overview of CSV handling in Go:
// https://earthly.dev/blog/golang-csv-files/
// ========================================================================== //

type CSVStore struct {
	path string
}

func NewCSVStore(p string) *CSVStore {
	return &CSVStore{path: p}
}

// readCSV reads a CSV file at the specified path.
// It returns a multidimensional array of strings.
func (cs *CSVStore) readCSV() ([][]string, error) {
	f, err := os.Open(cs.path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	r := csv.NewReader(f)
	data, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// writeCSV creates a new CSV file.
// This can be used in tandem with readCSV to read a CSV,
// delete a row from the slices,
// then write the new slices to a new CSV file that overwrites the original.
// Here is a helpful article on the writing to a CSV pattern: https
// ://gosamples.dev/write-csv/
func (cs *CSVStore) writeCSV(data [][]string) error {
	f, err := os.Create(cs.path)
	if err != nil {
		return err
	}

	defer f.Close()

	writer := csv.NewWriter(f)

	defer writer.Flush()

	err = writer.WriteAll(data)
	if err != nil {
		return err
	}

	return nil
}

// updateCSV appends a line to a CSV file.
func (cs *CSVStore) updateCSV(row []string) error {

	f, err := os.OpenFile(cs.path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	err = writer.Write(row)
	if err != nil {
		return err
	}

	return nil
}

func (cs *CSVStore) InsertCourse(c *models.Course) error {
	//TODO implement me
	panic("implement me")
}

func (cs *CSVStore) GetCourseByName(name string) (*models.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (cs *CSVStore) GetCourseByID(id string) (*models.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (cs *CSVStore) GetRoster(id string) ([]models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (cs *CSVStore) InsertUser(u *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (cs *CSVStore) GetUserByID(id string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (cs *CSVStore) GetUserByEmail(email string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (cs *CSVStore) GetByUsername(username string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}
