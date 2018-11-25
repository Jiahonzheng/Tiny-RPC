package tiny

import (
	"encoding/binary"
	"io"
	"net"
)

// Transport struct
type Transport struct {
	conn net.Conn
}

// NewTransport creates a transport
func NewTransport(conn net.Conn) *Transport {
	return &Transport{conn}
}

// Send data
func (t *Transport) Send(req Data) error {
	b, err := encode(req) // Encode req into bytes
	if err != nil {
		return err
	}
	buf := make([]byte, 4+len(b))
	binary.BigEndian.PutUint32(buf[:4], uint32(len(b))) // Set Header field
	copy(buf[4:], b)                                    // Set Data field
	_, err = t.conn.Write(buf)
	return err
}

// Receive data
func (t *Transport) Receive() (Data, error) {
	header := make([]byte, 4)
	_, err := io.ReadFull(t.conn, header)
	if err != nil {
		return Data{}, err
	}
	dataLen := binary.BigEndian.Uint32(header) // Read Header filed
	data := make([]byte, dataLen)              // Read Data Field
	_, err = io.ReadFull(t.conn, data)
	if err != nil {
		return Data{}, err
	}
	rsp, err := decode(data) // Decode rsp from bytes
	return rsp, err
}
