# Interactive Go Book 1.25

Welcome to **Interactive Go Book** — an educational CLI application with open-source code!

---

## About the Book

This is not just a text book, but an **interactive tutorial** written entirely in **Go**.  
The book combines:  

- **Theory** — chapters explaining key Go concepts.  
- **Practice** — exercises and tasks to reinforce your knowledge.  
- **Progress Saving** — your progress is stored *locally* in `progress.json`.  
- **CLI Project Insight** — see under the hood how a Go CLI tool is structured and works.  

---

## Features

- Navigate through chapters.  
- Read theoretical materials (Markdown).  
- Solve exercises and save progress.  
- Fully offline execution (internet needed only for installation).  

---

## Installation and Running

### Option 1. Using Precompiled Binary  
(Recommended for end users)  

1. Download the precompiled binary from [Releases](./releases).  
2. Extract the archive.  
3. Run the program:  
   ```bash
   ./interactiveGoBook
   ```  
   (on Windows: `interactiveGoBook.exe`)  

### Option 2. Building from Source  
(Suitable if you want to experiment or contribute)  

```bash
git clone https://github.com/Ikorby/Interactive-Go-Book.git
cd Interactive-Go-Book
go run ./cmd/interactiveGoBook
```

Or build the binary:  

```bash
go build -o interactiveGoBook ./cmd/interactiveGoBook
```

---

## Project Structure

```
Interactive Go Book + Exercises/
├── cmd/            # main.go — entry point
│   └── interactiveGoBook/
├── internal/       # project modules (theory, exercises, progress)
├── content/        # book chapters in Markdown format
├── go.mod
└── README.md
```

---

## Roadmap

- Add more chapters.  
- Integrate tests for exercises.  
- Improve console output (colors, Markdown formatting).  
- Add statistics tracking support.  

---

## Contributing

Pull requests are welcome! If you have ideas or fixes:  

1. Fork the repository.  
2. Create a feature branch (`git checkout -b feature/your-feature`).  
3. Make your changes and commit.  
4. Submit a pull request.  

---

## License

This project is licensed under the **GNU Affero General Public License v3.0**.

