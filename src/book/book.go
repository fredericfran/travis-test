package book

import (
	"fmt"
)

var book Book

const (
	//NoNameError is the error if there is no name in the record
	NoNameError = "no name in record"
	//NoEmailError is the error if there is no email in the record
	NoEmailError = "no email in the record"
	//NoRecordFound is the error if there is no record with specified name
	NoRecordFound = "no record found"
)

//Book is the type which implements the address book
type Book struct {
	records map[string]Record
}

func New() Book {
	return Book{records: make(map[string]Record)}
}

//Record is the type which implements an entry in the address book
type Record struct {
	Name string
	Email string
}

func (book *Book) Add(record Record) error {
	if err:=validateRecord(record); err != nil {
		return err
	}
	book.records[record.Name]= record
	return nil
}

func (book *Book) Record (name string) (Record, error) {
	record, found := book.records[name]
	if !found {
		return Record{},fmt.Errorf("%s: name %s submitted", NoRecordFound, name)
	} 
	
	return record, nil
}

func (book *Book) Delete (name string) error {
	if _, found := book.records[name]; !found {
		return fmt.Errorf("%s: name %s submitted", NoRecordFound, name)
	} 
	
	delete(book.records, name)
	
	return nil
}


func validateRecord(record Record) error {
	errString:=""
	if record.Name == "" {
		errString += fmt.Sprintf(" :%v", NoNameError)
	}
	
	if record.Email == "" {
		errString += fmt.Sprintf(" :%v", NoEmailError)
	}
	
	if errString != "" {
		errString = "errors when validating record" + errString
		return fmt.Errorf(errString)
	}
	
	return nil
}
