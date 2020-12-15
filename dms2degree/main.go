// main
//本程序将执行坐标点的进制转换
//dms=度分秒,默认通用输入格式为:dd.mmsss
//dd为整数度，mm为整数分，sss为ss.s秒
//角度制为(degree) 即 dd.dddddddd
//弧度制为(radian) 即 dd.dddddddd/180*Pi
//输入的坐标点以逗号(comma)分隔,例如:a11,27.5645394,112.10119848,95.46
//排列顺序为:点名，纬度，经度，大地高
//暂未对东西南北半球进行标记,默认为北纬、东经

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

//定义点的结构体
type PointBLH struct {
	name       string
	latitude   string
	longtitude string
	height     float64
}

//函数：根据读入的数据创建点结构体
func BuildPointBLH(point string) (pointblh *PointBLH) {

	comma := make([]int, 0, 1)

	for k, v := range point {
		if v == ',' {
			comma = append(comma, k)
		}
	}

	pointblh = new(PointBLH)

	pointblh.name = string(point[0:(comma[0])])
	pointblh.latitude = string(point[(comma[0])+1 : (comma[1])])
	pointblh.longtitude = string(point[(comma[1])+1 : (comma[2])])
	pointblh.height, _ = strconv.ParseFloat(point[(comma[2])+1:], 64)

	return

}

//函数：对dms格式的数据进行分割为 度dd  分mm 秒ss.s
func SpiltDms(dms string) (dd, mm int, sss float64) {
	for k, v := range dms {
		if v == '.' {
			dd, _ = strconv.Atoi(dms[0:k])
			mm, _ = strconv.Atoi(dms[k+1 : k+3])
			sss, _ = strconv.ParseFloat(dms[k+3:k+5]+"."+dms[k+5:], 64)
		}
	}
	return
}

//函数：对dms格式的数据进行转换为 度(角度制)degree
func Dms2Degree(dms string) (degree float64) {
	dd, mm, sss := SpiltDms(dms)
	degree = float64(dd) + (float64(mm)+(sss/60))/60
	return
}

//函数：将角度制(degree)转换为弧度制(radian)
func Degree2Radian(degree float64) (radian float64) {
	radian = degree / 180 * math.Pi
	return
}

func main() {

	//---------------------------------------------------
	fileNamein, err := os.Open("./coordout.txt")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer fileNamein.Close()
	pointin := bufio.NewReader(fileNamein)
	//---------------------------------------------------

	for {
		s1, _, err := pointin.ReadLine()

		if err == io.EOF {
			fmt.Println("read over !")
			return
		}

		s2 := BuildPointBLH(string(s1))

		name := s2.name
		latitude := Dms2Degree(s2.latitude)
		longtitude := Dms2Degree(s2.longtitude)
		height := s2.height

		fmt.Println(name, latitude, longtitude, height)

	}

}
