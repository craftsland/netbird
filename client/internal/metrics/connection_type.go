package metrics

// ConnectionType represents how a peer connection carries traffic.
//
// The values map one-to-one onto conntype.ConnPriority. They deliberately do not reuse the
// previous "ice" value: that one conflated direct peer-to-peer, ICE-over-TURN and the
// transient "no connection established yet" state, so samples tagged "ice" cannot be
// compared with the ones emitted now.
type ConnectionType string

const (
	// ConnectionTypeICEP2P is a direct peer-to-peer connection negotiated over ICE.
	ConnectionTypeICEP2P ConnectionType = "ice_p2p"

	// ConnectionTypeICETurn is an ICE connection traversing a TURN server. The traffic is
	// relayed even though ICE established it, which is why Conn.isRelayed counts it as
	// relayed rather than direct.
	ConnectionTypeICETurn ConnectionType = "ice_turn"

	// ConnectionTypeRelay is a connection carried by a NetBird relay.
	ConnectionTypeRelay ConnectionType = "relay"

	// ConnectionTypeUnknown is recorded when the connection priority is unset at sampling
	// time. Reporting it explicitly keeps those samples out of the peer-to-peer bucket
	// instead of silently inflating it.
	ConnectionTypeUnknown ConnectionType = "unknown"
)

// String returns the string representation of the connection type
func (c ConnectionType) String() string {
	return string(c)
}
