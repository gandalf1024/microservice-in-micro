package payment

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/micro/go-micro/util/log"
	proto "microservice-in-micro/part1/payment-srv/proto/payment"
)

// sendPayDoneEvt 发送支付事件
func (s *service) sendPayDoneEvt(orderId int64, state int32) {
	// 构建事件
	ev := &proto.PayEvent{
		Id:       uuid.New().String(),
		SentTime: time.Now().Unix(),
		OrderId:  orderId,
		State:    state,
	}

	log.Logf("[sendPayDoneEvt] 发送支付事件，%+v\n", ev)

	// 广播
	if err := payPublisher.Publish(context.Background(), ev); err != nil {
		log.Logf("[sendPayDoneEvt] 异常: %v", err)
	}
}
