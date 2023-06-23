package main

import (
	"fmt"
)

type celsius float64    //摄氏
type fahrenheit float64 //华氏
type kelvin float64     //开氏

// 摄氏度转华氏度  c to h
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

// 摄氏度转开氏度   c to k
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

// 华氏度转摄氏度  h to c
func (f fahrenheit) celsius() celsius {
	return celsius(f-32.0) * 5.0 / 9.0
}

// 华氏度转开氏度  f to k
func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

// 开氏度转摄氏度   k to c
func (k kelvin) celsius() celsius { //  ktoc(k kelvin) celsius {  //  函数名（函数接收者 接收者类型 ）函数类型//
	//	return celsius(k - 273.15)
	return celsius(k - 273.15)
}

// 开氏度转华氏度   k to f
func (k kelvin) fahrenheit() fahrenheit { //(方法接收者 接收者类型) 方法()方法类型  方法返回类型
	return k.celsius().fahrenheit() //return k.kelvin().fahrenheit
}
func main() {
	var k kelvin = 294.0
	//c := ktoc(k)
	c := k.celsius()
	f := k.fahrenheit()

	var d celsius = 294.0
	e := d.kelvin()
	g := d.fahrenheit()

	fmt.Println(d, "is", e)
	fmt.Print(d, "ºK is ", g, "ºc")

	fmt.Println(k, "is", f)
	fmt.Print(k, "ºK is ", c, "ºc")
}
