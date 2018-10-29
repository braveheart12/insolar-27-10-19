/*
 *    Copyright 2018 Insolar
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package hostnetwork

import (
	"fmt"
	"time"

	"github.com/insolar/insolar/configuration"
	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/log"
	"github.com/insolar/insolar/network"
	"github.com/insolar/insolar/network/transport"
	"github.com/insolar/insolar/network/transport/host"
	"github.com/insolar/insolar/network/transport/packet"
	"github.com/insolar/insolar/network/transport/packet/types"
	"github.com/insolar/insolar/network/transport/relay"
	"github.com/pkg/errors"
)

type hostTransport struct {
	transport transport.Transport
	origin    *host.Host
	handlers  map[types.PacketType]network.RequestHandler
}

type builder struct {
	sender *host.Host
	t      types.PacketType
	data   interface{}
}

func (b *builder) Type(packetType types.PacketType) network.RequestBuilder {
	b.t = packetType
	return b
}

func (b *builder) Data(data interface{}) network.RequestBuilder {
	b.data = data
	return b
}

func (b *builder) GetSender() core.RecordRef {
	return b.sender.NodeID
}

func (b *builder) GetSenderHost() *host.Host {
	return b.sender
}

func (b *builder) GetType() types.PacketType {
	return b.t
}

func (b *builder) GetData() interface{} {
	return b.data
}

func (b *builder) Build() network.Request {
	return b
}

type packetWrapper packet.Packet

func (p *packetWrapper) GetSender() core.RecordRef {
	return p.Sender.NodeID
}

func (p *packetWrapper) GetSenderHost() *host.Host {
	return p.Sender
}

func (p *packetWrapper) GetType() types.PacketType {
	return p.Type
}

func (p *packetWrapper) GetData() interface{} {
	return p.Data
}

type future struct {
	transport.Future
}

func (f future) Response() <-chan network.Response {
	in := transport.Future(f).Result()
	out := make(chan network.Response, cap(in))
	go func() {
		for packet := range in {
			out <- (*packetWrapper)(packet)
		}
	}()
	return out
}

func (f future) GetResponse(duration time.Duration) (network.Response, error) {
	select {
	case result, ok := <-f.Result():
		if !ok {
			return nil, transport.ErrChannelClosed
		}
		return (*packetWrapper)(result), nil
	case <-time.After(duration):
		f.Cancel()
		return nil, transport.ErrTimeout
	}
}

func (f future) GetRequest() network.Request {
	request := transport.Future(f).Request()
	return (*packetWrapper)(request)
}

// Listen start listening to network requests, should be started in goroutine.
func (h *hostTransport) Listen() error {
	go h.listen()
	return h.transport.Start()
}

func (h *hostTransport) listen() {
	for {
		select {
		case msg := <-h.transport.Packets():
			if msg == nil {
				log.Error("HostNetwork receiving channel is closed, disconnecting")
				h.Disconnect()
				break
			}
			if msg.Error != nil {
				log.Warnf("Received error response: %s", msg.Error.Error())
			}
			h.processMessage(msg)
		case <-h.transport.Stopped():
			h.transport.Close()
			return
		}
	}
}

func (h *hostTransport) processMessage(msg *packet.Packet) {
	handler, exist := h.handlers[msg.Type]
	if !exist {
		log.Errorf("No handler set for packet type %s from node %s",
			msg.Type.String(), msg.Sender.NodeID.String())
		return
	}
	response, err := handler((*packetWrapper)(msg))
	if err != nil {
		log.Errorf("Error handling request %s from node %s: %s",
			msg.Type.String(), msg.Sender.NodeID.String(), err)
		return
	}
	r := response.(*packetWrapper)
	h.transport.SendResponse(msg.RequestID, (*packet.Packet)(r))
}

// Disconnect stop listening to network requests.
func (h *hostTransport) Disconnect() error {
	h.transport.Stop()
	return nil
}

// PublicAddress returns public address that can be published for all nodes.
func (h *hostTransport) PublicAddress() string {
	return h.origin.Address.String()
}

// SendRequestPacket send request packet to a remote node.
func (h *hostTransport) SendRequestPacket(request network.Request, receiver *host.Host) (network.Future, error) {
	f, err := h.transport.SendRequest(h.buildRequest(request, receiver))
	if err != nil {
		return nil, err
	}
	return future{Future: f}, nil
}

func (h *hostTransport) buildRequest(request network.Request, receiver *host.Host) *packet.Packet {
	return packet.NewBuilder(h.origin).Receiver(receiver).
		Type(request.GetType()).Request(request.GetData()).Build()
}

// RegisterPacketHandler register a handler function to process incoming request packets of a specific type.
func (h *hostTransport) RegisterPacketHandler(t types.PacketType, handler network.RequestHandler) {
	_, exists := h.handlers[t]
	if exists {
		panic(fmt.Sprintf("multiple handlers for packet type %s are not supported!", t.String()))
	}
	h.handlers[t] = handler
}

// NewRequestBuilder create packet builder for an outgoing request with sender set to current node.
func (h *hostTransport) NewRequestBuilder() network.RequestBuilder {
	return &builder{sender: h.origin}
}

// BuildResponse create response to an incoming request with Data set to responseData.
func (h *hostTransport) BuildResponse(request network.Request, responseData interface{}) network.Response {
	sender := request.(*packetWrapper).Sender
	p := packet.NewBuilder(h.origin).Type(request.GetType()).Receiver(sender).Response(responseData).Build()
	return (*packetWrapper)(p)
}

func NewInternalTransport(conf configuration.Configuration) (InternalTransport, error) {
	tp, err := transport.NewTransport(conf.Host.Transport, relay.NewProxy())
	if err != nil {
		return nil, errors.Wrap(err, "error creating transport")
	}
	address, err := host.NewAddress(tp.PublicAddress())
	if err != nil {
		return nil, errors.Wrap(err, "error resolving address")
	}
	nodeID := core.NewRefFromBase58(conf.Node.Node.ID)
	if nodeID.Equal(core.RecordRef{}) {
		return nil, errors.Wrap(err, "error parsing NodeID from string (NodeID is zero)")
	}
	origin := &host.Host{NodeID: nodeID, Address: address}
	return &hostTransport{transport: tp, origin: origin, handlers: make(map[types.PacketType]network.RequestHandler)}, nil
}
