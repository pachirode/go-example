package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"Alice", 30}

	fmt.Printf("%%v 默认格式：%v\n", p)
	fmt.Printf("%%+v 包含字段名：%+v\n", p)
	fmt.Printf("%%#v Go 语法表示：%#v\n", p)
	fmt.Printf("%%T 类型表示：%T\n", p)
	fmt.Printf("%%%% 百分号：%%\n")

	fmt.Printf("%%t 布尔值：%t\n", true)
	{
		num := 255
		fmt.Printf("%%b 二进制：%b\n", num)
		fmt.Printf("%%c 字符：%c\n", num)
		fmt.Printf("%%d 十进制：%d\n", num)
		fmt.Printf("%%o 八进制：%o\n", num)
		fmt.Printf("%%O 带前缀的八进制：%O\n", num)
		fmt.Printf("%%q 单引号字符：%q\n", num)
		fmt.Printf("%%x 小写十六进制：%x\n", num)
		fmt.Printf("%%X 大写十六进制：%X\n", num)
		fmt.Printf("%%U Unicode 格式：%U\n", num)
	}

	{
		num := 12345.6789
		complexNum := 1.234 + 5.678i

		fmt.Printf("%%b 无小数点科学计数法：%b\n", num)
		fmt.Printf("%%e 科学计数法（小写 e）：%e\n", num)
		fmt.Printf("%%E 科学计数法（大写 E）：%E\n", num)
		fmt.Printf("%%f 十进制形式：%f\n", num)
		fmt.Printf("%%F 与 %%f 相同：%F\n", num)
		fmt.Printf("%%g 自动选择：%g\n", num)
		fmt.Printf("%%G 自动选择（大写）：%G\n", num)
		fmt.Printf("%%x 十六进制表示：%x\n", num)
		fmt.Printf("%%X 十六进制表示（大写）：%X\n", num)
		fmt.Println("----------")
		fmt.Printf("%%b 无小数点科学计数法（复数）：%b\n", complexNum)
		fmt.Printf("%%e 科学计数法（小写 e）（复数）：%e\n", complexNum)
		fmt.Printf("%%E 科学计数法（大写 E）（复数）：%E\n", complexNum)
		fmt.Printf("%%f 十进制形式（复数）：%f\n", complexNum)
		fmt.Printf("%%F 与 %%f 相同（复数）：%F\n", complexNum)
		fmt.Printf("%%g 自动选择（复数）：%g\n", complexNum)
		fmt.Printf("%%G 自动选择（大写）（复数）：%G\n", complexNum)
		fmt.Printf("%%x 十六进制表示（复数）：%x\n", complexNum)
		fmt.Printf("%%X 十六进制表示（大写）（复数）：%X\n", complexNum)
	}

	{
		str := "hello"
		str1 := `"hello"`
		data := []byte{72, 101, 108, 108, 111}

		fmt.Printf("%%s 原始字节：%s\n", str)
		fmt.Printf("%%q 双引号字符串：%q\n", str)
		fmt.Printf("%%q 双引号字符串：%q\n", str1)
		fmt.Printf("%%x 小写十六进制：%x\n", data)
		fmt.Printf("%%X 大写十六进制：%X\n", data)
	}
}
