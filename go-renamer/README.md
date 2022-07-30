# File Renaming Tool exercise

[Source](https://courses.calhoun.io/lessons/les_goph_79)

## Exercise details

In this exercise we are going to explore ways to navigate a file system by creating an application that will rename a bunch of user files in nested directories. The exact files you rename are up to you, but I have provided a sample directory in case you need some ideas. It has the files and directories shown below.

```
sample/
  birthday_001.txt
  birthday_002.txt
  birthday_003.txt
  birthday_004.txt
  christmas 2016 (1 of 100).txt
  christmas 2016 (2 of 100).txt
  christmas 2016 (3 of 100).txt
  christmas 2016 (4 of 100).txt
  christmas 2016 (5 of 100).txt
  nested/
    n_008.txt
    n_009.txt
    n_010.txt
```

The goal of our program is to rename a specific subset of these files. For instance, we might want to take all the files that end in `_NNN.txt` and rename them to instead read `(1 of 4).txt`. Or maybe we will want to rename all of the `XXXXXX_NNN.txt` files to instead read `NNN - XXXXXX.txt`. The exact naming pattern isn't really important, but what IS important is that you can write a program that only modifies the files you specifically want modified, and is able to rename them.
