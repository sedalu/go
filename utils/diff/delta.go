package diff
import "fmt"

type Delta map[byte][]interface{}

func (p *Delta) String() (s string) {
	for side, delta := range *p {
		for _, v := range delta {
			switch side {
			case LEFT:
				s += fmt.Sprintf("-")
			case RIGHT:
				s += fmt.Sprintf("+")
			default:
				s += fmt.Sprintf(" ")
			}

			s += fmt.Sprintf(" %v\n", v)
		}
	}

	return
}