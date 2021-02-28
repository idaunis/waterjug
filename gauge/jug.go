package gauge

type JugState int32

const (
	Empty JugState = 0
	Half  JugState = 1
	Full  JugState = 2
)

type Jug struct {
	Size int
	Max  int
}

func NewJug(max int) *Jug {
	return &Jug{0, max}
}

func (j *Jug) Fill() {
	j.Size = j.Max
}

func (j *Jug) Empty() {
	j.Size = 0
}

func (j *Jug) Transfer(d *Jug) int {
	amount := min(j.Size, d.Max-d.Size)
	j.Size -= amount
	d.Size += amount
	return amount
}

func (j *Jug) IsEmpty() bool {
	return j.Size == 0
}

func (j *Jug) IsFilled() bool {
	return j.Size == j.Max
}

func (j *Jug) GetState() JugState {
	if j.IsEmpty() {
		return Empty
	} else if j.IsFilled() {
		return Full
	}
	return Half
}
