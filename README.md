# Avatar-Generator
Generates random avatars that *look like* GitHub's identicons. These are not actual identicons.

## Usage

    go build
    ./avatargen --help
Output is `avatar.png`.
### Options
`--count=<block count at least 2>`     Number of blocks. By default, 5.

`--size=<block size at least 2>`       Size of each block in pixels. By default, 70.  

`--color=<hex value #000000 to #ffffff>` Color of each block. By default, #64C8C8

`--output=<path of output image>` By default, 'avatar.png' in the current directory