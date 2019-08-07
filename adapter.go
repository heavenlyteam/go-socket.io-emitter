package socketioemitter

import "fmt"

type Adapter struct {
	connector IConnector
}

type IAdapter interface {
	SetConnector(connector IConnector)
	GetConnector() IConnector
	EmitToID(socketID, event string, payload interface{})
	Broadcast(event string, payload interface{})
	sendMessage(m interface{}) error
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (a *Adapter) SetConnector(connector IConnector) {
	a.connector = connector
}

func (a *Adapter) GetConnector() IConnector {
	return a.connector
}

func (a *Adapter) EmitToID(socketID, event string, payload interface{}) {
	message := NewMessage(&EmitPayload{
		Event:       event,
		IsBroadcast: false,
		SocketID:    socketID,
		Data:        payload,
	})

	if err := a.sendMessage(message); err != nil {
		fmt.Print(err)
	}
}

func (a *Adapter) Broadcast(event string, payload interface{}) {
	message := NewMessage(&EmitPayload{
		Event:       event,
		IsBroadcast: true,
		Data:        payload,
	})

	if err := a.sendMessage(message); err != nil {
		fmt.Print(err)
	}
}

func (a *Adapter) sendMessage(m interface{}) (err error) {
	var encodedMessage []byte
	if encodedMessage, err = EncodeMessage(m); err != nil {
		return
	}

	if err = a.connector.PushEvent(encodedMessage); err != nil {
		return
	}

	return
}
