
#include <cstddef>
#include <fstream>
#include <iostream>
#include <string>
#include <string_view>

namespace aoc_o3 {

int largest2(std::string_view s) {
  size_t i1{}, i2{1};
  char c1{s.at(0)}, c2{s.at(1)};
  for (size_t i{1}; i < s.size(); ++i) {
    if (s.at(i) > c1 and i != s.size() - 1) {
      i1 = i;
      c1 = s.at(i);
      i2 = i + 1;
      c2 = s.at(i + 1);
    } else if (s.at(i) > c2) {
      i2 = i;
      c2 = s.at(i);
    }
  }
  return (c1 - '0') * 10 + (c2 - '0');
}

std::string largestn(std::string_view s, int idx) {
  if (idx < 1) {
    return "";
  }
  size_t len{s.size()};
  if (len < idx) {
    return "";
  }
  char hi{s.at(0)};
  size_t hi_idx{0};
  for (auto i{0}; i <= len - idx; ++i) {
    if (s.at(i) > hi) {
      hi = s.at(i);
      hi_idx = i;
    }
  }
  return hi + largestn(s.substr(hi_idx + 1, len - hi_idx - 1), idx - 1);
}

} // namespace aoc_o3

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
  std::string line{};
  long long sum{};
  while (std::getline(input, line)) {
    long long n = std::stoll(aoc_o3::largestn(line, 12));
    // std::cout << line << " -> " << n << "\n";
    sum += n;
  }
  std::cout << sum << std::endl;
  return 0;
};
