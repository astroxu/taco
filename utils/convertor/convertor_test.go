/**
 * @Author: Sean
 * @Date: 2022/5/20 14:14
 */

package convertor

import (
	"fmt"
	"github.com/seanxu24/taco/utils/internal"
	"testing"
)

func TestToChar(t *testing.T) {
	assert := internal.NewAssert(t, "TestToChar")

	tests := []struct {
		cases    string
		expected []string
	}{
		{"", []string{""}},
		{"abc", []string{"a", "b", "c"}},
		{"1 2#3", []string{"1", " ", "2", "#", "3"}},
	}
	for i := 0; i < len(tests); i++ {
		assert.Equal(tests[i].expected, ToChar(tests[i].cases))
	}
}

func TestToBool(t *testing.T) {
	assert := internal.NewAssert(t, "TestToBool")

	tests := []struct {
		cases    string
		expected bool
	}{
		{"1", true},
		{"true", true},
		{"True", true},
		{"false", false},
		{"False", false},
		{"0", false},
		{"123", false},
		{"0.0", false},
		{"abc", false},
	}

	for i := 0; i < len(tests); i++ {
		actual, _ := ToBool(tests[i].cases)
		assert.Equal(tests[i].expected, actual)
	}
}

func TestToBytes(t *testing.T) {
	assert := internal.NewAssert(t, "TestToBytes")

	tests := []struct {
		cases    any
		expected []byte
	}{
		{0, []byte{0, 0, 0, 0, 0, 0, 0, 0}},
		{false, []byte{102, 97, 108, 115, 101}},
		{"1", []byte{49}},
	}
	for i := 0; i < len(tests); i++ {
		actual, _ := ToBytes(tests[i].cases)
		assert.Equal(tests[i].expected, actual)
	}

	bytesData, err := ToBytes("abc")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	assert.Equal("abc", ToString(bytesData))
}

func TestToInt(t *testing.T) {
	assert := internal.NewAssert(t, "TestToInt")

	tests := []struct {
		cases    any
		expected int64
	}{
		{"123", 123},
		{"-123", -123},
		{123, 123},
		{uint(123), 123},
		{uint8(123), 123},
		{uint16(123), 123},
		{uint32(123), 123},
		{uint64(123), 123},
		{float32(12.3), 12},
		{float64(12.3), 12},
		{"abc", 0},
		{false, 0},
		{"11111111111111111111", 0},
	}

	for i := 0; i < len(tests); i++ {
		actual, _ := ToInt(tests[i].cases)
		assert.Equal(tests[i].expected, actual)
	}
}

func TestToFloat(t *testing.T) {
	assert := internal.NewAssert(t, "TestToFloat")

	tests := []struct {
		cases    any
		expected float64
	}{
		{"", 0},
		{"-1", -1},
		{"-.11", -0.11},
		{"1.23e3", 1230},
		{".123e10", 0.123e10},
		{"abc", 0},
		{int(0), 0},
		{int8(1), 1},
		{int16(-1), -1},
		{int32(123), 123},
		{int64(123), 123},
		{uint(123), 123},
		{uint8(123), 123},
		{uint16(123), 123},
		{uint32(123), 123},
		{uint64(123), 123},
		{float32(12.3), 12.300000190734863},
		{float64(12.3), 12.3},
	}

	for i := 0; i < len(tests); i++ {
		actual, _ := ToFloat(tests[i].cases)
		assert.Equal(tests[i].expected, actual)
	}
}

func TestToString(t *testing.T) {
	assert := internal.NewAssert(t, "TestToString")

	aMap := make(map[string]int)
	aMap["a"] = 1
	aMap["b"] = 2
	aMap["c"] = 3

	type TestStruct struct {
		Name string
	}
	aStruct := TestStruct{Name: "TestStruct"}

	tests := []struct {
		cases    any
		expected string
	}{
		{"", ""},
		{nil, ""},
		{int(0), "0"},
		{int8(1), "1"},
		{int16(-1), "-1"},
		{int32(123), "123"},
		{int64(123), "123"},
		{uint(123), "123"},
		{uint8(123), "123"},
		{uint16(123), "123"},
		{uint32(123), "123"},
		{uint64(123), "123"},
		{float32(12.3), "12.300000190734863"},
		{float64(12.3), "12.3"},
		{true, "true"},
		{false, "false"},
		{[]int{1, 2, 3}, "[1,2,3]"},
		{aMap, "{\"a\":1,\"b\":2,\"c\":3}"},
		{aStruct, "{\"Name\":\"TestStruct\"}"},
		{[]byte{104, 101, 108, 108, 111}, "hello"},
	}

	for i := 0; i < len(tests); i++ {
		actual := ToString(tests[i].cases)
		assert.Equal(tests[i].expected, actual)
	}
}

func TestToJson(t *testing.T) {
	assert := internal.NewAssert(t, "TestToJson")

	var aMap = map[string]int{"a": 1, "b": 2, "c": 3}
	mapJsonStr, _ := ToJson(aMap)
	assert.Equal("{\"a\":1,\"b\":2,\"c\":3}", mapJsonStr)

	type TestStruct struct {
		Name string
	}
	aStruct := TestStruct{Name: "TestStruct"}
	structJsonStr, _ := ToJson(aStruct)
	assert.Equal("{\"Name\":\"TestStruct\"}", structJsonStr)
}

func TestStructToMap(t *testing.T) {
	assert := internal.NewAssert(t, "TestStructToMap")

	type People struct {
		Name string `json:"name"`
		age  int
	}
	p := People{
		Name: "test",
		age:  100,
	}
	pm, _ := StructToMap(p)
	var expected = map[string]any{"name": "test"}
	assert.Equal(expected, pm)
}

func TestColorHexToRGB(t *testing.T) {
	assert := internal.NewAssert(t, "TestColorHexToRGB")

	colorHex := "#003366"
	r, g, b := ColorHexToRGB(colorHex)
	colorRGB := fmt.Sprintf("%d,%d,%d", r, g, b)
	expected := "0,51,102"

	assert.Equal(expected, colorRGB)
}

func TestColorRGBToHex(t *testing.T) {
	assert := internal.NewAssert(t, "TestColorRGBToHex")

	r, g, b := 0, 51, 102
	colorHex := ColorRGBToHex(r, g, b)
	expected := "#003366"

	assert.Equal(expected, colorHex)
}
