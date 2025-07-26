# 🕵️‍♂️ wscan

`wscan` is a fast CLI tool written in Go that recursively searches files in a directory using regular expressions. It outputs structured JSON results by default and supports filtering by extension or ignoring files/directories.

---

## 🚀 Features

- 🔍 Recursive regex scanning of files  
- ⚡ Fast and lightweight  
- 📂 Filter by file extensions  
- 🙈 Ignore specific extensions or folders  
- 🧾 JSON output by default  
- 🎯 Clean flag-based CLI interface  

---

## 📦 Usage

```bash
wscan get --dir="<path>" --regex="<pattern>" --ext="<.ext>" --ignore="<.ext|folder>"
```

### 🧩 Flags

| Flag       | Short | Description                                              |
|------------|-------|----------------------------------------------------------|
| `--dir`    | `-d`  | Directory to scan                                        |
| `--regex`  | `-r`  | Regular expression pattern to search for                |
| `--ext`    | `-e`  | Comma-separated extensions to include (e.g., `.go,.txt`) |
| `--ignore` | `-i`  | Comma-separated extensions or folders to ignore         |

---

## ✅ Examples

Search for `wail` inside `.txt` files:

```bash
wscan get --dir="C:\\Users\\Asus\\OneDrive\\Desktop\\wscan" --regex="wail" --ext=".txt"
```

Ignore `.log` files and the `vendor` directory:

```bash
wscan get --dir="C:\\Users\\Asus\\OneDrive\\Desktop\\wscan" --regex="error" --ignore=".log"
```

---

## 📤 Sample Output

```json
{
  "File": "C:\\Users\\Asus\\OneDrive\\Desktop\\wscan\\wail.txt",
  "Num of line": 2,
  "Match": ["wail"]
}
```

Each match result is printed as a separate JSON object.

---

## 🛠️ Build Instructions

To build the CLI tool from source:

```bash
git clone https://github.com/wailman24/cli-file-search.git
cd cli-file-search
go build -o wscan main.go
```

This will generate an executable named `wscan` in the current directory. You can run it directly:

```bash
./wscan get --dir="..." --regex="..."
```

---

## 👤 Author

**Wail Mansour Ouahchia**  
🔗 GitHub: [wailman24](https://github.com/wailman24)
