#include <fstream>
#include <iostream>
#include <string>
#include <string_view>
#include <vector>

namespace aoc_02 {

using std::string;
using std::string_view;
using std::vector;

int count_digit(int n) {
  int count{};
  for (; n != 0; n = n / 10) {
    count++;
  }
  return count;
}

bool is_invalid_p1(long long n) {
  string n_str{std::to_string(n)};
  size_t digits{n_str.size()};
  if (digits % 2 != 0) {
    return false;
  }
  string left{n_str.substr(0, digits / 2)};
  string right{n_str.substr(digits / 2, digits / 2)};
  return left == right;
}

vector<int> get_divs(int n) {
  vector<int> divs{};
  divs.push_back(1);
  for (int i = 2; i < n; ++i) {
    if (n % i == 0) {
      divs.push_back(i);
    }
  }
  return divs;
}

bool is_invalid_group(string_view s, int size) {
  string left{s.substr(0, size)};
  for (size_t i = size; i < s.size(); i += size) {
    string group{s.substr(i, size)};
    if (group != left) {
      return false;
    }
  }
  return true;
}

bool is_invalid_p2(long long n) {
  if (n <= 10) {
    return false;
  }
  string n_str{std::to_string(n)};
  auto divs{get_divs(n_str.size())};
  for (auto d : divs) {
    if (is_invalid_group(n_str, d)) {
      return true;
    }
  }
  return false;
}

long long sum_invalids(string_view range) {
  size_t idx{range.find('-')};

  string from_str{range.substr(0, idx)};
  long long from{std::stoll(from_str)};

  string to_str{range.substr(idx + 1, range.size() - idx)};
  long long to{std::stoll(to_str)};

  long long sum{};
  for (; from <= to; from++) {
    if (is_invalid_p2(from)) {
      sum += from;
    }
  }
  return sum;
}

}; // namespace aoc_02

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
  std::string range{};
  long long sum{};
  while (std::getline(input, range, ',')) {
    sum += aoc_02::sum_invalids(range);
  }
  std::cout << sum << std::endl;
  return 0;
}
