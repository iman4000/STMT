package trace

// FunctionCall represents a function call
type FunctionCall struct {
    From     string
    To       string
    Duration uint64
}

// CaptureDynamic captures function calls dynamically
func CaptureDynamic() []FunctionCall {
    // Implement dynamic tracing logic here
    return []FunctionCall{}
}