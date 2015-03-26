package euler

type SolutionFunc func()(uint64,error)

type solutionType struct {
    Name string
    Func SolutionFunc
}

var solutions []solutionType

func RegisterSolution(name string, solution SolutionFunc) {
    solutions = append(solutions, solutionType{name, solution})
}

func init() {
    solutions = make([]solutionType, 0, 100)
}
