#include <stdint.h>
#include <stdio.h>
#include <time.h>

#define TEST_LOOPS 100000000

// switched page_size from being the size itself to just he shift size
// no performance improvement though :(
uint64_t pagecount(uint64_t memory_size, uint64_t page_size) {
  return memory_size >> page_size;
}

int main (int argc, char** argv) {
  clock_t baseline_start, baseline_end, test_start, test_end;
  uint64_t memory_size, page_size;
  double clocks_elapsed, time_elapsed;
  int i, ignore = 0;

  uint64_t msizes[] = {1L << 32, 1L << 40, 1L << 52};
  // psizes are now just the shift sizes
  uint64_t psizes[] = {12, 16, 32};

  baseline_start = clock();
  for (i = 0; i < TEST_LOOPS; i++) {
    memory_size = msizes[i % 3];
    page_size = psizes[i % 3];
    ignore += 1 + memory_size + page_size; // so that this loop isn't just optimized away
  }
  baseline_end = clock();

  test_start = clock();
  for (i = 0; i < TEST_LOOPS; i++) {
    memory_size = msizes[i % 3];
    page_size = psizes[i % 3];

    // inlining gives the biggest improvement! wow huge ~1.7ns -> 0.25ns
    // wait is that too low to be plausible? hmmm
    ignore += (memory_size >> page_size) + memory_size + page_size;
  }
  test_end = clock();

  clocks_elapsed = test_end - test_start - (baseline_end - baseline_start);
  time_elapsed = clocks_elapsed / CLOCKS_PER_SEC;

  printf("%.2fs to run %d tests (%.2fns per test)\n", time_elapsed, TEST_LOOPS,
         time_elapsed * 1e9 / TEST_LOOPS);
  return ignore;
}

