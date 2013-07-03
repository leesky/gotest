// test
package main

import "fmt"

//func sum(a []int, c chan int) {
//	total := 0
//	for _, v := range a {
//		total += v
//	}
//	c <- total

//}

func main() {

	//a := []int{1, 2, 3, 4, 5, 6, 7, 8}

	//c := make(chan int)
	//go sum(a[:len(a)/2], c)
	//go sum(a[len(a)/2:], c)

	//x, y := <-c, <-c

	//fmt.Println(x, y)

	a := []int{1, 2, 3, 4, 5, 6}

	for _, v := range a {
		fmt.Println(v)
	}

	/*数据库test

	//db, err := sql.Open("mysql", "root:skype@/test?charset=utf8")
	//checkErr(err)

	////stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	////checkErr(err)

	////res, err := stmt.Exec("kiss", "world", "2013-07-02")
	////checkErr(err)

	////id, err := res.LastInsertId()
	////checkErr(err)

	////fmt.Println(id)

	//////更新
	////mt, err := db.Prepare("UPDATE userinfo SET username=? WHERE uid=?")
	////checkErr(err)

	////re, err := mt.Exec("pig", id)
	////checkErr(err)

	////after, err := re.RowsAffected()
	////checkErr(err)

	////fmt.Println(after)

	//up, err := db.Prepare("DELETE FROM userinfo WHERE uid=?")
	//checkErr(err)

	//rt, err := up.Exec(3)
	//checkErr(err)
	//fmt.Println(rt)

	//db.Close()
	*/

	//cookie := http.Cookie{Name: "username", Value: "lee", Expires: time.Now().AddDate(1, 0, 0)}

	//http.SetCookie(cookie)

}

//func checkErr(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
