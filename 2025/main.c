#include "solutions/solutions.h"
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
  if (argc != 4) {
    fprintf(stderr,
            "Invalid number of arguments. Got %d, expected 3.\n"
            "Usage: ./2025 <day> <part> <input file>\n",
            argc - 1);
    exit(1);
  }

  int day = atoi(argv[1]);
  if (day < 1 || day > 12) {
    fprintf(stderr, "Day must be a number in range of 1 - 12.");
    exit(1);
  }

  int part = atoi(argv[2]);
  if (part < 1 || part > 2) {
    fprintf(stderr, "Part must be a number 1 or 2");
    exit(1);
  }

  FILE *input_file = fopen(argv[3], "rb");
  if (input_file == NULL) {
    fprintf(stderr, "Could not open file %s", argv[3]);
    exit(1);
  }

  fseek(input_file, 0, SEEK_END);
  long size = ftell(input_file);
  fseek(input_file, 0, SEEK_SET);

  char *input_str = malloc(size + 1);

  fread(input_str, 1, size, input_file);
  input_str[size] = '\0';

  switch (day) {
  case 1:
    if (part == 1) {
      day1part1(input_str);
    } else {
      day1part2(input_str);
    }
    break;
  }

  free(input_str);
  fclose(input_file);
  exit(0);
}
