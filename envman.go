package envman

import (
	"os"
	"strconv"
	"time"
)

var (
	// floatBox stores all variables of type float.
	floatBox map[string]float64
	// intBox stores all variables of type int.
	intBox map[string]int
	// stringBox stores all variables of type string.
	stringBox map[string]string
	// timeBox stores all variables of type time.Duration.
	timeBox map[string]time.Duration
	// uintBox stores all variables of type uint.
	uintBox map[string]uint
)

func init() {
	Clear()
}

// Clear resets the map holding all environment variables that were set.
func Clear() {
	floatBox = make(map[string]float64)
	intBox = make(map[string]int)
	stringBox = make(map[string]string)
	timeBox = make(map[string]time.Duration)
	uintBox = make(map[string]uint)
}

// Float returns a variable of type float.
func Float(key string) float64 {
	return floatBox[key]
}

// Int returns a variable of type int.
func Int(key string) int {
	return intBox[key]
}

// SetFloat stores one or more environment variables
// intended to be of type float.
//
// If a key is already stored, it is overriden.
func SetFloat(keys ...string) error {
	for _, k := range keys {
		val, err := strconv.ParseFloat(os.Getenv(k), 64)

		if err != nil {
			return err
		}

		floatBox[k] = val
	}

	return nil
}

// SetInt stores one or more environment variables
// intended to be of type int.
//
// If a key is already stored, it is overriden.
func SetInt(keys ...string) error {
	for _, k := range keys {
		val, err := strconv.Atoi(os.Getenv(k))

		if err != nil {
			return err
		}

		intBox[k] = val
	}

	return nil
}

// SetString stores one or more environment variables
// intended to be of type string.
//
// If a key is already stored, it is overriden.
func SetString(keys ...string) {
	for _, k := range keys {
		val := os.Getenv(k)
		stringBox[k] = val
	}
}

// SetTime stores one or more environment variables
// intended to be of type time.Duration.
//
// If a key is already stored, it is overriden.
func SetTime(keys ...string) error {
	for _, k := range keys {
		val, err := strconv.Atoi(os.Getenv(k))

		if err != nil {
			return err
		}

		timeBox[k] = time.Duration(val)
	}

	return nil
}

// SetUint stores one or more environment variables
// intended to be of type uint.
//
// If a key is already stored, it is overriden.
func SetUint(keys ...string) error {
	for _, k := range keys {
		val, err := strconv.ParseUint(os.Getenv(k), 10, 0)

		if err != nil {
			return err
		}

		uintBox[k] = uint(val)
	}

	return nil
}

// String returns a variable of type string.
func String(key string) string {
	return stringBox[key]
}

// Time returns a variable of type time.Duration.
func Time(key string) time.Duration {
	return timeBox[key]
}

// Uint returns a variable of type uint.
func Uint(key string) uint {
	return uintBox[key]
}
