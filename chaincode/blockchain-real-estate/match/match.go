package match

import (
//"fmt"
// "github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
)

type Res struct {
	//num     int //用作计数
	process []string
}

// Res := &lib.Res
func Match(arr []string, m map[string][]string) /*(key string,val []string)*/ (r map[string][]string) { //arr 是 task , m 是 resource
	//key := &
	resource := make(map[string][]string)

	for key, val := range m { //for i := 0; i < len(arr); i++ { //for i, v := range arr {
		j := 0
		num := 0
		t := make([]string, len(arr))

		for i := 0; i < len(arr); i++ {
			//v := &val
			for n := 0; n < len(val /*.process*/); n++ {
				//strings.Compare(val.process[n], arr[i])
				if val /*.process*/ [n] == arr[i] {
					num++
					t[j] = val /*.process*/ [n]
					j++

				}
			}
		}
		if len(arr) == num {
			//fmt.Printf("%v 企业能接受 task: %v\n", key, arr)
			resource[key] = arr
		} else if num > 0 && num < len(arr) {
			//fmt.Printf("%v 企业部分接受 task:%v\n", key, t)
			resource[key] = t
		} else {
			//fmt.Printf("%v 企业不能加工 task\n ", key)
			resource[key] = nil
		}
	}
	return resource
}

// func main() {

// 	res1 := Res{[]string{"De", "Dr", "Bo", "M", "W"}}

// 	res2 := Res{ []string{"Dr", "Bo", "W"}}
// 	res3 := Res{ []string{"Bo", "W"}}
// 	res4 := Res{ []string{"De", "M", "Dr"}}
// 	res5 := Res{ []string{"Dr"}}

// 	resource := make(map[string]Res)

// 	resource["E1"] = res1
// 	resource["E2"] = res2
// 	resource["E3"] = res3
// 	resource["E4"] = res4
// 	resource["E5"] = res5
// 	//var task1 = []string{"Bo", "Dr"}
// 	//var task2 = []string{"De", "M", "Dr"}
// 	var task3 = []string{"De", "Bo", "M", "W"}
// 	Match(task3, resource)
// }
