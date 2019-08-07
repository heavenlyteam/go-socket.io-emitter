package socketioemitter

type Emitter struct {
	adapter IAdapter
}

type IEmitter interface {
	SetAdapter(adapter IAdapter)
	GetAdapter() IAdapter
	EmitToID(socketID, event string, payload interface{})
	Broadcast(event string, payload interface{})
}

func NewEmitter(a IAdapter) (e *Emitter) {
	emitter := Emitter{}
	emitter.SetAdapter(a)

	// pointer binding
	e = &emitter
	return
}

func (e *Emitter) SetAdapter(a IAdapter) {
	e.adapter = a
}

func (e *Emitter) GetAdapter() IAdapter {
	return e.adapter
}

func (e *Emitter) EmitToID(socketID, event string, payload interface{}) {
	e.adapter.EmitToID(socketID, event, payload)
}

func (e *Emitter) Broadcast(event string, payload interface{}) {
	e.adapter.Broadcast(event, payload)
}
