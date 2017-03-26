package dnet

var (
	DNET_MAGIC_NUMBER [4]byte = [4]byte{68, 110, 101, 116}
)

const (
	/* unauthenticated */

	RPC_MESSAGE_CODE_BYE = iota
	RPC_MESSAGE_CODE_HELLO
	RPC_MESSAGE_CODE_TELL_YOUR_VERSION

	// client gives it's address, server checks it. success - authenticated
	RPC_MESSAGE_CODE_AUTHENTICATE

	/* client authenticated */

	// Client checks server address. success ==  (two-way) authentication
	// established
	RPC_MESSAGE_CODE_TELL_YOUR_ADDRESS
	RPC_MESSAGE_CODE_PROVE_YOUR_ADDRESS

	/* both (two-way) authenticated */

	RPC_MESSAGE_CODE_DO_YOU_SEE_ADDRESS
	RPC_MESSAGE_CODE_TELL_KNOWN_ADDRESSES
	RPC_MESSAGE_CODE_I_SEE_ADDRESS
	RPC_MESSAGE_CODE_ESTABLISH_PROXY_CONNECTION_TO_ADDRESS
	RPC_MESSAGE_CODE_LIST_YOUR_SERVICES
	RPC_MESSAGE_CODE_CONNECT_TO_SERVICE
)

/*

	1 - server responting. by establishing connection to node, which asked
	    by using UDP broadcast message

	2 - request termination. for cases, when client already got enough data by
			simply connecting to node

	3 - Hello must necessaryly be made, as it should be the prove of what
			client is really a Dnet client and knows what it talking.

			Any errors in Client hello, should be treated by nodes with
			emmidiate connection closing.

*/

type RPCCommandCode uint16

type C2S_Hello_Step0_Request struct {
}
