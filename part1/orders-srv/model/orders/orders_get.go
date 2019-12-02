package orders

import (
	"bac/db"
	"github.com/micro/go-micro/util/log"
	proto "microservice-in-micro/part1/orders-srv/proto/orders"
)

// GetOrder 获取订单
func (s *service) GetOrder(orderId int64) (order *proto.Order, err error) {
	order = &proto.Order{}

	// 获取数据库
	o := db.GetDB()
	// 查询
	err = o.QueryRow("SELECT id, user_id, book_id, inv_his_id, state FROM orders WHERE id = ?", orderId).Scan(
		&order.Id, &order.UserId, &order.BookId, &order.InvHistoryId, &order.State)
	if err != nil {
		log.Logf("[GetOrder] 查询数据失败，err：%s", err)
		return
	}

	return
}
