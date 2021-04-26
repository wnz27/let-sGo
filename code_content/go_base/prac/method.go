/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/4/26 4:49 下午
* Description:
 */
package prac

import (
	"fmt"
	"math"
	"net/url"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}


type Path []Point

func (path Path) Distance() float64 {
	pathLength := 0.0
	for i, _ := range path {
		if i > 0 {
			pathLength += path[i - 1].Distance(path[i])
		}
	}
	return pathLength
}

//type Values map[string] []string
//
//func (v Values) Get(key string) string {
//	if vs := v[key]; len(vs) == 0 {
//		return vs[0]
//	}
//	return ""
//}
//
//func (v Values) Add(key, value string) {
//	v[key] = append(v[key], value)
//}

func TMethod() {
	p := Point{1, 2}
	q := Point{3, 4}
	fmt.Println(p.Distance(q))

	paths := Path{
		{1, 2},
		{2, 5},
		{5, 6},
	}
	fmt.Println(paths.Distance())


	r := &Point{1, 2}
	r.ScaleBy(2)

	p1 := Point{1, 2}
	p1.ScaleBy(2)


	p2 := Point{1, 2}
	(&p2).ScaleBy(2)

	fmt.Println(r, p1, p2)

	m := url.Values{"lang": {"en"}}

	m = nil
	fmt.Println(m.Get("item"))
	m.Add("item", "123")


}
