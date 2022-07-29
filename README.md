# Avatar-Generator
Generates random avatars that _look like_ GitHub's identicons. These are **not** actual identicons (yet?)!

## Usage

    go build
    ./avatargen --help
Output is `avatar.png`.

### Example
```
./avatargen --count=20 --size=30
```
<p style="display:block; margin: 0 auto">
<img src="./avatar.png" alt="avatar.png">
</p>

## Options

| options                                   |   Description                                         | 
|-------------------------------------------|-------------------------------------------------------|
|`--count=<block count at least 2>`         |   Number of blocks. By default, 5.                    |
|`--size=<block size at least 2>`           |   Size of each block in pixels. By default, 70.       | 
|`--fgcolor=<hex value #000000 to #ffffff>` |   Color of each block. By default, #64C8C8            |
|`--bgcolor=<hex value #000000 to #ffffff>` |   Background color. By default, #ffffff               |
|`--output=<path of output image>`          |   By default, 'avatar.png' in the current directory   |