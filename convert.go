package ky3k

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"reflect"
)

// 转换成 Decimal
func ConvertToDecimal(s interface{}) (decimal.Decimal, error) {
	var err error
	var d decimal.Decimal
	switch reflect.TypeOf(s).Kind() {
	case reflect.String:
		m, ok := s.(string)
		if ok {
			d, err = decimal.NewFromString(m)
		} else {
			err = errors.New(fmt.Sprintf("类型不匹配:%v", s))
		}
		break
	case reflect.Float64:
		m, ok := s.(float64)
		if ok {
			d = decimal.NewFromFloat(m)
		} else {
			err = errors.New(fmt.Sprintf("类型不匹配:%v", s))
		}
		break
	case reflect.Int64:
		m, ok := s.(int64)
		if ok {
			d = decimal.NewFromInt(m)
		} else {
			err = errors.New(fmt.Sprintf("类型不匹配:%v", s))
		}
		break
	case reflect.Int:
		m, ok := s.(int)
		if ok {
			d = decimal.NewFromInt(int64(m))
		} else {
			err = errors.New(fmt.Sprintf("类型不匹配:%v", s))
		}
		break
	default:
		err = errors.New(fmt.Sprintf("类型不匹配:%v", s))
	}
	return d, err
}

//
func ConvertYuanToFen(s float64) int64 {

	return decimal.NewFromFloat(s).
		Mul(decimal.NewFromInt(100)).
		IntPart()
}

// 毫金额转换元 四舍五入保留两位小数
func ConvertFenToYuan(s int64) float64 {

	d, _ := decimal.NewFromFloat(float64(s)).
		DivRound(decimal.NewFromFloat(100), 2).
		Float64()

	return d
}

// 计量单位转换毫米  米->毫米
func ConvertMToMM(s interface{}) (int64, error) {
	d, err := ConvertToDecimal(s)
	m := decimal.NewFromInt(100).Mul(d).IntPart()
	return m, err
}

// 计量单位转换为毫米 厘米->毫米
func ConvertCMToMM(s interface{}) (int64, error) {
	d, err := ConvertToDecimal(s)
	m := decimal.NewFromInt(10).Mul(d).IntPart()
	return m, err
}

// 重量单位转换毫克  千克->毫克  (1毫克(mg)=0.000001千克(kg)
func ConvertKGTOMG(s interface{}) (int64, error) {
	d, err := ConvertToDecimal(s)
	m := decimal.NewFromInt(1000000).Mul(d).IntPart()
	return m, err
}

// 重量单位转换毫克  克->毫克  (1毫克(mg)=0.001克(g)
func ConvertGTOMG(s interface{}) (int64, error) {
	d, err := ConvertToDecimal(s)
	m := decimal.NewFromInt(1000).Mul(d).IntPart()
	return m, err
}

//重量单位转换毫克  毫克->千克  1毫克(mg)=0.001克(g)
func ConvertMGToKG(s int64) (float64, bool) {
	d, b := decimal.NewFromFloat(float64(s)).
		Div(decimal.NewFromFloat(1000000)).
		Round(4).
		Float64()

	return d, b
}

// 重量单位转换毫克  毫克->克  1毫克(mg)=0.001克(g)
func ConvertMGToG(s int64) (float64, bool) {
	d, b := decimal.NewFromFloat(float64(s)).
		Div(decimal.NewFromFloat(1000)).
		Round(4).
		Float64()

	return d, b
}
