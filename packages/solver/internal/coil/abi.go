package coil

type ABIFunction struct {
	Name            string         `json:"name,omitempty"`
	Inputs          []ABIParameter `json:"inputs"`
	Outputs         []ABIParameter `json:"outputs"`
	StateMutability string         `json:"stateMutability"`
	Type            string         `json:"type"`
}

type ABIParameter struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	InternalType string `json:"internalType,omitempty"`
	Indexed      bool   `json:"indexed,omitempty"`
}
