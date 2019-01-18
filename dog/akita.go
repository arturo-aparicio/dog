package dog

import (
	"fmt"
        "github.com/jfrog-aparicio/food"
)

func AkitaPackageName () {
	fmt.Println ("Package Name: akita")
	BeaglePackageName()
	ColliePackageName()
	food.AlpoPackageName()
}
