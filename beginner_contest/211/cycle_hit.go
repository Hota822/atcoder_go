package main

import (
	"fmt"
	// "bufio"
	// "os"
	"strings"
	// "strconv"
	// "reflect"
)

func main() {
	var s1, s2, s3, s4 string
	fmt.Scan(&s1, &s2, &s3, &s4)
	sum_s := s1 + s2 + s3 + s4
	if strings.Contains(sum_s, "H") && strings.Contains(sum_s, "2B") && strings.Contains(sum_s, "3B") && strings.Contains(sum_s, "HR") && len(sum_s) == 7 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

// func main() {
// 	var s1, s2, s3, s4 string
// 	fmt.Scan(&s1, &s2, &s3, &s4)
// 	if s1 == "H" {
// 		if no3(s2, s3, s4, "H") {
// 			fmt.Println("No")
// 			return
// 		}
// 		if s2 == "2B" {
// 			if no2(s3, s4, "2B") {
// 				fmt.Println("No")
// 				return
// 			} else if same(s3, s4) {
// 				fmt.Println("No")
// 				return
// 			}
// 		} else if s2 == "3B" {
// 			if no2(s3, s4, "3B") {
// 				fmt.Println("No")
// 				return
// 			} else if same(s3, s4) {
// 				fmt.Println("No")
// 				return
// 			}
// 		} else {
// 			if no2(s3, s4, "HR") {
// 				fmt.Println("No")
// 				return
// 			} else if same(s3, s4) {
// 				fmt.Println("No")
// 				return
// 			}
// 		}
// 	} else if  s1 == "2B" {
// 		if no3(s2, s3, s4, "2B") {
// 			fmt.Println("No")
// 			return
// 		}
// 		if s2 == "H" {
// 			if no2(s3, s4, "H") {
// 				fmt.Println("No")
// 				return
// 			} else if same(s3, s4) {
// 				fmt.Println("No")
// 				return
// 			}
// 		} else if s2 == "3B" {
// 			if no2(s3, s4, "3B") {
// 				fmt.Println("No")
// 				return
// 			} else if same(s3, s4) {
// 				fmt.Println("No")
// 				return
// 			}
// 		} else {
// 			if no2(s3, s4, "HR") {
// 				fmt.Println("No")
// 				return
// 			} else if same(s3, s4) {
// 				fmt.Println("No")
// 				return
// 			}
// 		}
// 	} else if s1 == "3B" {
// 		if no3(s2, s3, s4, "3B") {
// 			fmt.Println("No")
// 			return
// 		}
// 		if s2 == "H" {
// 			if no2(s3, s4, "H") {
// 				fmt.Println("No")
// 				return
// 			} else if same(s3, s4) {
// 				fmt.Println("No")
// 				return
// 			}
// 		} else if s2 == "2B" {
// 			if no2(s3, s4, "2B") {
// 				fmt.Println("No")
// 				return
// 			} else if same(s3, s4) {
// 				fmt.Println("No")
// 				return
// 			}
// 		} else {
// 			if no2(s3, s4, "HR") {
// 				fmt.Println("No")
// 				return
// 			} else if same(s3, s4) {
// 				fmt.Println("No")
// 				return
// 			}
// 		}
// 	} else {
// 		if no3(s2, s3, s4, "HR") {
// 			fmt.Println("No")
// 			return
// 		}
// 		if s2 == "H" {
// 			if no2(s3, s4, "H") {
// 				fmt.Println("No")
// 				return
// 			} else if same(s3, s4) {
// 				fmt.Println("No")
// 				return
// 			}
// 		} else if s2 == "2B" {
// 			if no2(s3, s4, "2B") {
// 				fmt.Println("No")
// 				return
// 			} else if same(s3, s4) {
// 				fmt.Println("No")
// 				return
// 			}
// 		} else {
// 			if no2(s3, s4, "3B") {
// 				fmt.Println("No")
// 				return
// 			} else if same(s3, s4) {
// 				fmt.Println("No")
// 				return
// 			}
// 		}
// 	}
// 	fmt.Println("Yes")
// }

// func no3(t1 string, t2 string, t3 string, matcher string) bool {
// 	return t1 == matcher || t2 == matcher || t3 == matcher
// }

// func no2(t1 string, t2 string, matcher string) bool {
// 	return t1 == matcher || t2 == matcher
// }

// func same(t1 string, t2 string) bool {
// 	return t1 == t2
// }