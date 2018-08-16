package main

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPlus(t *testing.T) {
	a := 2
	b := 3
	c := Plus(a, b)
	if c != 5 {
		panic("wrong c result")
	}
}

func TestPlusNumbersAndWriteResultToTable(t *testing.T) {
	a := 9
	b := 7

	db := ConnectToDatabase()
	defer cleanTableAfterTest(db)

	c := Plus(a, b)
	err := WriteResultToTable(db, c)
	require.NoError(t, err)

	result, err := GetResultFromTable(db)
	require.NoError(t, err)
	assert.Equal(t, 16, result)
}

func cleanTableAfterTest(db *gorm.DB) {
	db.Delete(&Number{})
}
