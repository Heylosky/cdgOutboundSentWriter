package main

import (
	"github.com/cdgProcessor/outboundSentWriter/db"
	"github.com/cdgProcessor/outboundSentWriter/logger"
	"github.com/cdgProcessor/outboundSentWriter/messageQ"
	"go.uber.org/zap"
)

func main() {
	logger.InitLogger("./logs/outboundSentWriter.log")
	zap.L().Info("Outbound sent record DB processor starting...")

	outSentChan := make(chan []byte)

	go messageQ.MQRead(outSentChan, "obSentRecordToDb", "obSentRecord")

	db.Writer(outSentChan)
}
