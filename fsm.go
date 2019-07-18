package fsm

// StateFunc - state determined by the transition function to the new state.
// The state should determine whether it goes into some other state,
// including relying on incoming events through the channel (then the machine should be started as a goroutine).
// The target state must be initialized before returning.
type StateFunc func() (StateFunc, error)

// StateMachine enters the start state and then synchronously waits for receiving a new state,
// which it then also enters, or nil if the machine needs to complete its work.
// If the state is dependent on external events, it is better to deliver them through the channel;
// in this case, it is recommended to run StateMachine in a separate goroutine.
// StateMachine should only be called once, because next transitions are determined by states.
func StateMachine(start StateFunc) error {
	for currState := start; currState != nil; {
		var err error
		currState, err = currState()
		if err != nil {
			return err
		}
	}
	return nil
}
