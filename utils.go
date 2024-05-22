package utils

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func StatusText(code int) string {
	m := map[int]string{
		OK:                 "Success",
		BAD_REQUEST:        "Bad request, invalid parameters supplied.",
		LOGIN_FAILED:       "Login failed, please check your credentials, contact with administrator for help if still failed.",
		SESSION_EXPIRED:    "Session expired, please login again.",
		PERMISSION_DENIED:  "Permission denied.",
		NOT_FOUND:          "The requested resource could not be found.",
		SERVER_ERROR:       "An error occurred on the server when processing the URL. Please contact the system administrator.",
		GRPC_SERVICE_ERROR: "An error occurred when calling a gRPC service. Please contact the system administrator.",
		MYSQL_ERROR:        "An error occurred during connecting with MySQL, please contact with system administrator.",
	}
	if v, found := m[code]; found {
		return v
	}
	return ""
}

//AsOptions convert [a] to [a,a]
func AsOptions(list []string) [][]interface{} {
	output := [][]interface{}{}
	for _, item := range list {
		output = append(output, []interface{}{item, item})
	}
	return output
}

// Split by ",", ";"
func Split(s string) []string {
	out := strings.FieldsFunc(s, func(r rune) bool {
		switch r {
		case ',', ';':
			return true
		}
		return false
	})
	return out
}

func SliceToMap(colNames []string, rows [][]string) ([]map[string]string, error) {
	if len(rows) == 0 {
		return nil, nil
	}
	colNum := len(colNames)
	rowNum := len(rows)

	// create maps
	results := []map[string]string{}
	for i := 0; i < rowNum; i++ {
		if len(rows[i]) != colNum {
			return nil, errors.New("Error: ColNames length and record length not consistent.")
		}
		t_map := map[string]string{}
		for j := 0; j < colNum; j++ {
			t_map[colNames[j]] = rows[i][j]
		}
		results = append(results, t_map)
	}

	return results, nil
}

// SliceToMap Interface{} instead of string
func SliceToMapInterface(colNames []string, rows [][]interface{}) ([]map[string]interface{}, error) {
	if len(rows) == 0 {
		return nil, nil
	}
	colNum := len(colNames)
	rowNum := len(rows)

	// create maps
	results := []map[string]interface{}{}
	for i := 0; i < rowNum; i++ {
		if len(rows[i]) != colNum {
			return nil, errors.New("Error: ColNames length and record length not consistent.")
		}
		t_map := map[string]interface{}{}
		for j := 0; j < colNum; j++ {
			t_map[colNames[j]] = rows[i][j]
		}
		results = append(results, t_map)
	}

	return results, nil
}

func ReadUrl(apiUrl, user, pwd string, proxyUrl string) ([]byte, error) {
	var client = &http.Client{}
	var tr *http.Transport
	if proxyUrl != "" {
		proxy, _ := url.Parse(proxyUrl)
		tr = &http.Transport{
			Proxy:           http.ProxyURL(proxy),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	} else {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	client = &http.Client{
		Transport: tr,
	}

	//read body string
	req, err := http.NewRequest("GET", apiUrl, nil)
	req.SetBasicAuth(user, pwd)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	// Retrieve the body of the response
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		return nil, err
	}
	return []byte(body), nil
}

//StrInSlice check if string is in string list
func StrInSlice(str string, list []string) bool {
	for _, item := range list {
		if str == item {
			return true
		}
	}
	return false
}

// Remove string from a string list
func RemoveStrFromSlice(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func SplitByPipe(s string) []string {
	out := strings.FieldsFunc(s, func(r rune) bool {
		switch r {
		case '|':
			return true
		}
		return false
	})
	return out
}

//UniqueInts returns a unique subset of the int slice provided.
func UniqueInts(input []int) []int {
	u := make([]int, 0, len(input))
	m := make(map[int]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

//UniqueStrings returns a unique subset of the int slice provided.
func UniqueStrings(input []string) []string {
	u := make([]string, 0, len(input))
	m := make(map[string]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

//RemoveBlankStrings
func RemoveBlankStrings(input []string) []string {
	u := make([]string, 0, len(input))
	for _, val := range input {
		newVal := strings.TrimSpace(val)
		if newVal != "" {
			u = append(u, newVal)
		}
	}

	return u
}

func FilterUnicodeSymbol(s string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case 0x000A, 0x000B, 0x000C, 0x000D, 0x0085, 0x2028, 0x2029, 0xFFFD:
			return -1
		default:
			return r
		}
	}, s)
}

// random seed
func initRandSeed() {
	rand.Seed(time.Now().UnixNano())
}

// rando string charts
var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// gen randomstring with size input
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// UniqueSlice remote duplicate element from slice
func UniqueSlice(s []string) []string {
	unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
	for _, item := range s {
		elem := strings.TrimSpace(item)
		if len(elem) != 0 {
			if !unique[elem] {
				us = append(us, elem)
				unique[elem] = true
			}
		}
	}
	return us
}

func IsValidString(obj string) bool {
	return obj != "" && strings.TrimSpace(obj) != ""
}

// TrimSuffix remove trailing plus sign(s).
func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func ToUrlString(originStr, urlPath string) string {
	if !IsValidUrlString(originStr) {
		return ""
	}

	var urlStr string
	tempU, err := url.Parse(urlPath)
	if err != nil { // invalid url, only path
		if strings.HasPrefix(urlPath, "/") {
			urlStr = strings.TrimSuffix(originStr, "/") + urlPath
		} else {
			urlStr = strings.TrimSuffix(originStr, "/") + "/" + urlPath
		}
	} else {
		urlPathStr := urlPath[strings.Index(urlPath, tempU.Host)+len(tempU.Host):]
		if strings.HasPrefix(urlPathStr, "/") {
			urlStr = strings.TrimSuffix(originStr, "/") + urlPathStr
		} else {
			urlStr = strings.TrimSuffix(originStr, "/") + "/" + urlPathStr
		}
	}

	return urlStr
}

func IsValidUrlString(urlStr string) bool {
	_, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return false
	}
	return true
}
