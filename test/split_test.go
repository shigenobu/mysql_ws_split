package test

import (
	"encoding/json"
	"fmt"
	"github.com/shigenobu/mysql_ws_split/func"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimple(t *testing.T) {
	input := "aa,bb,cc"

	output := _func.Split(",", input)
	fmt.Println(output)

	var result []string
	json.Unmarshal([]byte(output), &result)
	assert.Equal(t, "aa", result[0])
	assert.Equal(t, "bb", result[1])
	assert.Equal(t, "cc", result[2])
	assert.Equal(t, len(result), 3)
}

func TestMultiLang(t *testing.T) {
	input := "田中。佐藤。清水。"

	output := _func.Split("。", input)
	fmt.Println(output)

	var result []string
	json.Unmarshal([]byte(output), &result)
	assert.Equal(t, "田中", result[0])
	assert.Equal(t, "佐藤", result[1])
	assert.Equal(t, "清水", result[2])
	assert.Equal(t, "", result[3])
	assert.Equal(t, len(result), 4)
}

func TestPattern(t *testing.T) {
	input := "hypertext language, programming"

	output := _func.Split("[\\s,]+", input)
	fmt.Println(output)

	var result []string
	json.Unmarshal([]byte(output), &result)
	assert.Equal(t, "hypertext", result[0])
	assert.Equal(t, "language", result[1])
	assert.Equal(t, "programming", result[2])
	assert.Equal(t, len(result), 3)
}

func TestPatternMultiLang(t *testing.T) {
	input := "田中　佐藤。清水、 鈴木"

	output := _func.Split("([　 。]|、 )", input)
	fmt.Println(output)

	var result []string
	json.Unmarshal([]byte(output), &result)
	assert.Equal(t, "田中", result[0])
	assert.Equal(t, "佐藤", result[1])
	assert.Equal(t, "清水", result[2])
	assert.Equal(t, "鈴木", result[3])
	assert.Equal(t, len(result), 4)
}

func TestNull(t *testing.T) {
	var input string
	var del string

	output := _func.Split(del, input)
	fmt.Println(output)

	var result []string
	json.Unmarshal([]byte(output), &result)
	assert.Equal(t, len(result), 0)
}

func TestEmpty(t *testing.T) {
	var input string = ""
	var del string = ""

	output := _func.Split(del, input)
	fmt.Println(output)

	var result []string
	json.Unmarshal([]byte(output), &result)
	assert.Equal(t, len(result), 0)
}