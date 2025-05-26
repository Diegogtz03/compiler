package VM

// Saves current call locals and temps
type Memory struct {
	IntLocals   []int
	FloatLocals []float64
	IntTemps    []int
	FloatTemps  []float64
	BoolTemps   []bool
}

type GlobalMemory struct {
	GlobalInts   []int
	GlobalFloats []float64
}
