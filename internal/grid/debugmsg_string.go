// Code generated by "stringer -type=debugMsg debug.go"; DO NOT EDIT.

package grid

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[debugShutdown-0]
	_ = x[debugKillInbound-1]
	_ = x[debugKillOutbound-2]
	_ = x[debugWaitForExit-3]
	_ = x[debugSetConnPingDuration-4]
	_ = x[debugSetClientPingDuration-5]
	_ = x[debugAddToDeadline-6]
}

const _debugMsg_name = "debugShutdowndebugKillInbounddebugKillOutbounddebugWaitForExitdebugSetConnPingDurationdebugSetClientPingDurationdebugAddToDeadline"

var _debugMsg_index = [...]uint8{0, 13, 29, 46, 62, 86, 112, 130}

func (i debugMsg) String() string {
	if i < 0 || i >= debugMsg(len(_debugMsg_index)-1) {
		return "debugMsg(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _debugMsg_name[_debugMsg_index[i]:_debugMsg_index[i+1]]
}
