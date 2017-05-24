package strain

const testVersion = 1

type Ints []int
type Lists [][]int
type Strings []string

func (collection Ints) Keep(f func(int) bool) (output Ints) {
	for _, element := range collection {
		if f(element) {
			output = append(output, element)
		}
	}
	return
}
func (collection Ints) Discard(f func(int) bool) (output Ints) {
	for _, element := range collection {
		if !f(element) {
			output = append(output, element)
		}
	}
	return
}
func (collection Lists) Keep(f func([]int) bool) (output Lists) {
	for _, element := range collection {
		if f(element) {
			output = append(output, element)
		}
	}
	return
}
func (collection Strings) Keep(f func(string) bool) (output Strings) {
	for _, element := range collection {
		if f(element) {
			output = append(output, element)
		}
	}
	return
}
