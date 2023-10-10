package main

import (
	"github.com/dtm-labs/client/dtmcli"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

const (
	dtmSrvAddr = "http://192.168.1.91:36789/api/dtmsvr"
	t1Addr     = "http://192.168.1.91:8082"
	t2Addr     = "http://192.168.1.91:8083"
)

func main() {
	req := &gin.H{"amount": 30}

	saga := dtmcli.NewSaga(dtmSrvAddr, uuid.NewString()).
		Add(t1Addr+"/api/bank1/TransOut", t1Addr+"/api/bank1/TransOutCompensate", req).
		Add(t2Addr+"/api/bank2/TransIn", t2Addr+"/api/bank2/TransInCompensate", req)

	err := saga.Submit()

	if err != nil {
		log.Panic().Msg(err.Error())
		// panic(err)
	}

	log.Info().Msgf("saga gid: %s", saga.Gid)
}
