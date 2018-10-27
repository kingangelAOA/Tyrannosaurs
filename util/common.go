package util

import (
	"reflect"
	. "tyrannosaurs/constant"
	"regexp"
	"fmt"
	"strings"
	"runtime"
	"bytes"
	"strconv"
)

const replaseFlag = "${"

func SliceExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic(SliceExistsError)
	}
	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}
	return false
}

func MapKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func ReplaceData(data *string, cache map[string]string) error {
	if !strings.Contains(*data, replaseFlag) {
		return nil
	}
	rgx := regexp.MustCompile(`\$\{(.*?)\}`)
	matches := rgx.FindAllStringSubmatch(*data, -1)
	for _, match := range matches {
		matchContent := match[0]
		key := match[1]
		value, ok := cache[key]
		if !ok {
			return fmt.Errorf("cache has not key '%s' to replase", key)
		}
		*data = strings.Replace(*data, matchContent, value, -1)
	}
	return nil
}

func ReplaceDataHasResult(data string, cache map[string]string) (string, error) {
	if !strings.Contains(data, replaseFlag) {
		return "", nil
	}
	rgx := regexp.MustCompile(`\$\{(.*?)\}`)
	matches := rgx.FindAllStringSubmatch(data, -1)
	for _, match := range matches {
		matchContent := match[0]
		key := match[1]
		value, ok := cache[key]
		if !ok {
			return "", fmt.Errorf("cache has not key '%s' to replase", key)
		}
		data = strings.Replace(data, matchContent, value, -1)
	}
	return data, nil
}

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}