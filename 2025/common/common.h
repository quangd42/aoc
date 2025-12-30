#include <cmath>
#include <cstdint>
#include <string>
#include <vector>

using std::string;
using std::string_view;
using std::vector;

using usize = size_t;
using u8 = uint8_t;
using u16 = uint16_t;
using u32 = uint32_t;
using u64 = uint64_t;
using i32 = int32_t;
using i64 = int64_t;
using f32 = float_t;
using f64 = double_t;
using std::size_t;

inline bool debug_enabled = false;

inline void db_print(const char* format, ...) {
  if (!debug_enabled) return;
  va_list args;
  va_start(args, format);

  std::vprintf(format, args);
  va_end(args);
}

inline auto trim(string_view s) {
  const string ws = " \t\n\r\f\v";
  const size_t first = s.find_first_not_of(ws);
  if (string::npos == first) { return s; }
  const size_t last = s.find_last_not_of(ws);
  return s.substr(first, last + 1 - first);
}

inline auto split(string_view str, string_view delim) {
  vector<string> out{};
  int cbegin_idx{};
  int cend_idx{};

  while (cend_idx != string::npos) {
    cbegin_idx = str.find_first_not_of(delim, cend_idx);
    cend_idx = str.find_first_of(delim, cbegin_idx);
    // apparently when len is negative substr will go to the end of string
    out.emplace_back(str.substr(cbegin_idx, cend_idx - cbegin_idx));
  }
  return out;
}
