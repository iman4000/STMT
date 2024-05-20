package trace

func CombineResults(dynamicCalls, staticCalls []FunctionCall) []FunctionCall {
    combined := make(map[string]FunctionCall)

    // Add dynamic calls to the combined map
    for _, call := range dynamicCalls {
        key := call.From + "->" + call.To
        combined[key] = call
    }

    // Add static calls to the combined map if they don't already exist
    for _, call := range staticCalls {
        key := call.From + "->" + call.To
        if _, exists := combined[key]; !exists {
            combined[key] = call
        }
    }

    var result []FunctionCall
    for _, call := range combined {
        result = append(result, call)
    }
    return result
}