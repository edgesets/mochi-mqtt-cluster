package mqttcluster

import (
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
	"log/slog"
)

var _ mqtt.Hook = new(BrokerHook)

type BrokerHook struct {
	mqtt.HookBase
}

// ID returns the ID of the hook.
func (h *BrokerHook) ID() string {
	return "broker"
}

// Provides indicates which methods a hook provides. The default is none - this method
// should be overridden by the embedding hook.
func (h *BrokerHook) Provides(b byte) bool {
	return false
}

// Init performs any pre-start initializations for the hook, such as connecting to databases
// or opening files.
func (h *BrokerHook) Init(config any) error {
	return nil
}

// SetOpts is called by the server to propagate internal values and generally should
// not be called manually.
func (h *BrokerHook) SetOpts(l *slog.Logger, opts *mqtt.HookOptions) {
	h.Log = l
	h.Opts = opts
}

// Stop is called to gracefully shut down the hook.
func (h *BrokerHook) Stop() error {
	return nil
}

// OnStarted is called when the server starts.
func (h *BrokerHook) OnStarted() {}

// OnStopped is called when the server stops.
func (h *BrokerHook) OnStopped() {}

// OnSubscribe is called when a client subscribes to one or more filters.
func (h *BrokerHook) OnSubscribe(cl *mqtt.Client, pk packets.Packet) packets.Packet {
	return pk
}

// OnSubscribed is called when a client subscribes to one or more filters.
func (h *BrokerHook) OnSubscribed(cl *mqtt.Client, pk packets.Packet, reasonCodes []byte) {}

// OnUnsubscribe is called when a client unsubscribes from one or more filters.
func (h *BrokerHook) OnUnsubscribe(cl *mqtt.Client, pk packets.Packet) packets.Packet {
	return pk
}

// OnUnsubscribed is called when a client unsubscribes from one or more filters.
func (h *BrokerHook) OnUnsubscribed(cl *mqtt.Client, pk packets.Packet) {}

// OnPublish is called when a client publishes a message.
func (h *BrokerHook) OnPublish(cl *mqtt.Client, pk packets.Packet) (packets.Packet, error) {
	return pk, nil
}
