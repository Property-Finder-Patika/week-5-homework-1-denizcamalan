package main

import "fmt"

// Package limit is 2 for every package.
// User 1 request every package
// User 2 request package is not being
// User 3 request package1 and package 3 ==> load limit
// User 4 request package1,package2 and package3, user doesn't use package1 and package3 because of sreaching the limit.



func main(){
	ProxyServer := newProxyServer()

	// --------- USER 1 --------------
	result,err := ProxyServer.clientRequest("package1")
	checkErr(result,err)
	result2,err := ProxyServer.clientRequest("package2")
	checkErr(result2,err)
	result3,err := ProxyServer.clientRequest("package3")
	checkErr(result3,err)

	// --------- USER 2 --------------
	// there is no package which name is package4
	result4,err := ProxyServer.clientRequest("package4")
	checkErr(result4,err)

	// --------- USER 3 --------------
	// the other usage request package1 and package3
	result,err = ProxyServer.clientRequest("package1")
	checkErr(result,err)
	result3,err = ProxyServer.clientRequest("package3")
	checkErr(result3,err)

	// --------- USER 4 --------------
	// the other user request package1,package2 and package3, limit is 2 so reached the limit of package1 and package3
	result,err = ProxyServer.clientRequest("package1")
	checkErr(result,err)
	result2,err = ProxyServer.clientRequest("package2")
	checkErr(result2,err)
	result3,err = ProxyServer.clientRequest("package3")
	checkErr(result3,err)
	fmt.Println(result,result3)

}

// check error if is not nil print else print value
func checkErr(str string, err error){
	if err != nil{
		fmt.Println(err)
	}else{
	fmt.Println(str)
	}
}