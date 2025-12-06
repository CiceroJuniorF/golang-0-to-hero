package tasks

import (
	"fmt"
	"time"
)

func Time() {
	const LAYOUT = "02/01/2006 15:00:00"
	fmt.Printf("Now is %s \n", time.Now().Format(LAYOUT))
}
