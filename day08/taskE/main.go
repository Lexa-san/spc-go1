package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(newX, newY float64) *Point {
	return &Point{newX, newY}
}

//Return X coord of Point
func (p *Point) GetX() float64 {
	return p.x
}

//Return Y coord of Point
func (p *Point) GetY() float64 {
	return p.y
}

//Set new X coord for Existing point
func (p *Point) SetX(newX float64) {
	p.x = newX
}

//Set new Y coord for Existing point
func (p *Point) SetY(newY float64) {
	p.y = newY
}

//Retunrt string info about point in format Point{X: value, Y: value}
func (p *Point) Stringify() string {
	return fmt.Sprintf("Point{X: %f, Y: %f}", p.x, p.y)
}

type Line struct {
	p1 Point
	p2 Point
}

func NewLine(newP1, newP2 Point) *Line {
	return &Line{newP1, newP2}
}

//Return p1 of line
func (l *Line) GetP1() Point {
	return l.p1
}

//Return p2 of line
func (l *Line) GetP2() Point {
	return l.p2
}

//Set new p1 for Existing Line
func (l *Line) SetP1(newP1 Point) {
	l.p1 = newP1
}

//Set new p2 for Existing Line
func (l *Line) SetP2(newP2 Point) {
	l.p2 = newP2
}

//Calc length of segment p1p2
func (l *Line) Length() float64 {
	return math.Sqrt(math.Pow(l.p2.GetX()-l.p1.GetX(), 2) + math.Pow(l.p2.GetY()-l.p1.GetY(), 2))
}

//Method return True if p3 belongs to p1p2 segment
//                       False otherwise
// (x-x1)/(x2-x1) == (y-y1)/(y2-y1)
func (l *Line) IsOnSegment(p3 Point) bool {
	partLeft := (p3.GetX() - l.p1.GetX()) / (l.p2.GetX() - l.p1.GetX())
	partRight := (p3.GetY() - l.p1.GetY()) / (l.p2.GetY() - l.p1.GetY())
	return partLeft == partRight
}

//Retunrt string info about line in format Line{P1: value, P2: value}
func (l *Line) Stringify() string {
	return fmt.Sprintf("Line{P1: %s, P2: %s}\n", l.p1.Stringify(), l.p2.Stringify())
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	pointsRaw := strings.Split(scanner.Text(), " ")
	if len(pointsRaw) != 4 {
		return
	}

	scanner.Scan()
	point3Raw := strings.Split(scanner.Text(), " ")
	if len(point3Raw) != 2 {
		return
	}

	var x, y float64
	x, _ = strconv.ParseFloat(pointsRaw[0], 64)
	y, _ = strconv.ParseFloat(pointsRaw[1], 64)
	point1 := NewPoint(x, y)
	//fmt.Println(point1.Stringify())

	x, _ = strconv.ParseFloat(pointsRaw[2], 64)
	y, _ = strconv.ParseFloat(pointsRaw[3], 64)
	point2 := NewPoint(x, y)
	//fmt.Println(point2.Stringify())

	line := NewLine(*point1, *point2)
	//fmt.Printf(line.Stringify())
	//fmt.Printf("Length of line: %f\n", line.Length())

	x, _ = strconv.ParseFloat(point3Raw[0], 64)
	y, _ = strconv.ParseFloat(point3Raw[1], 64)
	point3 := NewPoint(x, y)
	//fmt.Println(point3.Stringify())

	if line.IsOnSegment(*point3) {
		fmt.Printf("YES")
	} else {
		fmt.Printf("NO")
	}

}
