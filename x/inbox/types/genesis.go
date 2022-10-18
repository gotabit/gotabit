package types

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	for _, msg := range gs.Messages {
		if err := msg.Validate(); err != nil {
			return err
		}
	}
	return nil
}
