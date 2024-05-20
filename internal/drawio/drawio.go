package drawio

import (
    "fmt"
    "strings"

    "myproject/internal/trace"
)

// GenerateXML generates the Draw.io XML
func GenerateXML(calls []trace.FunctionCall) string {
    var sb strings.Builder

    sb.WriteString(`<mxGraphModel><root>`)
    sb.WriteString(`<mxCell id="0"/><mxCell id="1" parent="0"/>`)

    // Add nodes
    nodeIDs := make(map[string]string)
    idCounter := 2
    for _, call := range calls {
        if _, exists := nodeIDs[call.From]; !exists {
            nodeIDs[call.From] = fmt.Sprintf("%d", idCounter)
            sb.WriteString(fmt.Sprintf(`<mxCell id="%d" value="%s" style="ellipse" vertex="1" parent="1"/>`, idCounter, call.From))
            idCounter++
        }
        if _, exists := nodeIDs[call.To]; !exists {
            nodeIDs[call.To] = fmt.Sprintf("%d", idCounter)
            sb.WriteString(fmt.Sprintf(`<mxCell id="%d" value="%s" style="ellipse" vertex="1" parent="1"/>`, idCounter, call.To))
            idCounter++
        }
    }

    // Add edges
    for _, call := range calls {
        fromID := nodeIDs[call.From]
        toID := nodeIDs[call.To]
        sb.WriteString(fmt.Sprintf(`<mxCell id="%d" edge="1" source="%s" target="%s" parent="1"/>`, idCounter, fromID, toID))
        idCounter++
    }

    sb.WriteString(`</root></mxGraphModel>`)

    return sb.String()
}