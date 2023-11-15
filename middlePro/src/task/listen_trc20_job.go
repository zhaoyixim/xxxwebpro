package task

import (
	"sync"

	"github.com/pengk/summer/model/data"
	"github.com/pengk/summer/model/service"
	"github.com/pengk/summer/util/log"
)

type ListenTrc20Job struct {
}

var gListenTrc20JobLock sync.Mutex

func (r ListenTrc20Job) Run() {
	gListenTrc20JobLock.Lock()
	defer gListenTrc20JobLock.Unlock()
	//收款地址
	walletAddress, err := data.GetAvailableWalletAddress()
	if err != nil {
		log.Sugar.Error(err)
		return
	}
	if len(walletAddress) <= 0 {
		return
	}
	var wg sync.WaitGroup
	for _, address := range walletAddress {
		wg.Add(1)
		go service.NewTrc20CallBack(address.Token, &wg)
	}
	wg.Wait()
}
