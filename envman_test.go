package envman

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	a := assert.New(t)
	tests := []*struct {
		shouldBeSet bool
		varname     string
		expectedVal interface{}
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
			expectedVal: nil,
		},
	}

	for i, test := range tests {
		index := strconv.Itoa(i)

		if test.shouldBeSet {
			envbox[test.varname] = test.expectedVal
		}

		a.Exactly(test.expectedVal, Get(test.varname), index)
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
			a.IsType(0.0, Get(test.key), index)
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
			a.IsType(-1, Get(test.key), index)
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
		a.IsType("0", Get(test.key), index)

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
			a.IsType(time.Duration(0), Get(test.key), index)
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
			a.IsType(uint64(1), Get(test.key), index)
		}

		Clear()
	}
}
