package envman

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFloat(t *testing.T) {
	a := assert.New(t)
	tests := []*struct {
		shouldBeSet bool
		varname     string
		expectedVal float64
	}{
		// #0
		{
			shouldBeSet: true,
			varname:     "TEST",
			expectedVal: 1.0,
		},
		// #1
		{
			shouldBeSet: false,
			varname:     "TEST",
			expectedVal: 0.0,
		},
	}

	for i, test := range tests {
		index := strconv.Itoa(i)

		if test.shouldBeSet {
			floatBox[test.varname] = test.expectedVal
		}

		a.Exactly(test.expectedVal, Float(test.varname), index)
		Clear()
	}
}

func TestInt(t *testing.T) {
	a := assert.New(t)
	tests := []*struct {
		shouldBeSet bool
		varname     string
		expectedVal int
	}{
		// #0
		{
			shouldBeSet: true,
			varname:     "TEST",
			expectedVal: -1,
		},
		// #1
		{
			shouldBeSet: false,
			varname:     "TEST",
			expectedVal: 0,
		},
	}

	for i, test := range tests {
		index := strconv.Itoa(i)

		if test.shouldBeSet {
			intBox[test.varname] = test.expectedVal
		}

		a.Exactly(test.expectedVal, Int(test.varname), index)
		Clear()
	}
}

func TestSetFloat(t *testing.T) {
	a := assert.New(t)
	tests := []*struct {
		key      string
		val      string
		expected bool
	}{
		// #0
		{
			key:      "TEST",
			val:      "1",
			expected: true,
		},
		// #1
		{
			key:      "TEST",
			val:      "123.456",
			expected: true,
		},
		// #2
		{
			key:      "TEST",
			val:      "-1",
			expected: true,
		},
		// #3
		{
			key:      "TEST",
			val:      "test",
			expected: false,
		},
	}

	for i, test := range tests {
		index := strconv.Itoa(i)

		err := os.Setenv(test.key, test.val)

		a.Nil(err)

		err = SetFloat(test.key)

		a.Exactly(test.expected, err == nil, index)

		if err == nil {
			a.IsType(0.0, Float(test.key), index)
		}

		Clear()
	}
}

func TestSetInt(t *testing.T) {
	a := assert.New(t)
	tests := []*struct {
		key      string
		val      string
		expected bool
	}{
		// #0
		{
			key:      "TEST",
			val:      "1",
			expected: true,
		},
		// #1
		{
			key:      "TEST",
			val:      "123.456",
			expected: false,
		},
		// #2
		{
			key:      "TEST",
			val:      "-1",
			expected: true,
		},
		// #3
		{
			key:      "TEST",
			val:      "test",
			expected: false,
		},
	}

	for i, test := range tests {
		index := strconv.Itoa(i)

		err := os.Setenv(test.key, test.val)

		a.Nil(err)

		err = SetInt(test.key)

		a.Exactly(test.expected, err == nil, index)

		if err == nil {
			a.IsType(-1, Int(test.key), index)
		}

		Clear()
	}
}

func TestSetString(t *testing.T) {
	a := assert.New(t)
	tests := []*struct {
		key      string
		val      string
		expected bool
	}{
		// #0
		{
			key:      "TEST",
			val:      "1",
			expected: true,
		},
		// #1
		{
			key:      "TEST",
			val:      "123.456",
			expected: true,
		},
		// #2
		{
			key:      "TEST",
			val:      "-1",
			expected: true,
		},
		// #3
		{
			key:      "TEST",
			val:      "test",
			expected: true,
		},
	}

	for i, test := range tests {
		index := strconv.Itoa(i)

		err := os.Setenv(test.key, test.val)

		a.Nil(err)
		SetString(test.key)
		a.IsType("0", String(test.key), index)

		Clear()
	}
}

func TestSetTime(t *testing.T) {
	a := assert.New(t)
	tests := []*struct {
		key      string
		val      string
		expected bool
	}{
		// #0
		{
			key:      "TEST",
			val:      "1",
			expected: true,
		},
		// #1
		{
			key:      "TEST",
			val:      "123.456",
			expected: false,
		},
		// #2
		{
			key:      "TEST",
			val:      "-1",
			expected: true,
		},
		// #3
		{
			key:      "TEST",
			val:      "test",
			expected: false,
		},
	}

	for i, test := range tests {
		index := strconv.Itoa(i)

		err := os.Setenv(test.key, test.val)

		a.Nil(err)

		err = SetTime(test.key)

		a.Exactly(test.expected, err == nil, index)

		if err == nil {
			a.IsType(time.Duration(0), Time(test.key), index)
		}

		Clear()
	}
}

func TestSetUint(t *testing.T) {
	a := assert.New(t)
	tests := []*struct {
		key      string
		val      string
		expected bool
	}{
		// #0
		{
			key:      "TEST",
			val:      "1",
			expected: true,
		},
		// #1
		{
			key:      "TEST",
			val:      "123.456",
			expected: false,
		},
		// #2
		{
			key:      "TEST",
			val:      "-1",
			expected: false,
		},
		// #3
		{
			key:      "TEST",
			val:      "test",
			expected: false,
		},
	}

	for i, test := range tests {
		index := strconv.Itoa(i)

		err := os.Setenv(test.key, test.val)

		a.Nil(err)

		err = SetUint(test.key)

		a.Exactly(test.expected, err == nil, index)

		if err == nil {
			a.IsType(uint(1), Uint(test.key), index)
		}

		Clear()
	}
}

func TestString(t *testing.T) {
	a := assert.New(t)
	tests := []*struct {
		shouldBeSet bool
		varname     string
		expectedVal string
	}{
		// #0
		{
			shouldBeSet: true,
			varname:     "TEST",
			expectedVal: "test",
		},
		// #1
		{
			shouldBeSet: false,
			varname:     "TEST",
			expectedVal: "",
		},
	}

	for i, test := range tests {
		index := strconv.Itoa(i)

		if test.shouldBeSet {
			stringBox[test.varname] = test.expectedVal
		}

		a.Exactly(test.expectedVal, String(test.varname), index)
		Clear()
	}
}

func TestUint(t *testing.T) {
	a := assert.New(t)
	tests := []*struct {
		shouldBeSet bool
		varname     string
		expectedVal uint
	}{
		// #0
		{
			shouldBeSet: true,
			varname:     "TEST",
			expectedVal: 1,
		},
		// #1
		{
			shouldBeSet: false,
			varname:     "TEST",
			expectedVal: 0,
		},
	}

	for i, test := range tests {
		index := strconv.Itoa(i)

		if test.shouldBeSet {
			uintBox[test.varname] = test.expectedVal
		}

		a.Exactly(test.expectedVal, Uint(test.varname), index)
		Clear()
	}
}
