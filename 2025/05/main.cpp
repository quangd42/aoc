#include <algorithm>
#include <cstddef>
#include <deque>
#include <fstream>
#include <iostream>
#include <optional>

using std::pair;
using std::string;
using std::vector;
using u8 = unsigned int;
using u64 = long long;
using range = pair<u64, u64>;
using ranges = std::deque<pair<u64, u64>>;

range parse_range(const string& s) {
  auto d{s.find('-')};
  u64 from{std::stoll(s.substr(0, d))};
  u64 to{std::stoll(s.substr(d + 1, s.size() - d - 1))};
  return range(from, to);
}

void update_ranges(ranges& ranges, const string& s) {
  ranges.emplace_back(parse_range(s));
}

bool is_fresh(ranges& ranges, const string& s) {
  u64 id{std::stoll(s)};
  for (auto& range : ranges) {
    if (id >= range.first and id <= range.second)
      return true;
  }
  return false;
}

std::optional<range> merge_ranges(range& r1, range& r2) {
  range out{0, 0};
  if (r1.first <= r2.first) {
    if (r1.second < r2.first) {
      // no overlopping, return fail
      return std::nullopt;
    } else if (r1.second <= r2.second) {
      // overlapping
      out.first = r1.first;
      out.second = r2.second;
    } else {
      // r2 is subset of r1
      out = r1;
    }
  } else {
    // (r1.first > r2.first)
    if (r1.first > r2.second) {
      // no overlopping, return fail
      return std::nullopt;
    } else if (r2.second <= r1.second) {
      // overlapping
      out.first = r2.first;
      out.second = r1.second;
    } else {
      // r1 is subset of r2
      out = r2;
    }
  }
  return out;
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
  // part 1
  string line{};
  ranges rs{};
  while (std::getline(input, line)) {
    if (line.empty())
      break;
    update_ranges(rs, line);
  }
  long long sum{};
  while (std::getline(input, line)) {
    sum += is_fresh(rs, line);
  }
  std::cout << sum << std::endl;

  // part 2
  std::sort(rs.begin(), rs.end());
  bool merged{true};
  ranges rs2{};
  while (merged) {
    merged = false;
    while (rs.size() > 1) {
      range r1 = rs.front();
      rs.pop_front();
      range r2 = rs.front();
      auto maybe_range{merge_ranges(r1, r2)};
      if (!maybe_range) {
        rs2.push_back(r1);
      } else {
        rs.pop_front(); // r2
        rs2.push_back(maybe_range.value());
        merged = true;
      }
    }
    if (!rs.empty()) {
      range r1 = rs.front();
      rs.pop_front();
      rs2.push_back(r1);
    }
    rs2.swap(rs);
  }

  sum = 0;
  for (auto r : rs) {
    // std::cout << "range = " << r.first << " -> " << r.second << "\n";
    sum += r.second - r.first + 1;
  }
  std::cout << sum << std::endl;

  return 0;
};
