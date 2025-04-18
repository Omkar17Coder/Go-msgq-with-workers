package golanglearning

import (
	"fmt"
	"sync"

	"time"
)


func FanIn(done chan struct{},channels ...chan string) chan string{
	mergedChannel:=make(chan string)
	var wg sync.WaitGroup
	for _,chan:=range channels{

	}
	
}



func SimulateFanInFanOut(){
	chan1:=make(chan string)
	chan2:=make(chan string)
	chan3:=make(chan string)
	done:=make(chan struct{})

	mergedChannel:=FanIn(done,chan1,chan2,chan3)

}