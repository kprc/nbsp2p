package service

import "sync"

var(
	wg *sync.WaitGroup
)


func StartNatService()  {
	if wg == nil{
		wg = &sync.WaitGroup{}
	}
	
}

func StopNatService()  {
	

	wg.Wait()
}