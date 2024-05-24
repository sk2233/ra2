/*
@author: sk
@date: 2023/9/2
*/
package utils

import (
	"strconv"

	"github.com/beevik/etree"
)

func OpenXml(path string) *etree.Element {
	res := etree.NewDocument()
	err := res.ReadFromFile(path)
	HandleErr(err)
	return res.Root()
}

func XmlStr(elem *etree.Element, name, def string) string {
	return elem.SelectAttrValue(name, def)
}

func XmlUint(elem *etree.Element, name string, def uint64) uint64 {
	if res, err := strconv.ParseUint(elem.SelectAttrValue(name, ""), 10, 64); err == nil {
		return res
	}
	return def
}

func XmlInt(elem *etree.Element, name string, def int64) int64 {
	if res, err := strconv.ParseInt(elem.SelectAttrValue(name, ""), 10, 64); err == nil {
		return res
	}
	return def
}

func XmlFloat(elem *etree.Element, name string, def float64) float64 {
	if res, err := strconv.ParseFloat(elem.SelectAttrValue(name, ""), 64); err == nil {
		return res
	}
	return def
}
