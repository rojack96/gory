# Gory

`gory` is a CLI tool written in Go that simulates the behavior of the shell `history` command, with added features to simplify searching and interactively executing past commands.

---

## Key Features

* Reads commands directly from the user’s `~/.bash_history` or `~/.zsh_history` file.
* Automatically detects whether you're using Bash or Zsh.
* Allows searching for commands containing a specific string.
* Allows limiting the number of returned commands.
* Enables interactive execution of a selected command after a terminal confirmation prompt.

---

## Installation

You can install `gory` by building the Go project:

```bash
go install github.com/rojack96/gory/cmd/gory@latest
```

Make sure your `$GOPATH/bin` is in your `PATH`.

---

## Usage

### Options

* `-s`, `--search <string>`
  Search and display all commands in history that contain `<string>`.

* `-n`, `--number <num>`
  Limit the number of displayed commands to `<num>`.

---

### Examples

* Show the last 10 commands:

  ```bash
  gory -n 10
  ```

* Search all commands containing `docker`:

  ```bash
  gory -s docker
  ```

* Search and show up to 5 commands containing `git`:

  ```bash
  gory -s git -n 5
  ```

---

## Important Note

`gory` reads from `~/.bash_history` or `~/.zsh_history`, depending on the active shell.

> To ensure your shell history is up to date before using `gory`, it is recommended to run:

For Bash:

```bash
history -a
```

For Zsh:

```zsh
fc -W
```

---

## Interactive Execution

After displaying matching commands, `gory` will ask if you want to execute one of them via a simple confirmation prompt.

If you confirm, the selected command will be executed directly in your terminal.

---

## Contributing

Pull requests and issues are welcome!

---

## Third-Party Libraries

This project uses the following third-party libraries:

- [charmbracelet/huh](https://github.com/charmbracelet/huh) - MIT License

---

## License

[MIT License](LICENSE)