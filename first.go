package main

import "fmt"

//	func main() {
//		for i := 1; i <= 9; i++ {
//			for j := 1; j <= i; j++ {
//				fmt.Printf("%v*%v=%v \t", i, j, i*j)
//			}
//			fmt.Println("")
//		}
//	}
//
// func main() {
// lable1:
//
//		for i := 0; i < 2; i++ {
//			for j := 0; j < 10; j++ {
//				if j == 3 {
//					continue lable1
//				}
//				fmt.Printf("i=%v,j=%v\n", i, j)
//			}
//		}
//	}

//定义调用函数，和
//func sum(x int, y int) int {
//	num := x + y
//	return num
//}
//func main() {
//	num1 := sum(12, 3)
//	println(num1)
//
//}

// 可变参数
//
//	func sumf1(x ...int) int {
//		//x为可变参数,固定参数可以和可变参数一起，可变参数放后面
//		num := 0
//		for _, v := range x {
//			num = num + v
//		}
//		return num
//	}
//
//	func main() {
//		a := sumf1(12, 23, 66, 22) //表示把这些参数赋值给x
//		fmt.Println(a)
//	}
//
// 升序排序 封装成方法，之后再调用
func sortInAsc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}
	return slice
}
func main() {
	var sliceA = []int{12, 44, 23, 67, 32, 11, 34, 35, 40}
	a := sortInAsc(sliceA)
	fmt.Println(a)
}

//func mapsort(map1 map[string]string) string{
//	var slicekey []string
//	//1.把map对象的key放在切片里面
//	for k, _ := range map1 {
//		slicekey = append(slicekey, k)
//	}
//	//2.对key升序
//	sort.Strings(slicekey)
//}
//func main() {
//
//}
