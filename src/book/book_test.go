package book_test

import (
	"testing"
	"strings"
	"github.com/fredericfran/travis-test/src/book"
	"github.com/stretchr/testify/assert"
)

func TestAddRecordSuccess(t *testing.T) {
	//GIVEN
	record:=book.Record{Name: "adam", Email: "adam@test.com"}
	addressBook := book.New()
	
	//WHEN
	err:=addressBook.Add(record)
	
	//THEN
	assert.NoError(t, err)
}

func TestAddRecordValidateNameError(t *testing.T) {
	//GIVEN
	record:=book.Record{Email: "adam@test.com"}
	addressBook := book.New()
	
	//WHEN
	err:=addressBook.Add(record)
	
	//THEN
	assert.Error(t, err)
	assert.Equal(t, true, strings.Contains(err.Error(), book.NoNameError))
}

func TestAddRecordValidateEmailError(t *testing.T) {
	//GIVEN
	record:=book.Record{Name: "adam"}
	addressBook := book.New()

	//WHEN
	err:=addressBook.Add(record)

	//THEN
	assert.Error(t, err)
	assert.Equal(t, true, strings.Contains(err.Error(), book.NoEmailError))
}

func TestAddRecordVerifyRetrieve(t *testing.T) {
	//GIVEN
	record:=book.Record{Name: "adam", Email: "adam@test.com"}
	addressBook := book.New()
	err:=addressBook.Add(record)
	assert.NoError(t, err, "expected no error because setting object for testing")

	//WHEN
	actualRecord, err:=addressBook.Record("adam")

	//THEN
	assert.NoError(t, err)
	assert.Equal(t, record, actualRecord)
}

func TestAddRecordRetrieveError(t *testing.T) {
	//GIVEN
	addressBook := book.New()

	//WHEN
	_, err:=addressBook.Record("adam")

	//THEN
	assert.Error(t, err)
	assert.Equal(t, true, strings.Contains(err.Error(), book.NoRecordFound))
}

func TestDeleteRecordSuccess(t *testing.T) {
	//GIVEN
	record:=book.Record{Name: "adam", Email: "adam@test.com"}
	addressBook := book.New()
	err:=addressBook.Add(record)
	assert.NoError(t, err, "expected no error because setting object for testing")

	//WHEN
	err=addressBook.Delete("adam")

	//THEN
	assert.NoError(t, err)
	_,err=addressBook.Record("adam")
	assert.Error(t,err)
	assert.Equal(t, true, strings.Contains(err.Error(), book.NoRecordFound))
}

func TestDeleteRecordError(t *testing.T) {
	//GIVEN
	addressBook := book.New()

	//WHEN
	err:=addressBook.Delete("adam")

	//THEN
	assert.Error(t, err)
	assert.Equal(t, true, strings.Contains(err.Error(), book.NoRecordFound))
}