#include "../common/common.h"
#include <algorithm>
#include <cstdio>
#include <cstdlib>
#include <fstream>
#include <iostream>
#include <queue>

struct point {
  i32 x, y;

  bool operator<(const point& other) const {
    if (x != other.x) return x < other.x;
    return y < other.y;
  }
};

void print_point(const point& p) { db_print("(%d, %d)\n", p.x, p.y); }

u64 solve_part1(const vector<point>& points) {
  u64 out{};
  auto len{points.size()};
  for (auto i{0}; i + 1 < len; i++) {
    for (auto j{i + 1}; j < len; j++) {
      point a = points[i];
      point b = points[j];
      u64 w = std::llabs(a.x - b.x) + 1;
      u64 h = std::llabs(a.y - b.y) + 1;
      u64 area = w * h;
      if (area > out) { out = area; }
    }
  }
  return out;
}

struct line {
  point a, b;

  // in this problem a line can only be vertical or horizontal
  // assumption: a != b
  bool is_vertical() const { return a.x == b.x; }
  bool is_horizontal() const { return !is_vertical(); }
};

void print_line(const line& l) { db_print("(%d, %d) -> (%d, %d)\n", l.a.x, l.a.y, l.b.x, l.b.y); }

bool is_strictly_crossing(const line& l, const line& k) {
  if (l.is_vertical() == k.is_vertical()) return false;

  const line& v = (l.is_vertical()) ? l : k;
  const line& h = (l.is_horizontal()) ? l : k;

  auto [v_min_y, v_max_y] = std::minmax(v.a.y, v.b.y);
  auto [h_min_x, h_max_x] = std::minmax(h.a.x, h.b.x);

  // Intersection is strict if the fixed coordinate of one line
  // is strictly between the endpoints of the other.
  return (v.a.x > h_min_x and v.a.x < h_max_x and h.a.y > v_min_y and h.a.y < v_max_y);
}

enum class direction {
  left,
  right,
};

direction turn_from_point(point a, point b, point c) {
  i64 cross = (i64)(b.x - a.x) * (i64)(c.y - b.y) - (i64)(b.y - a.y) * (i64)(c.x - b.x);
  if (cross > 0) { return direction::left; }
  if (cross < 0) { return direction::right; }
  std::printf("input contains three straight points\n");
  std::abort();
}

point find_outer_b(point a, point b, direction dir) {
  if (line{a, b}.is_vertical()) {
    if (a.y < b.y) {
      // down
      return point{b.x + 1, dir == direction::left ? b.y + 1 : b.y - 1};
    } else {
      // up
      return point{b.x - 1, dir == direction::left ? b.y - 1 : b.y + 1};
    }
  } else {
    // horizontal
    if (a.x < b.x) {
      // right
      return point{dir == direction::left ? b.x + 1 : b.x - 1, b.y - 1};
    } else {
      // left
      return point{dir == direction::left ? b.x - 1 : b.x + 1, b.y + 1};
    }
  }
}

vector<line> polygon_boundary(const vector<point>& points) {
  auto n{points.size()};
  vector<point> outer_points{};
  outer_points.reserve(n);
  for (auto i{0}; i < n; i++) {
    auto a = points[i];
    auto b = points[(i + 1) % n];
    auto c = points[(i + 2) % n];
    direction dir = turn_from_point(a, b, c);
    point outer_b = find_outer_b(a, b, dir);
    outer_points.emplace_back(outer_b);
  }
  vector<line> outer_lines{};
  outer_lines.reserve(n);
  for (auto i{0}; i < n; i++) {
    outer_lines.emplace_back(line{outer_points[i], outer_points[(i + 1) % n]});
  }
  return outer_lines;
}

vector<line> rectangle_borders(point a, point c) {
  point b{c.x, a.y};
  point d{a.x, c.y};
  return {line{a, b}, line{b, c}, line{c, d}, line{d, a}};
}

struct candidate {
  point a, c;
  u64 area;

  bool operator<(const candidate& other) const { return area < other.area; }
};

u64 solve_part2(const vector<point>& points) {
  auto boundaries{polygon_boundary(points)};
  std::priority_queue<candidate> pq{};

  auto len{points.size()};
  for (auto i{0}; i + 1 < len; i++) {
    for (auto j{i + 1}; j < len; j++) {
      const point& p = points[i];
      const point& q = points[j];
      u64 w = std::llabs(p.x - q.x) + 1;
      u64 h = std::llabs(p.y - q.y) + 1;
      pq.emplace(candidate{p, q, w * h});
    }
  }

  while (!pq.empty()) {
    bool valid = true;
    auto cand = pq.top();
    pq.pop();
    auto cand_borders{rectangle_borders(cand.a, cand.c)};
    for (const auto& l : cand_borders) {
      for (const auto& k : boundaries) {
        if (is_strictly_crossing(l, k)) {
          valid = false;
          break;
        }
      }
      if (!valid) break;
    }
    if (valid) return cand.area;
  }
  std::printf("something went wrong");
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
  string line{};
  vector<point> red_points{};
  while (std::getline(input, line)) {
    line = trim(line);
    auto coord_s{split(line, ",")};
    if (coord_s.size() < 2) {
      std::printf("line missing value pair x, y\n");
      std::abort();
    }
    i32 x = std::stoll(coord_s[0]);
    i32 y = std::stoll(coord_s[1]);
    red_points.emplace_back(point{x, y});
  }

  std::cout << solve_part2(red_points) << std::endl;
  return 0;
};
