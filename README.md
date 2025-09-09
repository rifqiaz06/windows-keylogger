# ⌨️ Windows Keylogger (Educational Purpose Only)

⚠️ **Disclaimer**:
This project is intended **for educational and research purposes only**.
**Do not use this software for any malicious or illegal activity** such as stealing personal data, credentials, or sensitive information.
The author is **not responsible for any misuse** of this code.

---

## 📖 Overview

A simple Windows keylogger written in Go that uses the [`go-hook`](https://github.com/moutend/go-hook) package.
This program captures keyboard events and stores them into a `log.txt` file.

Use cases:

* Learn how **event hooks** work in Windows
* Security research & digital forensics
* Debugging keyboard input in applications

---

## 🛠️ Installation

1. Make sure you have **Go 1.16+** installed
   Download: [https://go.dev/dl/](https://go.dev/dl/)

2. Clone the repository:

   ```bash
   git clone https://github.com/username/windows-keylogger.git
   cd windows-keylogger
   ```

3. Install dependencies:

   ```bash
   go get github.com/moutend/go-hook/pkg/keyboard
   go get github.com/moutend/go-hook/pkg/types
   ```

---

## 🚀 Usage

Build the program:

```bash
go build -o keylogger.exe
```

Run the program:

```bash
keylogger.exe
```

Terminal output:

```
start capturing keyboard input
```

All keystrokes will be saved into a `log.txt` file in the same directory.

---

## 📝 Example Output

`log.txt` contents:

```
A B C D 1 2 3 ENTER SPACE
```

---

## ⚠️ Limitations

* Works only on **Windows** (due to Windows API hooks).
* Captures **keyboard events only** (no mouse or clipboard).
* Logs are saved in plain text → anyone with access can read them.
