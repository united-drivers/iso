package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"text/template"

	"github.com/moul/funcmap"
)

var tmplStr = `syntax = "proto3";

/** Currency codes - ISO 4217
 *
 * http://www.iso.org/iso/home/standards/currency_codes.htm
 */

package currency;

enum CurrencyCode {
  _ = 0; // zero vale

{{range .}}{{if .CurrencyNumber}}/**
 * {{.CurrencyName}}
 *
{{range .CountryNames}} * * {{.}}
{{end}}
 */
  {{.CurrencyCode}} = {{.CurrencyNumber}};

{{end}}{{end}}}
`

type currency struct {
	CurrencyName   string
	CurrencyCode   string
	CurrencyNumber int
	CountryNames   []string
	MajorExponent  *int
}

func main() {
	inputStr, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	var input interface{}
	if err := json.Unmarshal(inputStr, &input); err != nil {
		panic(err)
	}

	currencies := map[string]currency{}
	inputCurrencies := ((((input.(map[string]interface{}))["ISO_4217"].(map[string]interface{}))["CcyTbl"]).(map[string]interface{})["CcyNtry"]).([]interface{})
	for _, entryInterface := range inputCurrencies {
		entry := entryInterface.(map[string]interface{})
		if entry["CcyNbr"] == nil {
			continue
		}
		newCurrency := currency{}
		countryName := nodeToString(entry["CtryNm"])
		currencyNumber := nodeToString(entry["CcyNbr"])
		newCurrency.CurrencyNumber, _ = strconv.Atoi(currencyNumber)
		if entry["CcyNm"] != nil {
			newCurrency.CurrencyName = nodeToString(entry["CcyNm"])
		}
		if entry["Ccy"] != nil {
			newCurrency.CurrencyCode = nodeToString(entry["Ccy"])
		}
		newCurrency.CountryNames = []string{countryName}
		if existing, found := currencies[newCurrency.CurrencyCode]; found {
			newCurrency.CountryNames = append(existing.CountryNames, countryName)
		}
		currencies[newCurrency.CurrencyCode] = newCurrency
	}

	tmpl, err := template.New("").Funcs(funcmap.FuncMap).Parse(tmplStr)
	if err != nil {
		panic(err)
	}

	if err := tmpl.Execute(os.Stdout, currencies); err != nil {
		panic(err)
	}
}

func nodeToString(node interface{}) string {
	if str, ok := node.(string); ok {
		return str
	}
	return node.(map[string]interface{})["$t"].(string)
}
