# EHFOLDER
Plan your new folders and files and create them using the EHFOLDER CLI.

## 📚 Tutorial

### ✅ Prerequisites

Before you begin, make sure you have the following installed on your system:

* [Go (Golang)](https://golang.org/dl/) installed and properly set up in your system’s `PATH`.
* A text editor of your choice (e.g., VS Code, Sublime Text, etc.).

---

### 📝 Syntax Guide

Each folder is indicated by a line starting with `#`, and files within the folder are listed below it, each prefixed by `-`. So, to specify a:
  - **folder**: use the "#" character followed by the a white space and the name of the folder to be created.
  - **file**: use the "-" character followed by the a white space and the name of the file to be created.

**Example:**

```
# FolderOne
- file1.txt
- file2.md
# FolderTwo
- script.go
- README2.md
```

---

### 🚀 Step-by-Step Instructions

1. **Create a new folder** anywhere on your system to work in.
2. **Add the `main.go` file** to this folder.
3. **Create an input file** in the same folder, named something like `input-structure.txt` or `input-structure.md`. This file should describe the folder/file structure using the syntax shown above.
4. **Open a terminal**, and navigate to the folder you created using the `cd` command.
5. **Run the script** using the following commands:

#### If your input file is a `.txt` file:

```bash
go run main.go -input=input-structure.txt -output=./output-folder
```

#### If your input file is a `.md` file:

```bash
go run main.go -input=input-structure.md -output=./output-folder
```

This will generate the described folder and file structure inside the specified `output-folder`.
