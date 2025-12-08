package point

import "math"

type Point3D struct {
	X int
	Y int
	Z int
}

func NewPoint3D(x, y, z int) Point3D {
	return Point3D{x, y, z}
}

func (p *Point3D) DistanceVector(p2 Point3D) Point3D {
	x := p2.X - p.X
	y := p2.Y - p.Y
	z := p2.Z - p.Z
	return Point3D{x, y, z}
}

func (p *Point3D) Distance(p2 Point3D) float64 {
	distance := p.DistanceVector(p2)
	return math.Sqrt(math.Pow(float64(distance.X), 2) + math.Pow(float64(distance.Y), 2) + math.Pow(float64(distance.Z), 2))
}
