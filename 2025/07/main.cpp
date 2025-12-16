#include "../common/common.h"
#include <cassert>
#include <fstream>
#include <iostream>
#include <stack>
#include <unordered_set>

using set = std::unordered_set<int>;

set find_splitters(const string& row) {
  set out{};
  for (auto i{0}; i < row.size(); i++) {
    if (row.at(i) == '^') {
      out.insert(i);
    }
  }
  return out;
}

int split(set& beams, set& splitters) {
  int split_count{};
  set updated{};
  for (auto b : beams) {
    bool split{false};
    for (auto s : splitters) {
      if (s == b) {
        split = true;
        split_count++;
        break;
      }
    }
    if (!split) {
      updated.insert(b);
    } else {
      updated.insert(b - 1);
      updated.insert(b + 1);
    }
  }
  beams = updated;
  return split_count;
}

void render(string& row, const set& beams) {
  for (auto b : beams) {
    row.at(b) = '|';
  }
}

void add_beams(vector<int>& beams, vector<int>& new_beams) {
  vector<int> combined{beams};
  for (auto nb : new_beams) {
    bool add{true};
    for (auto b : beams) {
      if (nb == b) {
        add = false;
        break;
      }
    }
    if (add) {
      combined.push_back(nb);
    }
  }
  beams = combined;
}

void print_grid(vector<string>& grid) {
  for (auto row : grid) {
    std::cout << row << std::endl;
  }
}

int part1(std::ifstream& input) {

  // pass arg example_expected.txt
  // std::ifstream expected(argv[2]);
  // if (!expected) {
  //   std::cerr << "expected file cannot be opened" << std::endl;
  //   return 1;
  // }

  string line{};
  vector<string> grid{};
  while (std::getline(input, line)) {
    grid.emplace_back(line);
  }
  // vector<string> expected_grid{};
  // while (std::getline(expected, line)) {
  //   expected_grid.emplace_back(line);
  // }

  std::unordered_set<int> beams{};
  beams.insert(grid[0].find('S'));

  int split_count{};
  for (auto i{1}; i < grid.size(); i++) {
    auto& row{grid.at(i)};
    // find splitters
    auto splitters = find_splitters(row);
    // if beam meets splitter, remove it from beam list and add back two new beams
    split_count += split(beams, splitters);
    // render beams
    render(row, beams);
  }

  print_grid(grid);
  std::cout << std::endl << "split " << split_count << " times";

  // for (auto i{0}; i < grid.size(); i++) {
  //   assert(grid[i] == expected_grid[i]);
  // }
  return 0;
}

struct coord {
  size_t x, y;
};

int part2dfs(std::ifstream& input) {
  string line{};
  vector<string> grid{};
  while (std::getline(input, line)) {
    grid.emplace_back(line);
  }

  u32 sum{};
  size_t max_y{grid.size() - 1};
  size_t max_x{grid[0].size() - 1};
  coord start{grid[0].find('S'), 0};
  std::stack<coord> stack{};
  stack.push(start);

  while (!stack.empty()) {
    auto p{stack.top()};
    stack.pop();
    while (p.y <= max_y) {
      // std::cout << "p(" << p.x << ", " << p.y << ")" << std::endl;
      // std::cout << grid[p.y][p.x] << std::endl;

      if (grid[p.y][p.x] != '^') {
        p.y++;
      } else {
        // stack.push({p.x == max_x ? max_x : p.x + 1, p.y});
        stack.push({p.x + 1, p.y});
        // p.x = (p.x == 0) ? 0 : p.x - 1;
        p.x--;
      }
    }
    sum += 1;
    std::cout << sum << " ";
  }
  std::cout << sum << " different timelines" << "\n";
  return 0;
}

int part2dp(std::ifstream& input) {
  string line{};
  vector<string> grid{};
  while (std::getline(input, line)) {
    grid.emplace_back(line);
  }

  u64 sum{};
  size_t max_y{grid.size() - 1};
  size_t max_x{grid[0].size() - 1};

  vector<vector<u64>> tlg{};
  for (auto _ : grid) {
    tlg.emplace_back(vector<u64>(max_x + 1, 0ull));
  }

  size_t start_x{grid[0].find('S')};

  // set up beams
  std::unordered_set<int> beams{};
  beams.insert(start_x);
  // set up timeline grid
  tlg.at(0).at(start_x) = 1;

  // draw all beams on the grid
  for (auto i{1}; i < grid.size(); i++) {
    auto& row{grid.at(i)};
    // find splitters
    auto splitters = find_splitters(row);
    // if beam meets splitter, remove it from beam list and add back two new beams
    split(beams, splitters);
    // render beams
    render(row, beams);
    // count how many ways to get to each beam
    for (auto b : beams) {
      auto& prev_row{grid.at(i - 1)};
      auto& cur_row{grid.at(i)};
      if (auto prev_cell{prev_row.at(b)}; prev_cell == '|' or prev_cell == 'S') {
        tlg.at(i).at(b) = tlg.at(i - 1).at(b);
      }
      if (b > 0 and cur_row.at(b - 1) == '^') {
        tlg.at(i).at(b) += tlg.at(i - 1).at(b - 1);
      }
      if (b < max_x and cur_row.at(b + 1) == '^') {
        tlg.at(i).at(b) += tlg.at(i - 1).at(b + 1);
      }
    }
  }

  // print_grid(grid);

  // for (auto r : tlg) {
  //   for (auto n : r) {
  //     if (n == 0) {
  //       std::cout << '.' << " ";
  //     } else {
  //       std::cout << n << " ";
  //     }
  //   }
  //   std::cout << std::endl;
  // }

  for (u64 tl : tlg.at(tlg.size() - 1)) {
    sum += tl;
  }

  std::cout << sum << " different timelines" << "\n";
  return 0;
}

int main(int argc, char* argv[]) {
  if (argc < 2) {
    std::cerr << "missing input argv" << std::endl;
    return 1;
  }
  std::ifstream input(argv[1]);
  if (!input) {
    std::cerr << "input file cannot be opened" << std::endl;
    return 1;
  }
  if (part2dp(input)) {
    return 1;
  }
  return 0;
};
