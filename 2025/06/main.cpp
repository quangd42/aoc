#include <fstream>
#include <functional>
#include <iostream>
#include <numeric>
#include <ostream>
#include <stdexcept>
#include <string>

#include "../common/common.h"

string trim(const string& str) {
  const string ws = " \t\n\r\f\v";

  const size_t first = str.find_first_not_of(ws);
  if (string::npos == first) {
    return str;
  }
  const size_t last = str.find_last_not_of(ws);

  // Extract the trimmed substring
  return str.substr(first, last + 1 - first);
}

vector<string> split(string& line) {
  line = trim(line);
  const string ws = " \t\n\r\f\v";
  vector<string> out{};
  int cbegin_idx{};
  int cend_idx{};

  while (cend_idx != string::npos) {
    cbegin_idx = line.find_first_not_of(ws, cend_idx);
    cend_idx = line.find_first_of(ws, cbegin_idx);
    // apparently when len is negative substr will go to the end of string
    out.emplace_back(line.substr(cbegin_idx, cend_idx - cbegin_idx));
  }
  return out;
}

void print_fields(vector<string>& fields) {
  for (auto f : fields) {
    std::cout << f << " ";
  }
  std::cout << std::endl;
}

int part1(std::ifstream& input) {
  string line{};
  u64 sum{};
  vector<vector<u64>> table{};
  vector<string> opers{};
  vector<u64> sums{};
  while (std::getline(input, line)) {
    auto fields{split(line)};
    for (auto i{0}; i < fields.size(); i++) {
      if (i >= table.size()) {
        table.emplace_back(vector<u64>{});
      }
      try {
        table[i].push_back(std::stoull(fields.at(i)));
      } catch (const std::invalid_argument& e) {
        opers = fields;
        break;
      }
    }
    print_fields(fields);
  }
  for (auto i{0}; i < table.size(); i++) {
    std::cout << "parsed nums: ";
    for (auto n : table[i]) {
      std::cout << n << " ";
    }
    std::cout << "oper = " << opers.at(i);
    u64 sub_acc{};
    if (opers.at(i) == "+") {
      sub_acc = std::accumulate(table[i].cbegin(), table[i].cend(), 0ull);
    } else if (opers.at(i) == "*") {
      sub_acc = std::accumulate(table[i].cbegin(), table[i].cend(), 1ull, std::multiplies<u64>());
    } else {
      std::cout << "parsing operator failed" << "\n";
      return 1;
    }
    sums.emplace_back(sub_acc);
    std::cout << " sub_acc = " << sub_acc << std::endl;
  }
  sum = std::accumulate(sums.cbegin(), sums.cend(), 0ull);
  std::cout << "part1: " << sum << std::endl;
  return 0;
}

int part2(std::ifstream& input) {
  string line{};
  vector<string> lines{};
  u64 sum{};
  size_t max_len{};
  while (std::getline(input, line)) {
    if (line.size() > max_len) {
      // find longest line
      max_len = line.size();
    }
    lines.push_back(line);
  }

  // find the last line to use as the guide
  string opers = lines.at(lines.size() - 1);
  // but need to pad the last line to have the same length as the longest line
  if (size_t pad_size = max_len - opers.size(); pad_size > 0) {
    string pad(pad_size, ' ');
    opers += pad;
  }

  vector<u64> numbers{};
  for (int i{static_cast<int>(opers.size() - 1)}; i >= 0; i--) {
    /* add number to numbers */
    string num{};
    for (int j{0}; j < lines.size() - 1; j++) {
      // concat all chars in column i (except for the guide), ignoring empty char
      if (auto ch{lines.at(j).at(i)}; ch != ' ') {
        num += ch;
      }
    }
    if (num.empty()) {
      // ignore empty columns
      continue;
    }
    // finally convert column to number and add to number list
    numbers.push_back(std::stoul(num));

    /* if operator (add or mult) is found in the guide, do it and reset numbers */
    if (opers.at(i) == '+') {
      sum += std::accumulate(numbers.cbegin(), numbers.cend(), 0ull);
      numbers.clear();
    } else if (opers.at(i) == '*') {
      sum += std::accumulate(numbers.cbegin(), numbers.cend(), 1ull, std::multiplies<u64>());
      numbers.clear();
    } else if (opers.at(i) != ' ') {
      std::cout << "parsing operator failed" << "\n";
      return 1;
    }
  }
  std::cout << "part2: " << sum << std::endl;
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
  // if (part1(input))
  //   return 1;
  if (part2(input))
    return 1;
  return 0;
};
