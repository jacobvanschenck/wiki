# Wiki CLI

`wiki` is a simple command-line interface (CLI) tool written in Go that allows users to search Wikipedia directly from the terminal. 

## Features

- Use the `-i` flag to get only the introduction section of the article for a quick summary.

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/jacobvanschenck/wiki.git
   ```

2. **Navigate to the project directory:**

   ```bash
   cd wiki
   ```

3. **Build the project:**

   ```bash
   go build -o wiki
   ```

4. **Move the executable to a directory in your PATH (optional):**

   ```bash
   mv wiki /usr/local/bin/
   ```

## Usage

Run the `wiki` command followed by your search term to retrieve the content of a Wikipedia article.

```bash
./wiki "search term"
```

For example:

```bash
./wiki "Albert Einstein"
```

### Summary Option

Use the `-i` flag to fetch only the introduction of the article for a brief overview:

```bash
./wiki -i "search term"
```

For example:

```bash
./wiki -i "Quantum Mechanics"
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to help improve the `wiki` CLI tool.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
```
