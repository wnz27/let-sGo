/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/5/1 00:56 5月
 **/
package main

import "fmt"

// Definition for Employee.
type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

// 有误 没有深度优先
func getImportance(employees []*Employee, id int) int {

	userImportant := make(map[int]int)
	for _, e := range employees {
		userImportant[e.Id] = e.Importance
	}
	for _, e := range employees {
		if e.Id == id {
			important := e.Importance
			if len(e.Subordinates) == 0 {
				return important
			}
			for _, sub_id := range e.Subordinates {
				imp, ok := userImportant[sub_id]; if ok {
					important += imp
				} else {
					panic("user not found")
				}
			}
			fmt.Println(e.Subordinates)
			fmt.Println(important)
			return important
		}
	}
	return 0
}

func getImportance2(employees []*Employee, id int) (total int) {
	mp := map[int]*Employee{}
	for _, employee := range employees {
		mp[employee.Id] = employee
	}

	var dfs func(int)
	dfs = func(id int) {
		employee := mp[id]
		total += employee.Importance
		for _, subId := range employee.Subordinates {
			dfs(subId)
		}
	}
	dfs(id)
	return
}
