# gopulse (Learning Reference)

A CLI tool that takes a list of URLs and checks their HTTP status. I am building this specifically to practice Go's concurrency primitives rather than relying on standard thread pools.

## Core Concepts to Remember

*   **Goroutines:** A goroutine is simply a function running concurrently alongside other code[span_0](start_span)[span_0](end_span). Go's runtime automatically multiplexes goroutines onto OS threads, managing their scheduling so I do not have to[span_1](start_span)[span_1](end_span). They are extremely cheap to create[span_2](start_span)[span_2](end_span).
*   **Channels:** Used to pass data safely between goroutines. They are inherently composable, making it easier to coordinate inputs from multiple subsystems[span_3](start_span)[span_3](end_span). 
*   **Context:** The `context` package is used to provide an API for canceling branches of a call-graph and enforcing deadlines[span_4](start_span)[span_4](end_span). 
*   **Compilation:** The `go build` command compiles the Go code into a final executable program that can be run directly[span_5](start_span)[span_5](end_span).

## Project Milestones

### 01_sequential
*   **Goal:** Basic HTTP fetching and error handling.
*   **Notes:** Go does not use `try/catch`. Errors are returned as standard values. I am bundling the HTTP result and the `error` into a single `Result` struct so the caller decides how to handle failures.

### 02_workerpool
*   **Goal:** Introduce concurrency.
*   **Notes:** Replaces the sequential loop with a fixed pool of goroutines. Uses a `jobs` channel to feed URLs to the workers and a `results` channel to get the data back. Uses a `sync.WaitGroup` to block the main program until all workers finish their tasks[span_6](start_span)[span_6](end_span).

### 03_contexts
*   **Goal:** Handle hanging network calls.
*   **Notes:** Uses `context.WithTimeout` to forcefully cancel hanging HTTP requests[span_7](start_span)[span_7](end_span). This ensures workers fail fast and don't leak resources if a server is unresponsive.

### 04_reporting
*   **Goal:** Final output formatting.
*   **Notes:** Aggregates the results stream to calculate success rates and average response times.

## Directory Structure
```text
gopulse/
├── go.mod                  // Module declaration
├── targets.json            // Test URLs
├── README.md               // This file
├── main.go                 // Active working file
└── milestones/             // Reference snapshots
    ├── 01_sequential/main.go
    ├── 02_workerpool/main.go
    ├── 03_contexts/main.go
    └── 04_reporting/main.go
