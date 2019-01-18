package dog

import (
	"fmt"
	"github.com/jfrog-aparicio/food"
)

func BeaglePackageName () {
	fmt.Println ("Package Name: beagle")
	ColliePackageName()
	food.PurinaPackageName()
}
