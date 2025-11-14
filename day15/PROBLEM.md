<p>You appear back inside your own mini submarine! Each Historian drives their mini submarine in a different direction; maybe the Chief has his own submarine down here somewhere as well?</p>
<p>You look up to see a vast school of <a href="/2021/day/6">lanternfish</a> swimming past you. On closer inspection, they seem quite anxious, so you drive your mini submarine over to see if you can help.</p>
<p>Because lanternfish populations grow rapidly, they need a lot of food, and that food needs to be stored somewhere. That's why these lanternfish have built elaborate warehouse complexes operated by robots!</p>
<p>These lanternfish seem so anxious because they have lost control of the robot that operates one of their most important warehouses! It is currently running <span title="Wesnoth players might solve their Warehouse Woes with a Warehouse Wose!">amok</span>, pushing around boxes in the warehouse with no regard for lanternfish logistics <em>or</em> lanternfish inventory management strategies.</p>
<p>Right now, none of the lanternfish are brave enough to swim up to an unpredictable robot so they could shut it off. However, if you could anticipate the robot's movements, maybe they could find a safe option.</p>
<p>The lanternfish already have a map of the warehouse and a list of movements the robot will <em>attempt</em> to make (your puzzle input). The problem is that the movements will sometimes fail as boxes are shifted around, making the actual movements of the robot difficult to predict.</p>
<p>For example:</p>

```
##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^
```

<p>As the robot (<code>@</code>) attempts to move, if there are any boxes (<code>O</code>) in the way, the robot will also attempt to push those boxes. However, if this action would cause the robot or a box to move into a wall (<code>#</code>), nothing moves instead, including the robot. The initial positions of these are shown on the map at the top of the document the lanternfish gave you.</p>
<p>The rest of the document describes the <em>moves</em> (<code>^</code> for up, <code>v</code> for down, <code><</code> for left, <code>></code> for right) that the robot will attempt to make, in order. (The moves form a single giant sequence; they are broken into multiple lines just to make copy-pasting easier. Newlines within the move sequence should be ignored.)</p>
<p>Here is a smaller example to get started:</p>

```
########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<
```

<p>Were the robot to attempt the given sequence of moves, it would push around the boxes as follows:</p>

```
Initial state:
########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

Move <:
########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

Move ^:
########
#.@O.O.#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

Move ^:
########
#.@O.O.#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

Move >:
########
#..@OO.#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

Move >:
########
#...@OO#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

Move >:
########
#...@OO#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

Move v:
########
#....OO#
##..@..#
#...O..#
#.#.O..#
#...O..#
#...O..#
########

Move v:
########
#....OO#
##..@..#
#...O..#
#.#.O..#
#...O..#
#...O..#
########

Move <:
########
#....OO#
##.@...#
#...O..#
#.#.O..#
#...O..#
#...O..#
########

Move v:
########
#....OO#
##.....#
#..@O..#
#.#.O..#
#...O..#
#...O..#
########

Move >:
########
#....OO#
##.....#
#...@O.#
#.#.O..#
#...O..#
#...O..#
########

Move >:
########
#....OO#
##.....#
#....@O#
#.#.O..#
#...O..#
#...O..#
########

Move v:
########
#....OO#
##.....#
#.....O#
#.#.O@.#
#...O..#
#...O..#
########

Move <:
########
#....OO#
##.....#
#.....O#
#.#O@..#
#...O..#
#...O..#
########

Move <:
########
#....OO#
##.....#
#.....O#
#.#O@..#
#...O..#
#...O..#
########
```

<p>The larger example has many more moves; after the robot has finished those moves, the warehouse would look like this:</p>

```
##########
#.O.O.OOO#
#........#
#OO......#
#OO@.....#
#O#.....O#
#O.....OO#
#O.....OO#
#OO....OO#
##########
```

<p>The lanternfish use their own custom Goods Positioning System (GPS for short) to track the locations of the boxes. The <em>GPS coordinate</em> of a box is equal to 100 times its distance from the top edge of the map plus its distance from the left edge of the map. (This process does not stop at wall tiles; measure all the way to the edges of the map.)</p>
<p>So, the box shown below has a distance of <code>1</code> from the top edge of the map and <code>4</code> from the left edge of the map, resulting in a GPS coordinate of <code>100 * 1 + 4 = 104</code>.</p>

```
#######
#...O..
#......
```

<p>The lanternfish would like to know the <em>sum of all boxes' GPS coordinates</em> after the robot finishes moving. In the larger example, the sum of all boxes' GPS coordinates is `10092`. In the smaller example, the sum is `2028`.</p>
<p>Predict the motion of the robot and boxes in the warehouse. After the robot is finished moving, <em>what is the sum of all boxes' GPS coordinates?</em></p>
