package vec

const EPSILON = 1.0e-5

type Vector []float64

type IntVector []int

func lenMap(vecs ...Vector)

func mapVec(f func(...float64), vecs ...Vector) Vector {

}

func VecDot(a, b Vector) float64 {
	shortest := func() int {
		aLen := len(a)
		bLen := len(b)
		if aLen < bLen {
			return aLen
		} else {
			return bLen
		}
	}()

}
