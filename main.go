package main

import (
	// "errors"
	// "fmt"
	// "math"
	// "encoding/json"
	"fmt"
	// "time"
	// "strconv"
	"os"
	"os/exec"
)

// 9.函数

// func add(a int, b int) int {
// 	return a + b
// }

// func add2(a, b int) int {
// 	return a + b
// }

// func exists(m map[string]string, k string) (v string, ok bool) {
// 	v, ok = m[k]
// 	return v, ok
// }

// 10.指针

// func add2(n int) {
// 	n += 2
// }

// func add2ptr(n *int) {
// 	*n += 2
// }

// 11.结构体

// type user struct {
// 	name     string
// 	password string
// }

// func checkPassword(u user, password string) bool {
// 	return u.password == password
// }

// func checkPassword2(u *user, password string) bool {
// 	return u.password == password
// }

// 12.结构体方法
// type user struct {
// 	name     string
// 	password string
// }

// func (u user) checkPassword(password string) bool {
// 	return u.password == password
// }

// func (u *user) resetPassword(password string) {
// 	u.password = password
// }

// 13.错误处理

// type user struct {
// 	name     string
// 	password string
// }

// func findUser(users []user, name string) (v *user, err error) {
// 	for _, u := range users {
// 		if u.name == name {
// 			return &u, nil
// 		}
// 	}
// 	return nil, errors.New("user not found")
// }

// 15.字符串格式化
// type point struct {
// 	x, y int
// }

// 16.JSON处理

// type userInfo struct {
// 	Name  string
// 	Age   int `json:"age"`
// 	Hobby []string
// }

