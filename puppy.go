package puppy

import (
	"github.com/MikeBazh/dog"
)

func Bark() string {
	return "Woof!!!!"
}
func BigBark() string {
	return dog.WhenGrownUp(Bark())
}
