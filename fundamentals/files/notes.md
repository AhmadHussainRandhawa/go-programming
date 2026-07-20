# Go File I/O — Revision Notes

> Core mental model: everything in Unix is a file — regular files, keyboard, terminal, pipes, sockets. This is *the* idea that makes Go's `os` package design make sense, and it becomes critical later in networking and systems programming.

---

## Table of Contents

1. [Reading Files](#1-reading-files)
2. [Writing Files](#2-writing-files)
3. [Creating Files](#3-creating-files)
4. [Appending to Files](#4-appending-to-files)
5. [Reading Line by Line & With Buffers](#5-reading-line-by-line--with-buffers)
6. [Copying Files](#6-copying-files)
7. [File Metadata & Existence Checks](#7-file-metadata--existence-checks)
8. [Renaming & Deleting](#8-renaming--deleting)
9. [Directories](#9-directories)
10. [Temporary Files](#10-temporary-files)
11. [Reading from Stdin](#11-reading-from-stdin)
12. [The Golden Pattern](#12-the-golden-pattern)
13. [The 10 APIs to Memorize](#13-the-10-apis-to-memorize)
14. [Practice Project: CLI Notes App](#14-practice-project-cli-notes-app)

---

## 1. Reading Files

### 1.1 Read an Entire File — `os.ReadFile`

**`hello.txt`:**
```
Hello Ahmad
Welcome to Go.
```

**Code:**
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}
```

**Output:**
```
Hello Ahmad
Welcome to Go.
```

- Simplest way to slurp a whole file into memory.
- Returns `[]byte` — convert with `string(data)`.
- Good for small/medium files; not ideal for huge files (loads everything into RAM at once).

---

## 2. Writing Files

### 2.1 Write a Whole File — `os.WriteFile`

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.WriteFile(
		"hello.txt",
		[]byte("Hello Ahmad"),
		0644,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
}
```

- **Creates** the file if it doesn't exist.
- **Overwrites** it completely if it does exist.
- `0644` = Unix permission bits (owner: read/write, group/others: read only).

---

## 3. Creating Files

### 3.1 `os.Create`

```go
package main

import (
	"os"
)

func main() {
	file, err := os.Create("notes.txt")
	if err != nil {
		return
	}
	defer file.Close()
}
```

- Creates a new file (truncates if it already exists).
- Returns a `*os.File` — a handle you can write to directly.

### 3.2 Writing Using the File Object

```go
file, err := os.Create("notes.txt")
if err != nil {
	return
}
defer file.Close()

file.WriteString("Hello\n")
file.WriteString("Go is awesome\n")
```

- Use this when you want more control than `os.WriteFile` (e.g., writing multiple times, streaming data).

---

## 4. Appending to Files

### 4.1 Append (file must already exist)

```go
file, err := os.OpenFile(
	"notes.txt",
	os.O_APPEND|os.O_WRONLY,
	0644,
)
if err != nil {
	return
}
defer file.Close()

file.WriteString("Another line\n")
```

### 4.2 Append + Create If Missing

```go
file, err := os.OpenFile(
	"notes.txt",
	os.O_APPEND|os.O_CREATE|os.O_WRONLY,
	0644,
)
if err != nil {
	return
}
defer file.Close()
```

> **This flag combo (`O_APPEND|O_CREATE|O_WRONLY`) is extremely common for logging systems** — always append, never overwrite, and auto-create the file on first run.

**Key `os.OpenFile` flags:**

| Flag          | Meaning                          |
|---------------|-----------------------------------|
| `O_RDONLY`    | Read only                         |
| `O_WRONLY`    | Write only                        |
| `O_RDWR`      | Read and write                    |
| `O_APPEND`    | Append on write                   |
| `O_CREATE`    | Create if it doesn't exist        |
| `O_TRUNC`     | Truncate file to 0 length on open |
| `O_EXCL`      | Fail if file already exists (used with `O_CREATE`) |

---

## 5. Reading Line by Line & With Buffers

### 5.1 Line by Line — `bufio.Scanner`

**`fruits.txt`:**
```
apple
banana
orange
```

**Code:**
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("fruits.txt")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
```

- Ideal for text files (config files, logs, CSV-like data).
- Always check `scanner.Err()` after the loop — `Scan()` swallows errors silently otherwise.

### 5.2 Reading With a Fixed Buffer

```go
buf := make([]byte, 10)

n, err := file.Read(buf)

fmt.Println(n)
fmt.Println(string(buf[:n]))
```

- Reads at most `len(buf)` bytes at a time.
- Useful for **large files** where loading everything into memory is wasteful — you process data in chunks.

---

## 6. Copying Files

### 6.1 Idiomatic File Copy — `io.Copy`

```go
src, _ := os.Open("a.txt")
defer src.Close()

dst, _ := os.Create("b.txt")
defer dst.Close()

io.Copy(dst, src)
```

- This is **the** idiomatic Go pattern for copying data between any two `io.Reader`/`io.Writer` — not just files (also works for network connections, buffers, etc.).

---

## 7. File Metadata & Existence Checks

### 7.1 Get File Info — `os.Stat`

```go
info, err := os.Stat("hello.txt")
if err != nil {
	return
}

fmt.Println(info.Name())
fmt.Println(info.Size())
fmt.Println(info.IsDir())
fmt.Println(info.ModTime())
```

### 7.2 Check If a File Exists

```go
_, err := os.Stat("hello.txt")

if err == nil {
	fmt.Println("exists")
}
```

Or, more explicitly:

```go
if os.IsNotExist(err) {
	fmt.Println("does not exist")
}
```

> **Tip:** In modern Go, prefer `errors.Is(err, os.ErrNotExist)` over `os.IsNotExist(err)` for wrapped errors, though `os.IsNotExist` is still widely seen in real code and tutorials.

---

## 8. Renaming & Deleting

### 8.1 Rename a File

```go
err := os.Rename(
	"old.txt",
	"new.txt",
)
```

### 8.2 Delete a File

```go
err := os.Remove("hello.txt")
```

---

## 9. Directories

### 9.1 Create a Directory

```go
err := os.Mkdir(
	"data",
	0755,
)
```

- Fails if parent directories don't exist.

### 9.2 Create Nested Directories

```go
err := os.MkdirAll(
	"data/images/profile",
	0755,
)
```

- Creates all missing parent directories — the safe default to reach for.

### 9.3 Read Directory Contents

```go
entries, err := os.ReadDir(".")
if err != nil {
	return
}

for _, entry := range entries {
	fmt.Println(entry.Name())
}
```

### 9.4 Check If an Entry Is a Directory

```go
for _, entry := range entries {
	fmt.Println(
		entry.Name(),
		entry.IsDir(),
	)
}
```

---

## 10. Temporary Files

```go
file, err := os.CreateTemp(
	"",
	"example-*.txt",
)
if err != nil {
	return
}
defer os.Remove(file.Name())
defer file.Close()

fmt.Println(file.Name())
```

- Empty string `""` for the directory means "use the OS default temp dir."
- The `*` in the pattern gets replaced with a random string to avoid collisions.
- Always clean up temp files with `defer os.Remove(...)`.

---

## 11. Reading from Stdin

```go
reader := bufio.NewReader(os.Stdin)

text, _ := reader.ReadString('\n')

fmt.Println(text)
```

- `os.Stdin` is itself just a `*os.File` — reinforcing the "everything is a file" idea.
- Same pattern extends to pipes and sockets later on.

---

## 12. The Golden Pattern

```go
file, err := os.Open("data.txt")
if err != nil {
	return
}
defer file.Close()

// work with file
```

> **This is the single most important pattern in Go file handling.** You will write some variant of *open → check error → defer Close() → use* hundreds of times in real Go code. Internalize it until it's automatic.

---

## 13. The 10 APIs to Memorize

| API              | Purpose                                      |
|------------------|-----------------------------------------------|
| `os.ReadFile`    | Read entire file into memory                  |
| `os.WriteFile`   | Write entire file (create/overwrite)          |
| `os.Open`        | Open file for reading                         |
| `os.Create`      | Create file for writing (truncates existing)  |
| `os.OpenFile`    | Fine-grained open (flags + permissions)        |
| `os.Stat`        | Get file metadata                             |
| `os.Remove`      | Delete a file                                 |
| `os.Rename`      | Rename/move a file                            |
| `os.MkdirAll`    | Create nested directories                     |
| `os.ReadDir`     | List directory contents                       |

Plus one habit to memorize:

```go
defer file.Close()
```

---

## 14. Practice Project: CLI Notes App

Build a tiny CLI note-taking application:

```
notes add "Learn Go"
notes add "Learn Networking"
notes list
notes clear
```

**Skills this exercises:**

- [ ] Creating files
- [ ] Appending to files
- [ ] Reading files
- [ ] Deleting files
- [ ] Error handling
- [ ] `defer`
- [ ] `bufio`
- [ ] `os.OpenFile`

> This is one of the best small exercises for mastering Go's file I/O — it forces you to touch nearly every API in this document in a single, coherent project.