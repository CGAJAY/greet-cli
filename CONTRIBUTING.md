# Contributing to greet-cli

Thank you for your interest in contributing to **greet-cli**! 🚀 This project is designed to be beginner-friendly while exploring Go, Linux, and terminal applications. Whether you’re fixing bugs, adding features, or improving documentation, your contributions are welcome.

---

## Table of Contents

- [Contributing to greet-cli](#contributing-to-greet-cli)
  - [Table of Contents](#table-of-contents)
  - [How to Contribute](#how-to-contribute)
  - [Reporting Issues](#reporting-issues)
  - [Submitting Pull Requests](#submitting-pull-requests)
  - [Coding Guidelines](#coding-guidelines)
  - [Project Roadmap](#project-roadmap)
  - [Community \& Support](#community--support)

---

## How to Contribute

1. **Fork the repository**: Click the “Fork” button on GitHub to create your own copy of the project.
2. **Clone your fork locally**:

```bash
git clone https://github.com/<your-username>/greet-cli.git
cd greet-cli
```

3. **Create a new branch** for your contribution:

```bash
git checkout -b feature/my-new-feature
```

4. **Make your changes**: Fix bugs, add features (colors, text effects, music, interactive modes), or improve documentation.
5. **Test your changes** locally:

```bash
go run . -name=Alice
```

6. **Commit and push**:

```bash
git add .
git commit -m "Add feature: rainbow text effect"
git push origin feature/my-new-feature
```

7. **Submit a pull request** on GitHub.

---

## Reporting Issues

-   Check existing issues before opening a new one.
-   Provide a clear description and steps to reproduce.
-   Include your OS, Go version, and any relevant logs.

Example:

```
OS: Ubuntu 22.04
Go: 1.22
Command: go run . -name=Alice
Error: unexpected output with colors flag
```

---

## Submitting Pull Requests

-   Base your PR on the `dev` branch.
-   Include a clear title and description.
-   Reference related issues:

```markdown
Fixes #12
```

-   Keep commits small and focused.
-   Ensure your code passes existing tests.

---

## Coding Guidelines

-   **Follow Go conventions**: `gofmt` or `go fmt` your code.
-   **Function names**:

    -   Exported: `Greet`, `PlayMusic`
    -   Unexported: `sayHi`, `applyEffect`

-   **Folder structure**:

```
greet-cli/
├── main.go        # CLI entry point
├── greet/
│   └── greet.go   # Greeting logic
├── effects/
│   └── effects.go # Text effects & colors
├── sound/
│   └── sound.go   # Optional music/audio
└── README.md
```

-   **Tests**: Add tests for new functions when possible.

---

## Project Roadmap

1. **Text Effects & Colors**: Bold, underline, rainbow, ASCII art banners
2. **Interactive Mode**: Real-time prompts for name, color, effect, music
3. **Music & Sounds**: Play short audio files during greetings
4. **Presets & Themes**: Save favorite combinations for quick greetings
5. **Community Contributions**: Share new effects, sounds, or interactive features; add tutorials or example scripts

---

## Community & Support

-   Join discussions in [Issues](https://github.com/CGAJAY/greet-cli/issues)
-   Ask questions or suggest features
-   Be respectful and follow good open-source etiquette

---

Thank you for helping make **greet-cli** better! Every contribution—big or small—makes a difference. 💙
