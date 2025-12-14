#include <cassert>
#include <cstddef>
#include <fstream>
#include <iostream>
#include <vector>

namespace aoc_04 {

using std::string;
using std::vector;
using u8 = unsigned int;
using usize = std::size_t;

u8 tp_at(const vector<string> &grid, usize x, usize y) {
  assert(grid.size() > 0 and grid[0].size() > 0);
  usize max_y = grid.size() - 1;
  usize max_x = grid[0].size() - 1;

  int range[3]{-1, 0, 1};
  u8 tp_count{0};
  for (auto i : range) {
    for (auto j : range) {
      if (x + i < 0 or x + i > max_x or y + j < 0 or y + j > max_y or
          (i == 0 and j == 0)) {
        continue;
      }
      if (grid[x + i][y + j] == '@') {
        tp_count++;
      }
    }
  }
  return tp_count;
}

u8 count_tp(const vector<string> &grid) {
  assert(grid.size() > 0 and grid[0].size() > 0);
  u8 count{};
  for (usize x{0}; x < grid.size(); ++x) {
    for (usize y{0}; y < grid[0].size(); ++y) {
      // std::cout << grid[x][y];
      if (grid[x][y] == '@') {
        count += (tp_at(grid, x, y) < 4);
      }
    }
    // std::cout << std::endl;
  }
  return count;
}

u8 remove_tp(vector<string> &grid) {
  assert(grid.size() > 0 and grid[0].size() > 0);
  u8 count{};
  while (true) {
    u8 removed{0};
    for (usize x{0}; x < grid.size(); ++x) {
      for (usize y{0}; y < grid[0].size(); ++y) {
        auto &pos{grid.at(x).at(y)};
        if (pos == '@' and tp_at(grid, x, y) < 4) {
          pos = '.';
          removed += 1;
        }
      }
    }
    if (removed == 0)
      break;
    count += removed;
  }
  for (usize x{0}; x < grid.size(); ++x) {
    for (usize y{0}; y < grid[0].size(); ++y) {
      std::cout << grid[x][y];
    }
    std::cout << std::endl;
  }
  return count;
}

} // namespace aoc_04

int main(int argc, char *argv[]) {
  if (argc < 2) {
    std::cerr << "missing input argv" << std::endl;
    return 1;
  }
  std::ifstream input(argv[1]);
  if (!input) {
    std::cerr << "input file cannot be opened" << std::endl;
    return 1;
  }
  std::vector<std::string> grid{};
  std::string line{};
  while (std::getline(input, line)) {
    grid.emplace_back(line);
  }
  aoc_04::u8 sum{aoc_04::remove_tp(grid)};
  std::cout << sum << std::endl;
  return 0;
};
