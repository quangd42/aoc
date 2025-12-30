// fast_bfs.cpp
// Build: g++ -O3 -march=native -DNDEBUG -std=c++20 fast_bfs.cpp -o fast_bfs
#include <array>
#include <charconv>
#include <cstdint>
#include <fstream>
#include <iostream>
#include <string>
#include <string_view>

static constexpr uint8_t MAX_N = 10;       // goal.size() <= 10
static constexpr uint8_t MAX_BUTTONS = 16; // support up to 16 buttons

struct Parsed {
  uint16_t goal = 0;                         // bit i is 1 if goal[i] == '#'
  std::array<uint16_t, MAX_BUTTONS> masks{}; // button toggle masks
  uint8_t nb = 0;                            // number of buttons
  uint8_t n = 0;                             // goal length (<= 10)
};

static inline Parsed parse_machine(std::string_view line) {
  Parsed p{};

  // Parse goal between [ ... ]
  const size_t l = line.find('[');
  const size_t r = line.find(']', l + 1);
  if (l == std::string_view::npos || r == std::string_view::npos || r <= l + 1) {
    std::cerr << "Invalid line (missing goal): " << std::string(line) << "\n";
    std::abort();
  }

  const size_t len = r - l - 1;
  if (len > MAX_N) {
    std::cerr << "goal.size() > " << unsigned(MAX_N) << " not supported: " << len << "\n";
    std::abort();
  }
  p.n = static_cast<uint8_t>(len);

  for (uint8_t i = 0; i < p.n; ++i) {
    const char c = line[l + 1 + i];
    if (c == '#')
      p.goal |= (uint16_t(1) << i);
    else if (c == '.') { /* no-op */
    } else {
      std::cerr << "Invalid goal char: '" << c << "' in line: " << std::string(line) << "\n";
      std::abort();
    }
  }

  // Parse each button group "(...)" containing comma-separated indices.
  size_t pos = r + 1;
  while (true) {
    const size_t a = line.find('(', pos);
    if (a == std::string_view::npos) break;
    const size_t b = line.find(')', a + 1);
    if (b == std::string_view::npos) {
      std::cerr << "Invalid line (missing ')'): " << std::string(line) << "\n";
      std::abort();
    }
    if (p.nb >= MAX_BUTTONS) {
      std::cerr << "buttons > " << unsigned(MAX_BUTTONS) << " not supported\n";
      std::abort();
    }

    uint16_t mask = 0;
    size_t cur = a;

    while (cur < b) {
      // Skip spaces and commas
      while (cur < b && (line[cur] == ' ' || line[cur] == ',')) ++cur;
      if (cur >= b) break;

      unsigned idx = 0;
      const char* begin = line.data() + cur;
      const char* end = line.data() + b;
      auto res = std::from_chars(begin, end, idx);
      if (res.ec != std::errc()) {
        std::cerr << "Invalid button index in line: " << std::string(line) << "\n";
        std::abort();
      }
      cur = static_cast<size_t>(res.ptr - line.data());

      if (idx >= p.n) {
        std::cerr << "Button index out of range (" << idx << " >= " << unsigned(p.n)
                  << ") in line: " << std::string(line) << "\n";
        std::abort();
      }

      // XOR is robust even if duplicates appear in a button definition
      mask ^= (uint16_t(1) << idx);
    }

    p.masks[p.nb++] = mask;
    pos = b + 1;
  }

  return p;
}

static inline uint16_t bfs(const Parsed& p) {
  // State space for n<=10 is <=1024. Queue capacity 1024 is always safe.
  std::array<int16_t, 1024> dist;
  dist.fill(-1);

  std::array<uint16_t, 1024> q;
  uint16_t head = 0, tail = 0;

  dist[0] = 0;
  q[tail++] = 0;

  while (head != tail) {
    const uint16_t s = q[head++];
    if (s == p.goal) return static_cast<uint16_t>(dist[s]);
    for (uint8_t i = 0; i < p.nb; ++i) {
      const uint16_t t = static_cast<uint16_t>(s ^ p.masks[i]);
      if (dist[t] == -1) {
        dist[t] = dist[s] + 1;
        q[tail++] = t;
      }
    }
  }

  return 0xFFFF; // unreachable
}

int main(int argc, char* argv[]) {
  if (argc < 2) {
    std::cerr << "missing input argv\n";
    return 1;
  }
  std::ifstream input(argv[1]);
  if (!input) {
    std::cerr << "input file cannot be opened\n";
    return 1;
  }

  uint64_t sum = 0;
  std::string line;
  while (std::getline(input, line)) {
    if (line.empty()) continue;
    const Parsed p = parse_machine(line);
    const uint16_t steps = bfs(p);
    if (steps == 0xFFFF) {
      std::cerr << "Unreachable goal for line: " << line << "\n";
      return 2;
    }
    sum += steps;
  }

  std::cout << sum << "\n";
  return 0;
}
