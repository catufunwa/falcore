// +build windows 

package falcore

import (
	"net"
)

// only valid on non-windows
func (srv *Server) setupNonBlockingListener(err error, l *net.TCPListener) error {
	return nil
}

func (srv *Server) setNoDelay(c net.Conn, noDelay bool) {
	return false
}
