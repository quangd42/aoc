#include "../common/common.h"
#include <cstdio>
#include <cstdlib>
#include <fstream>
#include <iostream>
#include <string>

using button = vector<usize>;

struct machine {
  vector<bool> goal;
  vector<button> buttons;

  friend std::ostream& operator<<(std::ostream& os, const machine& m) {
    os << "[";
    for (auto g : m.goal) { os << (g ? '#' : '.'); }
    os << "] ";

    for (const auto& b : m.buttons) {
      os << "(";
      auto b_size = b.size();
      for (auto i{0}; i < b_size; i++) {
        os << b[i];
        if (i != b_size - 1) os << ",";
      }
      os << ") ";
    }
    os << std::endl;
    return os;
  }
};

machine parse_machine(const string& line) {
  vector<bool> goal{};
  {
    auto start = line.find("[");
    auto end = line.find("]");
    string goal_string = line.substr(start + 1, end - start - 1);
    for (auto& c : goal_string) {
      if (c == '.') {
        goal.emplace_back(false);
      } else if (c == '#') {
        goal.emplace_back(true);
      } else {
        std::printf("invalid goal string: %s\n", line.c_str());
      }
    }
  }

  vector<button> buttons{};
  {
    auto end = 0;
    auto start = 0;
    while (true) {
      start = line.find("(", end + 1);
      end = line.find(")", start);
      if (start == string::npos) break;
      auto button_strings = split(line.substr(start + 1, end - start - 1), ",");
      button b{};
      for (const auto& bs : button_strings) { b.emplace_back(std::stoul(bs)); }
      buttons.emplace_back(b);
    }
  }
  return {goal, buttons};
}

struct state {
  state* prev = nullptr;
  vector<bool> current;
  vector<state> next = {};

  state(usize len) { current = vector<bool>(len, false); }
  state(vector<bool> v, state* prev = nullptr) : current(std::move(v)), prev(prev) {}

  void populate_next_states(const vector<button>& buttons) {
    if (!next.empty()) return;
    next.reserve(buttons.size());
    for (const auto& button : buttons) {
      auto new_current{current};
      for (auto light : button) { new_current[light] = !new_current[light]; }
      next.emplace_back(std::move(new_current), this);
    }
  }

  friend std::ostream& operator<<(std::ostream& os, const state& s) {
    os << "[";
    for (auto l : s.current) { os << (l ? '#' : '.'); }
    os << "] \n";

    for (const auto& n : s.next) { os << " - " << n << std::endl; }
    return os;
  }
};

u64 bfs_state(state& root, const machine& m) {
  std::deque<state*> q{};
  q.push_back(&root);

  while (!q.empty()) {
    state* node{q.front()};
    q.pop_front();
    if (node->current == m.goal) {
      usize dist{};
      for (state* cur = node; cur->prev != nullptr; cur = cur->prev) dist++;
      return dist;
    }

    node->populate_next_states(m.buttons);
    for (auto& branch : node->next) { q.push_back(&branch); }
  }

  std::printf("Something went wrong, bfs search queue is empty\n");
  std::abort();
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
  u64 sum{};
  string line{};
  while (std::getline(input, line)) {
    // do something
    machine m = parse_machine(line);
    state s(m.goal.size());
    sum += bfs_state(s, m);
  }
  std::cout << sum << std::endl;
  return 0;
};
