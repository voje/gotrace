package main

import __log "github.com/voje/gotrace/log"

func someBytes(b []byte) {
	__traceID := __log.ID()
	__log.L.Printf("[%d] someBytes(%s)\n", __traceID, __log.Format(b))

	return
}

func main() {
	__traceID := __log.ID()
	__log.L.Printf("[%d] main(%s)\n", __traceID, __log.Format())

	someBytes([]byte{0x61, 0x62, 0x63})
}

var _ = __log.Setup("stderr", "", 1024)
