package lengthconv

import (
	"fmt"
)

type Incun float64
type Meter float64

func (i Incun) String() string  {
	return fmt.Sprintf("%g英寸",i)
}
func (m Meter) String() string {
	return fmt.Sprintf("%g米",m)

}