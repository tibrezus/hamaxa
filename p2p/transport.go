package p2p

// Peer  is an iterface rappresenting an agent node
type Peer interface {

}


// Transport is anything that handles th communication
// between the agent nodes in the network
type Transport interface {
	ListenAndAccept() error
}