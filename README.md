# vex

A lightweight TUI for composing and editing shell commands with vim keybindings.

Built with [goeditor](https://github.com/ionut-t/goeditor).

## Install

```sh
go install github.com/ionut-t/vex@latest
```

## Shell integration

```zsh
# Run vex and push the result to the shell input buffer
function vex() {
  local cmd
  cmd=$(command vex "$@")
  [[ -n "$cmd" ]] && print -z "$cmd"
}

# Edit the current line in vex with ctrl+e (insert and normal mode)
_vex_edit_line() {
  local cmd
  cmd=$(command vex "$BUFFER")
  [[ -n "$cmd" ]] && BUFFER="$cmd" && CURSOR=$#BUFFER
}
zle -N _vex_edit_line
bindkey '^e' _vex_edit_line
bindkey -M vicmd '^e' _vex_edit_line
```

## Usage

```sh
vex                        # open empty editor
vex git log --oneline      # pre-fill with a command
```

Compound commands must be quoted:

```sh
vex "go build && go install"
```

## Keybindings

| Key          | Action                        |
| ------------ | ----------------------------- |
| `:wq` / `:x` | confirm (inserts into prompt) |
| `:q!`        | discard & exit                |
| `ctrl+c`     | discard & exit                |
