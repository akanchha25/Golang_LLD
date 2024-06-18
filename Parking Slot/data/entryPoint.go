package data


// EntryPoint represents an entry point with a name and an open status.
type EntryPoint struct {
    Name  string
    Open bool
}

// NewEntryPoint creates a new EntryPoint with the provided name and open status.
func NewEntryPoint(name string, open bool) *EntryPoint {
    return &EntryPoint{
        Name:  name,
        Open: open,
    }
}

// GetName returns the name of the entry point.
func (e *EntryPoint) GetName() string {
    return e.Name
}

// IsOpen returns the open status of the entry point.
func (e *EntryPoint) IsOpen() bool {
    return e.Open
}
