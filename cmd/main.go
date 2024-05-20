package main

import (
    "log"
    "os"

    "myproject/internal/confluence"
    "myproject/internal/drawio"
    "myproject/internal/ebpf"
    "myproject/internal/trace"
)

func main() {
    if len(os.Args) != 2 {
        log.Fatalf("Usage: %s /path/to/project", os.Args[0])
    }
    projectPath := os.Args[1]

    // Setup eBPF
    if err := ebpf.Setup(projectPath); err != nil {
        log.Fatalf("eBPF setup failed: %v", err)
    }
    defer ebpf.Cleanup()

    // Capture dynamic function calls
    dynamicCalls := trace.CaptureDynamic()

    // Perform static analysis
    staticCalls := trace.StaticAnalysis(projectPath)

    // Combine dynamic and static results
    combinedCalls := trace.CombineResults(dynamicCalls, staticCalls)

    // Generate Draw.io diagram XML
    drawioXML := drawio.GenerateXML(combinedCalls)

    // Create Confluence page with Draw.io diagram
    if err := confluence.CreatePage(drawioXML); err != nil {
        log.Fatalf("Failed to create Confluence page: %v", err)
    }

    log.Println("Confluence page created successfully")
}