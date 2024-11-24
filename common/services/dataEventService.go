package services

import "sync"

type EventModel struct {
	Event string //事件类型
	Data  any    // 数据
}

type EventChan chan EventModel

type EventBus struct {
	mu         sync.Mutex
	subscriber map[string][]EventChan
}

var EB = &EventBus{
	subscriber: map[string][]EventChan{},
}

func (eb *EventBus) Subscribe(topic string) EventChan {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	ch := make(EventChan)
	eb.subscriber[topic] = append(eb.subscriber[topic], ch)
	return ch
}

func (eb *EventBus) Publish(topic string, data EventModel) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	subscribers := eb.subscriber[topic]
	go func() {
		for _, subscriber := range subscribers {
			subscriber <- data
		}
	}()
}

func (eb *EventBus) UnSubscribe(topic string, ch EventChan) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	if subscriber, ok := eb.subscriber[topic]; ok {
		for i, eventChan := range subscriber {
			if eventChan == ch {
				eb.subscriber[topic] = append(eb.subscriber[topic][:i], eb.subscriber[topic][i+1:]...)
				close(ch)
				for range ch {

				}
				return
			}
		}
	}

}
