package envman

import (
	"os"
	"strconv"
	"time"
)

// envbox stores all the environment variables stored by the setters.
var envbox map[string]interface{}

func init() {
	Clear()
}

// Get returns an already stored value or nil.
func Get(key string) interface{} {
	return envbox[key]
}

// Clear resets the map holding all environment variables that were set.
func Clear() {
	envbox = make(map[string]interface{})
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

		envbox[k] = val
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

		envbox[k] = val
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
		envbox[k] = val
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

		envbox[k] = time.Duration(val)
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

		envbox[k] = val
	}

	return nil
}
