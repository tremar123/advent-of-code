#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void day1part1(char *input) {
  int zeroes = 0;
  int current = 50;

  char *line = strtok(input, "\n");
  for (; line; line = strtok(NULL, "\n")) {
    // +1 to exclude first char which is L or R for direction
    int clicks = atoi(line + 1);
    if (*line == 'L') {
      int move_by = clicks % 100;
      current -= move_by;
      if (current < 0) {
        current = 100 + current;
      }
    } else {
      int move_by = clicks % 100;
      current += move_by;
      if (current > 99) {
        current = 0 + current - 100;
      }
    }

    if (current == 0) {
      zeroes++;
    }
  }

  printf("%d\n", zeroes);
}

void day1part2(char *input) {
  int zeroes = 0;
  int current = 50;

  char *line = strtok(input, "\n");
  for (; line; line = strtok(NULL, "\n")) {
    // +1 to exclude first char which is L or R for direction
    int clicks = atoi(line + 1);
    if (*line == 'L') {
      int move_by = clicks % 100;
      int last = current;
      current -= move_by;
      if (current < 0) {
        if (last != 0) {
          zeroes++;
        }
        current = 100 + current;
      }
      if (current == 0) {
        zeroes++;
      }
      zeroes += clicks / 100;
    } else {
      int move_by = clicks % 100;
      current += move_by;
      if (current > 99) {
        current = 0 + current - 100;
        zeroes++;
      }
      zeroes += clicks % 100 == 0 ? 0 : clicks / 100;
    }
  }

  printf("%d\n", zeroes);
}
