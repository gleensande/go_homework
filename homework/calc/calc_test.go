package main

import (
	"bytes"
	"testing"
)

const testSimpleString = `1 2 3 4 + * + =`
const testSimpleResult = `15`

func TestCalcSimple(t *testing.T) {
	out := new(bytes.Buffer)
	err := calc(testSimpleString, out)
	if err != nil {
		t.Errorf("test for OK Failed - error")
	}
	result := out.String()
	if result != testSimpleResult {
		t.Errorf("test for OK Failed - results not match\nGot:\n%v\nExpected:\n%v", result, testSimpleResult)
	}
}

const testSimpleString1 = `1 2 + 3 4 + * =`
const testSimpleResult1 = `21`

func TestCalcSimple1(t *testing.T) {
	out := new(bytes.Buffer)
	err := calc(testSimpleString1, out)
	if err != nil {
		t.Errorf("test for OK Failed - error")
	}
	result := out.String()
	if result != testSimpleResult1 {
		t.Errorf("test for OK Failed - results not match\nGot:\n%v\nExpected:\n%v", result, testSimpleResult1)
	}
}

const testAbnormalString = `1 2 3 * * * * = = = =`
const testAbnormalResult = ``
const testAbnormalError = `abnormal syntax`

func TestCalcAbnormanl(t *testing.T) {
	out := new(bytes.Buffer)
	err := calc(testAbnormalString, out)
	if err != nil {
		if err.Error() != testAbnormalError {
			t.Errorf("test for OK Failed - errors not match\nGot:\n%v\nExpected:\n%v", err.Error(), testAbnormalError)
		}
	}
	result := out.String()
	if result != testAbnormalResult {
		t.Errorf("test for OK Failed - results not match\nGot:\n%v\nExpected:\n%v", result, testAbnormalResult)
	}
}

const testDivisionString = `8 2 / 2 /`
const testDivisionResult = `2`

func TestCalcDivision(t *testing.T) {
	out := new(bytes.Buffer)
	err := calc(testDivisionString, out)
	if err != nil {
		t.Errorf("test for OK Failed - error")
	}
	result := out.String()
	if result != testDivisionResult {
		t.Errorf("test for OK Failed - results not match\nGot:\n%v\nExpected:\n%v", result, testDivisionResult)
	}
}

const testMultiplexString = `5 5 * 4 *`
const testMultiplexResult = `100`

func TestCalcMultiplex(t *testing.T) {
	out := new(bytes.Buffer)
	err := calc(testMultiplexString, out)
	if err != nil {
		t.Errorf("test for OK Failed - error")
	}
	result := out.String()
	if result != testMultiplexResult {
		t.Errorf("test for OK Failed - results not match\nGot:\n%v\nExpected:\n%v", result, testMultiplexResult)
	}
}
