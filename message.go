package socketioemitter

type EmitPayload struct {
	Event       string      `json:"event"`
	IsBroadcast bool        `json:"is_broadcast"`
	SocketID    string      `json:"socket_id"`
	Data        interface{} `json:"data"`
}

type Message struct {
	UID    string `json:"uid"`
	Packet Packet `json:"packet"`
	Opts   Opts   `json:"opts"`
}

type Packet struct {
	Type      int         `json:"type"`
	Data      interface{} `json:"data"`
	Namespace string      `json:"namespace"`
}

type Opts struct {
	Rooms []string `json:"rooms"`
	Flags []string `json:"flags"`
}

func NewPacket(payload interface{}, nsp string) (p Packet) {
	p.Data = payload
	p.Type = 2
	p.Namespace = nsp
	return
}

func NewOpts(rooms, flags []string) (o Opts) {
	o.Rooms = rooms
	o.Flags = flags
	return
}

func NewMessage(payload interface{}) (m Message) {
	packet := NewPacket(payload, "/")
	opts := NewOpts([]string{}, []string{})
	m.Opts = opts
	m.Packet = packet
	m.UID = "emitter"
	return
}
