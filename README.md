# Go Programming Journey

A structured, Git-driven learning repository documenting my journey of learning Go through hands-on coding exercises, small programs, and practical experiments.

Rather than treating this repository as a collection of tutorial code, I maintain it using professional software engineering practices. Each topic is developed incrementally, committed as a focused change, and organized into logical modules to create a clean and reviewable project history.

---

## Purpose

This repository serves as an engineering notebook that documents my progress while learning Go.

My goals are to:

* Learn Go from first principles through practical implementation.
* Build strong fundamentals before moving to backend systems and distributed software.
* Maintain a clean and meaningful Git history.
* Practice professional repository organization and version control workflows.
* Create a long-term reference for concepts I have learned.

---

## Learning Approach

Instead of copying code into a single project, each concept is implemented independently in its own directory.

Development follows a disciplined Git workflow:

* Feature branches for major learning milestones.
* Small, atomic commits focused on individual concepts.
* Conventional Commit messages.
* Annotated release tags for completed modules.
* Reviewed merges into the `main` branch.

This allows the repository history to reflect the progression of learning rather than simply the order of tutorial videos.

---

## Repository Structure

```text
go-programming/
├── fundamentals/
│   ├── hello/
│   ├── variables/
│   ├── input/
│   ├── type-conversion/
│   ├── time/
│   ├── pointers/
│   ├── arrays/
│   ├── slices/
│   ├── maps/
│   ├── structs/
│   ├── if-else/
│   ├── switch/
│   ├── loops/
│   ├── functions/
│   ├── struct-methods/
│   ├── defer/
│   └── files/
├── go.mod
└── README.md
```

As additional modules are completed, new top-level directories will be added while preserving a clear and consistent project layout.

---

## Progress

### ✅ Module 1 — Go Fundamentals

* Hello World
* Variables and Constants
* User Input
* Type Conversion
* Time Handling
* Pointers
* Arrays
* Slices
* Maps
* Structs
* Conditional Statements
* Switch Statements
* Loops
* Functions
* Struct Methods
* Defer
* File Operations

### 🚧 Planned Modules

* HTTP Client
* HTTP Server & REST APIs
* MongoDB
* Concurrency
* Utilities

---

## Running Examples

Each topic is self-contained and can be executed independently.

Example:

```bash
go run fundamentals/hello/main.go
```

Another example:

```bash
go run fundamentals/maps/main.go
```

---

## Technologies

* Go
* Go Modules
* Git
* GitHub

---

## Course Reference

The examples in this repository are based on the excellent Go programming course by **Hitesh Choudhary**.

This repository is my own implementation, experimentation, and learning record while following the course.

---

## License

This repository is available under the MIT License unless stated otherwise.
