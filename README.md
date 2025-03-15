# splash!

A parser/serializer-tool for converting your Base16 color palettes between programs.

By defining modules for parsing and serializing to the internal representation
of a color scheme, this tool is designed to easily create new adapters to
convert e.g., iTerm colors (XML) to ghostty colors automatically.

## Philosophy

Being a terminal ricer, I have spent
countless hours copy-pasting, (failing at) creating Vim macros just to get
\<insert-colorscheme-here\> in my i3, Alacritty, Neovim and polybar configs.
This project is an attempt to save some time [in the long run](https://www.xkcd.com/974/).
