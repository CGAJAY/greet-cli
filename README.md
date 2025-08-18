# greet-cli

A beginner-friendly guide and project to learn how to build, run, and install Go CLI tools.

---

## 📌 1. Prerequisites: Linux Basics

Before diving into Go, you should understand a few Linux basics:

-   **Executables** → Files you can run like commands (`ls`, `cd`, etc.).
-   **PATH** → An environment variable that tells Linux where to look for executables.
-   **/usr/local/bin** → System-wide folder where executables are installed (accessible by all users).
-   **~/go/bin** → User-specific folder where Go places executables when you `go install` (accessible only to you unless added to PATH).

👉 To see your PATH, run:

```bash
echo $PATH
```

---

## 📌 2. Install Go on Linux

1. Download Go from the official site:
    ```bash
    wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
    ```
2. Extract and move to `/usr/local`:
    ```bash
    sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
    ```
3. Add Go to your PATH by editing `~/.bashrc` or `~/.zshrc`:
    ```bash
    export GOROOT=/usr/local/go
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
    ```
4. Reload shell:
    ```bash
    source ~/.bashrc
    ```
5. Verify installation:
    ```bash
    go version
    ```

---

## 📌 3. What are GOROOT, GOPATH, PATH?

-   **GOROOT** → Where Go itself is installed (default: `/usr/local/go`).
-   **GOPATH** → Workspace where your Go projects and installed executables live (default: `~/go`).
-   **PATH** → System search path for executables. Adding `$GOPATH/bin` lets you run Go-installed tools as commands.

Example:

-   `/usr/local/go` → Go compiler, stdlib, core tools.
-   `~/go/bin` → Executables you build/install.
-   `/usr/local/bin` → System-wide executables installed manually.

---

## 📌 4. Project Setup: greet-cli

Create your project folder:

```bash
mkdir -p ~/projects/greet-cli
cd ~/projects/greet-cli
```

Initialize a Go module (naming convention: `github.com/<username>/<project>`):

```bash
go mod init github.com/CGAJAY/greet-cli
```

This will create a `go.mod` file in the root of your project folder.

-   The go.mod file is similar to Node.js’s package.json
-   It defines your project’s name, Go version, and tracks dependencies as you add them.

This naming is important because:

-   If you push your project to GitHub, others can `go get github.com/CGAJAY/greet-cli`.
-   If you use just `greet-cli`, it won’t be unique → conflict with others’ packages.

---

## 📌 5. Project Files

### Create `greet.go`

This file defines a separate package called greet. Inside it, we place a function that handles the greeting logic. By keeping this in its own package, the project stays organized and makes it easy to reuse the greeting logic in other programs if needed.

In Go, the first letter of a function name, variable name is very important:

-   Capitalized (`e.g., Greet`) → the function is exported (public) and can be used in other packages (`like main`).
-   Lowercase (`e.g., sayHi`) → the function is unexported (private) and can only be used inside the same package.

### `main.go`

This is the entry point of the program. It uses Go’s flag package to read input from the command line (`like --name=John`). After parsing the input, it calls the Greet function from the greet package to display the greeting.

---

## 📌 6. Running the Program

If you have **multiple files** (`main.go`, `greet.go`, etc.), just run:

```bash
go run . -name=Alice
```

-   The `.` tells Go: _run everything in the current folder._
-   Without `.`, you’d have to type:
    ```bash
    go run main.go greet.go -name=Alice
    ```

---

## 📌 7. Building the Executable

Build the program into a binary (executable):

```bash
go build -o greet-cli
```

Now run it directly:

```bash
./greet-cli -name=Alice
```

-   By default, the executable name is your folder/module name.
-   You can override with `-o` (e.g., `go build -o myapp`).

---

## 📌 8. Installing the Executable

To install into your GOPATH bin (`~/go/bin`):

```bash
go install
```

Now you can run it as a command from anywhere:

```bash
greet-cli -name=Alice
```

If it says **command not found**, ensure your PATH includes `$GOPATH/bin`.

---

## 📌 9. Running vs Building vs Installing

-   **go run** → Compiles & runs immediately (no permanent executable).
-   **go build** → Compiles & creates an executable in the current folder.
-   **go install** → Builds & places the executable in `~/go/bin` (or `$GOPATH/bin`), making it globally available if PATH is set.

---

## 📌 10. Importing & Exporting Functions

-   Functions starting with **uppercase** (like `Greet`) are **exported** → usable in other files/packages.
-   Functions starting with **lowercase** (like `sayHi`) are **unexported** → private to the package.

Example:

```go
// exported
func Greet() {}

// unexported
func sayHi() {}
```

---

## 11. Explain What Packages Mean (and Why Not to Mix)

In Go, **a package is like a folder of tools**.  
Each file in that folder belongs to the same package.

-   If you write `package main`, Go knows this is the **entry point** of a program.  
    It will look for a `func main()` and start execution from there.
-   If you write something else, like `package greet`, then you are making a **reusable library** (not an app).  
    This kind of package is imported and used by other Go programs.

⚠️ **Why not to mix?**  
You cannot have both `package main` and `package greet` in the **same folder**, because Go expects all files in a single folder to belong to **one package only**.

### Example (✅ Correct):

/greet-cli
│
├── main.go // package main
└── greet/
└── greet.go // package greet

### Example (❌ Wrong — cannot mix):

/greet-cli
│
├── main.go // package main
└── greet.go // package greet (❌ cannot live in the same folder as main)

## 📌 12. Summary of Key Points

-   **GOROOT** → Where Go is installed (`/usr/local/go`).
-   **GOPATH** → Workspace for your projects & binaries (`~/go`).
-   **PATH** → System search path; must include `$GOPATH/bin`.
-   **go run** → Quick testing.
-   **go build** → Create an executable locally.
-   **go install** → Create an executable in `~/go/bin`.
-   Use **`go run .`** when running multiple files.

---

## 🚀 Final Test

1. Build locally:

    ```bash
    go build -o greet-cli
    ./greet-cli -name=Alice
    ```

2. Install globally:
    ```bash
    go install
    greet-cli -name=Alice
    ```

Now you’ve made your own CLI tool in Go! 🎉
