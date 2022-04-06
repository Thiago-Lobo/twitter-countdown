package logic

import (
	"time"
)

type fn func() bool

func ScheduleTask(theTask fn, thresholdTime time.Time, period time.Duration, runAtStart bool) {
	done := make(chan bool)

	taskClosure := func () {	
		if (theTask()) {
			done<-true
		}
	}
	
	ticker := time.NewTicker(period)

	go func() {
		if (runAtStart) {
			taskClosure()	
		}
		
		for {
			<-ticker.C
			taskClosure()
		}
	}()

	<-done
}
