package dnet

const (
	DNET_MAGIC_NUMBER = [4]byte{68, 110, 101, 116}
)

const (
	/* unauthenticated */

	RPC_MESSAGE_CODE_HELLO             = iota /* 3 */
	RPC_MESSAGE_CODE_BYE               = iota /* 2  */
	RPC_MESSAGE_CODE_TELL_YOUR_VERSION = iota

	// client gives it's address, server checks it. success - authenticated
	RPC_MESSAGE_CODE_AUTHENTICATE = iota

	/* client authenticated */

	// Client checks server address. success ==  (two-way) authentication
	// established
	RPC_MESSAGE_CODE_TELL_YOUR_ADDRESS  = iota
	RPC_MESSAGE_CODE_PROVE_YOUR_ADDRESS = iota

	/* both (two-way) authenticated */

	RPC_MESSAGE_CODE_DO_YOU_SEE_ADDRESS                    = iota
	RPC_MESSAGE_CODE_TELL_KNOWN_ADDRESSES                  = iota
	RPC_MESSAGE_CODE_I_SEE_ADDRESS                         = iota /* 1 */
	RPC_MESSAGE_CODE_ESTABLISH_PROXY_CONNECTION_TO_ADDRESS = iota
	RPC_MESSAGE_CODE_LIST_YOUR_SERVICES                    = iota
	RPC_MESSAGE_CODE_CONNECT_TO_SERVICE                    = iota
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
