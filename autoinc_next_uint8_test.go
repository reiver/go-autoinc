package autoinc_test

import (
	"testing"

	"github.com/reiver/go-autoinc"
)

func TestAutoInc_Next_uint8(t *testing.T) {

	tests := []struct{
		TimesBefore uint
		Expected uint8
	}{
		{
			Expected: 1,
		},



		{
			TimesBefore: 0,
			Expected:    1,
		},
		{
			TimesBefore: 1,
			Expected:    2,
		},
		{
			TimesBefore: 2,
			Expected:    3,
		},
		{
			TimesBefore: 3,
			Expected:    4,
		},
		{
			TimesBefore: 4,
			Expected:    5,
		},
		{
			TimesBefore: 5,
			Expected:    6,
		},
		{
			TimesBefore: 6,
			Expected:    7,
		},
		{
			TimesBefore: 7,
			Expected:    8,
		},
		{
			TimesBefore: 8,
			Expected:    9,
		},
		{
			TimesBefore: 9,
			Expected:    10,
		},
		{
			TimesBefore: 10,
			Expected:    11,
		},
		{
			TimesBefore: 11,
			Expected:    12,
		},
		{
			TimesBefore: 12,
			Expected:    13,
		},
		{
			TimesBefore: 13,
			Expected:    14,
		},
		{
			TimesBefore: 14,
			Expected:    15,
		},
		{
			TimesBefore: 15,
			Expected:    16,
		},
		{
			TimesBefore: 16,
			Expected:    17,
		},
		{
			TimesBefore: 17,
			Expected:    18,
		},
		{
			TimesBefore: 18,
			Expected:    19,
		},
		{
			TimesBefore: 19,
			Expected:    20,
		},



		{
			TimesBefore: 124,
			Expected:    125,
		},
		{
			TimesBefore: 125,
			Expected:    126,
		},
		{
			TimesBefore: 126,
			Expected:    127,
		},
		{
			TimesBefore: 127,
			Expected:    128,
		},
		{
			TimesBefore: 128,
			Expected:    129,
		},



		{
			TimesBefore: 250,
			Expected:    251,
		},
		{
			TimesBefore: 251,
			Expected:    252,
		},
		{
			TimesBefore: 252,
			Expected:    253,
		},
		{
			TimesBefore: 253,
			Expected:    254,
		},
		{
			TimesBefore: 254,
			Expected:    255,
		},
	}

	for testNumber, test := range tests {

		var serial autoinc.AutoInc[uint8]

		for index := uint(0); index < test.TimesBefore; index ++ {
			_, err := serial.Next()
			if nil != err {
				t.Errorf("For test #%d and index #%d, did not expect an error but actually get one.", testNumber, index)
				t.Logf("ERROR: %s", err)
				t.Logf("TIMES-BEFORE: %d", test.TimesBefore)
				continue
			}
		}

		actual, err := serial.Next()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually get one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("TIMES-BEFORE: %d", test.TimesBefore)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual next AUTO_INCREMENT value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			t.Logf("TIMES-BEFORE: %d", test.TimesBefore)
			continue
		}
	}
}

func TestAutoInc_Next_uint8_fail(t *testing.T) {

	tests := []struct{
		TimesBefore uint
		ExpectedError string
	}{
		{
			TimesBefore: 255,
			ExpectedError: "overflow",
		},
	}

	for testNumber, test := range tests {

		var serial autoinc.AutoInc[uint8]

		for index := uint(0); index < test.TimesBefore; index ++ {
			_, err := serial.Next()
			if nil != err {
				t.Errorf("For test #%d and index #%d, did not expect an error but actually get one.", testNumber, index)
				t.Logf("ERROR: %s", err)
				t.Logf("TIMES-BEFORE: %d", test.TimesBefore)
				continue
			}
		}

		_, err := serial.Next()
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("TIMES-BEFORE: %d", test.TimesBefore)
			continue
		}

		actual := err.Error()
		expected := test.ExpectedError

		if expected != actual {
			t.Errorf("For test #%d, the actual error value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("TIMES-BEFORE: %d", test.TimesBefore)
			continue
		}
	}
}