func main() {
	// 1.变量声明

	// var a = "initial"

	// var b, c int = 1, 2

	// var d = true

	// var e float64

	// f := float32(e)

	// g := a + "foo"
	// fmt.Println(a, b, c, d, e, f)
	// fmt.Println(g)

	// const s string = "constant"
	// const h = 500000000
	// const i = 3e20 / h
	// fmt.Println(s, h, i, math.Sin(h), math.Sin(i))

	// 2.if else

	// if 7%2 == 0 {
	// 	fmt.Println("7 is even")
	// } else {
	// 	fmt.Println("7 is odd")
	// }

	// if 8%4 == 0 {
	// 	fmt.Println("8 is divisible by 4")
	// }

	// if num := 9; num < 0 {
	// 	fmt.Println(num, "is negative")
	// } else if num < 10 {
	// 	fmt.Println(num, "has 1 digit")
	// } else {
	// 	fmt.Println(num, "has multiple digits")
	// }

	// 3.循环(仅有for)

	// i := 1
	// for n := 0; n <= 5; n++ {
	// 	if n%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(n)
	// }
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// 4.switch

	// a := 2
	// switch a {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// case 4, 5:
	// 	fmt.Println("four or five")
	// default:
	// 	fmt.Println("other")
	// }

	// t := time.Now()
	// switch {
	// case t.Hour() < 12:
	// 	fmt.Println("It's before noon")
	// default:
	// 	fmt.Println("It's after noon")
	// }

	// 5.数组

	// var a [5]int
	// a[4] = 100
	// fmt.Println(a[4], len(a))

	// b := [5]int{1, 2, 3, 4, 5}
	// fmt.Println(b)

	// var twoD [2][3]int
	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		twoD[i][j] = i + j
	// 	}
	// }
	// fmt.Println("2d: ", twoD)

	// 6.切片

	// s := make([]string, 3)
	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"
	// fmt.Println("get:", s[2])
	// fmt.Println("len:", len(s))

	// s = append(s, "d")
	// s = append(s, "e", "f")
	// fmt.Println(s)

	// c := make([]string, len(s))
	// copy(c, s)
	// fmt.Println(c)

	// fmt.Println(s[2:5])
	// fmt.Println(s[:5])
	// fmt.Println(s[2:])

	// good := []string{"g", "o", "o", "d"}
	// fmt.Println(good)

	// 7.map

	// m := make(map[string]int)
	// m["one"] = 1
	// m["two"] = 2
	// fmt.Println(m)
	// fmt.Println(len(m))
	// fmt.Println(m["one"])
	// fmt.Println(m["unknow"])

	// r, ok := m["unknow"]
	// fmt.Println(r, ok)

	// delete(m, "one")

	// m2 := map[string]int{"one": 1, "two": 2}
	// var m3 = map[string]int{"one": 1, "two": 2}
	// fmt.Println(m2, m3)

	// 8.range

	// nums := []int{2, 3, 4}
	// sum := 0
	// for i, num := range nums {
	// 	sum += num
	// 	if num == 2 {
	// 		fmt.Println("index:", i, "num:", num)
	// 	}
	// }
	// fmt.Println(sum)

	// m := map[string]string{"a": "A", "b": "B"}
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }
	// for k := range m {
	// 	fmt.Println("key:", k)
	// }

	// 9.函数

	// res := add(1, 2)
	// fmt.Println(res)

	// v, ok := exists(map[string]string{"a": "A"}, "a")
	// fmt.Println(v, ok)

	// 10.指针

	// n := 5
	// add2(n)
	// fmt.Println(n)
	// add2ptr(&n)
	// fmt.Println(n)

	// 11.结构体

	// a := user{name: "wang", password: "1024"}
	// b := user{"wang", "1024"}
	// c := user{name: "wang"}
	// c.password = "1024"
	// var d user
	// d.name = "wang"
	// d.password = "1024"

	// fmt.Println(a, b, c, d)
	// fmt.Println(checkPassword(a, "haha"))
	// fmt.Println(checkPassword2(&a, "haha"))

	// 12.结构体方法

	// a := user{name: "wang", password: "1024"}
	// a.resetPassword("2048")
	// fmt.Println(a.checkPassword("2048"))

	// 13.错误处理

	// u, err := findUser([]user{{"wang", "1024"}}, "wang")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(u.name)

	// if u, err := findUser([]user{{"wang", "1024"}}, "li"); err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Println(u.name)
	// }

	// 14.字符串操作

	// a := "hello"
	// fmt.Println(strings.Contains(a, "ll"))
	// fmt.Println(strings.Count(a, "l"))
	// fmt.Println(strings.HasPrefix(a, "he"))
	// fmt.Println(strings.HasSuffix(a, "llo"))
	// fmt.Println(strings.Index(a, "ll"))
	// fmt.Println(strings.Join([]string{"he", "llo"}, "-"))
	// fmt.Println(strings.Repeat(a, 2))
	// fmt.Println(strings.Replace(a, "e", "E", -1))
	// fmt.Println(strings.Split("a-b-c", "-"))
	// fmt.Println(strings.ToLower(a))
	// fmt.Println(strings.ToUpper(a))
	// fmt.Println(len(a))
	// b := "你好"
	// fmt.Println(len(b))

	// 15.字符串格式化

	// s := "hello"
	// n := 123
	// p := point{1, 2}
	// fmt.Println(s, n)
	// fmt.Println(p)

	// fmt.Printf("s=%v\n", s)
	// fmt.Printf("n=%v\n", n)
	// fmt.Printf("p=%v\n", p)
	// fmt.Printf("p=%+v\n", p)
	// fmt.Printf("p=%#v\n", p)

	// f := 3.141592653
	// fmt.Println(f)
	// fmt.Printf("%.2f\n", f)

	// 16.JSON处理

	// a := userInfo{Name: "wang", Age: 18, Hobby: []string{"Golang", "TypeScript"}}
	// buf, err := json.Marshal(a)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(buf)
	// fmt.Println(string(buf))

	// buf, err = json.MarshalIndent(a, "", "\t")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(buf))

	// var b userInfo
	// err = json.Unmarshal(buf, &b)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%#v\n", b)

	// 17.时间处理

	// now := time.Now()
	// fmt.Println(now)
	// t := time.Date(2022, 3, 27, 1, 25, 36, 0, time.UTC)
	// t2 := time.Date(2022, 3, 27, 2, 30, 36, 0, time.UTC)
	// fmt.Println(t)
	// fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
	// fmt.Println(t.Format("2006-01-02 15:04:05"))
	// diff := t2.Sub(t)
	// fmt.Println(diff)
	// fmt.Println(diff.Minutes(), diff.Seconds())
	// t3, err := time.Parse("2006-01-02 15:04:05", "2022-03-27 01:25:36")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(t3 == t)
	// fmt.Println(now.Unix())

	// 18.数字解析

	// f, _ := strconv.ParseFloat("1.234", 64)
	// fmt.Println(f)

	// n, _ := strconv.ParseInt("111", 10, 64)
	// fmt.Println(n)

	// n, _ = strconv.ParseInt("0x1000", 0, 64)
	// fmt.Println(n)

	// n2, _ := strconv.Atoi("123")
	// fmt.Println(n2)

	// n2, err := strconv.Atoi("AAA")
	// fmt.Println(n2, err)

	// 19.进程信息
	fmt.Println(os.Args)
	fmt.Println(os.Getenv("PATH"))
	fmt.Println(os.Setenv("AA", "BB"))

	buf, err := exec.Command("grep", "1.2.3.4", "/etc/hosts").CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
}
