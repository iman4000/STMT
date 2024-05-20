package ebpf

import (
    "log"
    "github.com/cilium/ebpf"
    "github.com/cilium/ebpf/link"
    "github.com/cilium/ebpf/rlimit"
)

// bpfObjects and other generated structs here...

// Setup initializes eBPF
func Setup(programPath string) error {
    // Allow the current process to lock memory for eBPF resources.
    if err := rlimit.RemoveMemlock(); err != nil {
        return err
    }

    // Load pre-compiled programs into the kernel.
    objs := bpfObjects{}
    if err := loadBpfObjects(&objs, nil); err != nil {
        return err
    }

    // Attach the program to the entry and return of functions
    err := attachUprobe(programPath, "trace_entry", objs.TraceEntry)
    if err != nil {
        return err
    }

    err = attachUretprobe(programPath, "trace_return", objs.TraceReturn)
    if err != nil {
        return err
    }

    log.Println("eBPF program loaded and attached")
    return nil
}

func attachUprobe(path, funcName string, prog *ebpf.Program) error {
    uprobe, err := link.Uprobe(path, funcName, prog, nil)
    if err != nil {
        return err
    }
    defer uprobe.Close()
    return nil
}

func attachUretprobe(path, funcName string, prog *ebpf.Program) error {
    uretprobe, err := link.Uretprobe(path, funcName, prog, nil)
    if err != nil {
        return err
    }
    defer uretprobe.Close()
    return nil
}

// Cleanup performs any necessary cleanup
func Cleanup() {
    // Close eBPF resources here
}