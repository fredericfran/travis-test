package book_test

import (
	"testing"
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


