package trace

import (
    "encoding/json"
    "log"
    "os/exec"
)

type SemgrepResult struct {
    Results []struct {
        CheckID   string `json:"check_id"`
        Path      string `json:"path"`
        StartLine int    `json:"start"`
        EndLine   int    `json:"end"`
        Extra     struct {
            Message  string `json:"message"`
            Metadata struct {
                Metadata map[string]string `json:"metadata"`
            } `json:"metadata"`
        } `json:"extra"`
    } `json:"results"`
}

func StaticAnalysis(projectPath string) []FunctionCall {
    var functionCalls []FunctionCall

    // Run Semgrep
    cmd := exec.Command("semgrep", "--config", "rules.yml", "--json", projectPath)
    output, err := cmd.Output()
    if err != nil {
        log.Fatalf("Static analysis failed: %v", err)
    }

    // Parse Semgrep output
    var result SemgrepResult
    if err := json.Unmarshal(output, &result); err != nil {
        log.Fatalf("Failed to parse Semgrep output: %v", err)
    }

    // Extract function calls from Semgrep results
    for _, res := range result.Results {
        // Simplified extraction logic
        functionCalls = append(functionCalls, FunctionCall{
            From:     res.Path,
            To:       res.Extra.Metadata["function"],
            Duration: 0, // Static analysis doesn't provide duration
        })
    }

    return functionCalls
}