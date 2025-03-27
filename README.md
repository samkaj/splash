# splash!

[![codecov](https://codecov.io/gh/samkaj/splash/graph/badge.svg?token=BJBZN4QZDW)](https://codecov.io/gh/samkaj/splash)

A parser/serializer-tool for converting your Base16 color palettes between programs.

## Usage

**Define a JSON palette with your theme**, e.g.:

```json
{
  "b00": "#181818",
  "b01": "#282828",
  "b02": "#383838",
  "b03": "#585858",
  "b04": "#b8b8b8",
  "b05": "#d8d8d8",
  "b06": "#e8e8e8",
  "b07": "#f8f8f8",
  "b08": "#ab4642",
  "b09": "#dc9656",
  "b0a": "#f7ca88",
  "b0b": "#a1b56c",
  "b0c": "#86c1b9",
  "b0d": "#7cafc2",
  "b0e": "#ba8baf",
  "b0f": "#a16946"
}
```

**Compile**

```sh
go build -o splash cmd/cli/main.go
```

**Generate a color scheme**

```sh
./splash -i <your-palette>.json wezterm nvim
```

**Usage**

```sh 
./splash -h
```

## How to use my generated color scheme?

When you have generated a scheme, check the instructions for the tool.

## Philosophy

Being a terminal ricer, I have spent
countless hours copy-pasting, (failing at) creating Vim macros just to get
\<insert-colorscheme-here\> in my i3, Alacritty, Neovim and polybar configs.
This project is an attempt to save some time [in the long run](https://www.xkcd.com/974/).
