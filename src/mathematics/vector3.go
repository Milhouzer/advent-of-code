package mathematics

import "math"

//		Y
//	 	^
//		|
//		|
//		|
//		|
//		|
//		+----------------> x
type Vector3 struct {
	X float64
	Y float64
	Z float64
}

func (v *Vector3) IsZero() bool {
	return v.X == 0 && v.Y == 0 && v.Z == 0
}

// Add other to a [Vector3]
func (v *Vector3) Add(other *Vector3) *Vector3 {
	return &Vector3{X: v.X + other.X, Y: v.Y + other.Y, Z: v.Z + other.Z}
}

// Subtract subtracts two vectors and returns the result.
func (v *Vector3) Subtract(other *Vector3) *Vector3 {
	return &Vector3{X: v.X - other.X, Y: v.Y - other.Y, Z: v.Z - other.Z}
}

func (v *Vector3) Mul(f float64) *Vector3 {
	return &Vector3{X: v.X * f, Y: v.Y * f, Z: v.Z * f}
}

// Magnitude returns the magnitude (length) of the vector.
func (v *Vector3) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normalize normalizes the vector to a unit vector.
func (v *Vector3) Normalize() *Vector3 {
	mag := v.Magnitude()
	if mag == 0 {
		return &Vector3{0, 0, 0}
	}
	return &Vector3{X: v.X / mag, Y: v.Y / mag, Z: v.Z / mag}
}

// RotateAroundAxis rotates the vector `v` by an angle `theta` (in radians) around an arbitrary axis `axis`.
// Rodrigues formula: https://en.wikipedia.org/wiki/Rodrigues%27_rotation_formula
func (v *Vector3) RotateAroundAxis(axis Vector3, theta float64) Vector3 {
	k := axis.Normalize()

	u1 := Vector3{
		X: v.X * math.Cos(theta),
		Y: v.Y * math.Cos(theta),
		Z: v.Z * math.Cos(theta),
	}

	crossProduct := v.Cross(k)
	u2 := Vector3{
		X: crossProduct.X * math.Sin(theta),
		Y: crossProduct.Y * math.Sin(theta),
		Z: crossProduct.Z * math.Sin(theta),
	}

	dotProduct := v.Dot(k)
	u3 := Vector3{
		X: k.X * dotProduct * (1 - math.Cos(theta)),
		Y: k.Y * dotProduct * (1 - math.Cos(theta)),
		Z: k.Z * dotProduct * (1 - math.Cos(theta)),
	}

	return Vector3{
		X: u1.X + u2.X + u3.X,
		Y: u1.Y + u2.Y + u3.Y,
		Z: u1.Z + u2.Z + u3.Z,
	}
}

// Fast clockwise rotation for a 2d vector contained in the z = 0 plane
func (v *Vector3) RotateRight() *Vector3 {
	return &Vector3{X: v.Y, Y: -v.X, Z: 0}
}

// Dot product between two vectors
func (v *Vector3) Dot(other *Vector3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

// Cross product between two vectors
func (u *Vector3) Cross(v *Vector3) *Vector3 {
	return &Vector3{
		X: u.Y*v.Z - u.Z*v.Y,
		Y: u.Z*v.X - u.X*v.Z,
		Z: u.X*v.Y - u.Y*v.X,
	}
}

func (u *Vector3) Equals(v *Vector3) bool {
	return u.X == v.X && u.Y == v.Y && u.Z == v.Z
}

// Reverse return the opposite of a vector
func (v *Vector3) Reverse() Vector3 {
	return Vector3{X: -v.X, Y: -v.Y, Z: -v.Z}
}
