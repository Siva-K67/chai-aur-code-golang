# Chai aur Code - Golang

A hands-on journey through Go fundamentals to intermediate backend concepts, built by typing out every example by hand rather than copy-pasting — the goal was muscle memory and actually understanding the "why" behind each concept, not just seeing it work once.

## 📚 Learning Sources

- **[Chai aur Code](https://youtube.com/playlist?list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&si=pOk7wSbOJvHWnwGr)** by Hitesh Choudhary — YouTube playlist covering Go fundamentals
- **[How I Write HTTP Services in Go After 13 Years](https://grafana.com/blog/how-i-write-http-services-in-go-after-13-years/)** — Grafana Labs blog on idiomatic, production-grade HTTP service design in Go
- Claude (Anthropic) — used as a learning aid for debugging, explanations, and reinforcing concepts

## 🗂️ Structure

Each folder is a self-contained topic with its own `main.go` (and tests, where applicable), organized roughly in the order concepts were learned:

| Folder | Topic |
|---|---|
| `01-basics-and-setup` | Hello World, GOPATH, lexer/types, variables, comma-ok syntax, type conversions, time handling, system info, memory management |
| `02-core-data-structs` | Pointers, arrays, slices, maps, structs |
| `03-control-flow` | If-else, switch-case, loops (break/continue/goto), functions, methods, defer |
| `04-error-handling` | Error basics, returning errors, wrapping errors, `errors.Is`, `errors.As` |
| `05-files-and-web-basics` | File I/O, handling web requests, query params & JSON, POST JSON bodies |
| `06-courses-api` | A small REST API project tying together earlier concepts (DB integration with PostgreSQL) |
| `07-concurrency` | Goroutines, WaitGroups, race conditions, mutexes, channels, deadlocks |
| `08-interfaces` | Basic interfaces, interface variables, empty interface, mocking |
| `09-testing` | Basic tests, table-driven tests, testing with mocks, `httptest`, table-driven `httptest` |
| `10-context` | Context basics, HTTP context |
| `11-generics` | Generics basics |
| `12-logging` | Logging basics, structured logging, logger as a dependency |

## 🎯 Approach

- Every file was typed out manually — no copy-pasting from tutorials — to build real familiarity with Go syntax and idioms
- Concepts build progressively: basics → data structures → control flow → error handling → web/files → a real API project → concurrency → interfaces → testing → context → generics → logging
- The `06-courses-api` folder marks the shift from isolated concept snippets to an actual working project
- To run the files, type **go run .** in the terminal. To run the test files, use **go test** or **got test -v**.

## 🛠️ Tech

- **Language:** Go
- **Database:** PostgreSQL (for `06-courses-api`)
- **Testing:** Go's standard `testing` package, `httptest`

---

*This repo is a personal learning log — code quality reflects a learning-in-progress mindset rather than production polish, especially in earlier folders.*
