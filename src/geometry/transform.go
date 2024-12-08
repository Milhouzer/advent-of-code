package geometry

type Transform struct {
	Position Vector3
	Rotation Vector3
}

func (t *Transform) Translate(v *Vector3) {
	t.Position = *t.Position.Add(v)
}

// Fast clockwise rotation
func (t *Transform) RotateRight() {
	t.Rotation = *t.Rotation.RotateRight()
}
