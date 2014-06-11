CLI2048
=======

Console version of the popular '2048' game written in Go.

Usage
-----

    cli2048 [flags]
        -v          Display version
        --version
        -h          Display help info
        --help

Controls
--------

    UP      w, i, [UP ARROW]
    DOWN    s, k, [DOWN ARROW]
    LEFT    a, j, [LEFT ARROW]
    RIGHT   d, l, [RIGHT ARROW]

    HELP    h

    RESET       r, n
    NEW GAME    r, n

    QUIT    q, [ESC]

Objective
---------
The game takes place on a 4x4 grid with tiles randomly generated in empty grid squares. Your goal is move the tiles around the board until they combine like numbers until one or more tiles reaches 2048. After every movement another tile will be generated. The game is over when there are no possible moves left. Points are earned by combining like tiles.

Rules
-----
Tiles may be shifted either UP, DOWN, LEFT, or RIGHT

All tiles will be shifted at the same time in one of the four directions

When two tiles with the same value are shifted into each other they will combine their values and position on the board
