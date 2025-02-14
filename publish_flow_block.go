package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func (publisher *Publisher) startNotifyFlowHandler() {
	notifyFlowChan := publisher.chanManager.NotifyFlowSafe(make(chan bool))
	publisher.disablePublishDueToFlowMux.Lock()
	publisher.disablePublishDueToFlow = false
	publisher.disablePublishDueToFlowMux.Unlock()

	for ok := range notifyFlowChan {
		publisher.disablePublishDueToFlowMux.Lock()
		if ok {
			publisher.options.Logger.Warnf("pausing publishing due to flow request from server")
			publisher.disablePublishDueToFlow = true
		} else {
			publisher.disablePublishDueToFlow = false
			publisher.options.Logger.Warnf("resuming publishing due to flow request from server")
		}
		publisher.disablePublishDueToFlowMux.Unlock()
	}
}

func (publisher *Publisher) startNotifyBlockedHandler() {
	blockings := publisher.connManager.NotifyBlockedSafe(make(chan amqp.Blocking))
	//publisher.disablePublishDueToBlockedMux.Lock()
	//publisher.disablePublishDueToBlocked = false
	//publisher.disablePublishDueToBlockedMux.Unlock()

	publisher.disablePublishDueToBlockedMux.Lock()
	defer publisher.disablePublishDueToBlockedMux.Unlock() // 确保锁在函数结束时被解锁
	publisher.disablePublishDueToBlocked = false
	//for b := range blockings {
	//	publisher.disablePublishDueToBlockedMux.Lock()
	//	if b.Active {
	//		publisher.options.Logger.Warnf("pausing publishing due to TCP blocking from server")
	//		publisher.disablePublishDueToBlocked = true
	//	} else {
	//		publisher.disablePublishDueToBlocked = false
	//		publisher.options.Logger.Warnf("resuming publishing due to TCP blocking from server")
	//	}
	//	publisher.disablePublishDueToBlockedMux.Unlock()
	//}
	for {
		select {
		case b, ok := <-blockings:
			if !ok {
				// 通道关闭，退出 goroutine
				return
			}
			publisher.disablePublishDueToBlockedMux.Lock()
			defer publisher.disablePublishDueToBlockedMux.Unlock() // 确保每次循环中锁都被解锁
			if b.Active {
				publisher.options.Logger.Warnf("pausing publishing due to TCP blocking from server")
				publisher.disablePublishDueToBlocked = true
			} else {
				publisher.disablePublishDueToBlocked = false
				publisher.options.Logger.Warnf("resuming publishing due to TCP blocking from server")
			}
		case <-publisher.stopCh:
			// 接收到停止信号，退出 goroutine
			return
		}
	}
}
