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

}
