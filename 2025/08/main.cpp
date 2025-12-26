#include "../common/common.h"
#include <cassert>
#include <cstddef>
#include <fstream>
#include <iostream>
#include <queue>
#include <unordered_map>

struct vertex {
  u64 x, y, z;

  bool operator==(const vertex& u) const { return x == u.x and y == u.y and z == u.z; }
  bool operator!=(const vertex& u) const { return not operator==(u); }
};

namespace std {
template <> struct hash<vertex> {
  size_t operator()(const vertex& u) const { return (u.x << 42) | (u.y << 21) | u.z; }
};
} // namespace std

struct edge {
  vertex v, u;
  u64 weight;

  edge(vertex v, vertex u)
      : v(v), u(u),
        weight((v.x - u.x) * (v.x - u.x) + (v.y - u.y) * (v.y - u.y) + (v.z - u.z) * (v.z - u.z)) {}
  bool operator<(const edge& other) const { return weight < other.weight; }
  bool operator>(const edge& other) const { return weight > other.weight; }
};

struct set_data {
  vertex parent;
  u64 rank;
};

struct disjoint_set {
  std::unordered_map<vertex, set_data> nodes;

  disjoint_set(const vector<vertex>& vertices) {
    nodes.reserve(vertices.size());
    for (const auto& v : vertices) {
      nodes[v] = {v, 0};
    }
  }

  auto find(const vertex& v) {
    auto& sd{nodes.at(v)};
    auto& parent = sd.parent;
    if (parent == v) {
      return v;
    }
    sd.parent = find(parent);
    return parent;
  }

  bool union_by_rank(const vertex& v, const vertex& u) {
    vertex root_v = find(v);
    vertex root_u = find(u);

    if (root_v != root_u) {
      auto& data_v = nodes.at(root_v);
      auto& data_u = nodes.at(root_u);

      if (data_v.rank < data_u.rank) {
        data_v.parent = root_u;
      } else if (data_v.rank > data_u.rank) {
        data_u.parent = root_v;
      } else {
        data_u.parent = root_v;
        data_v.rank += 1;
      }
      return true;
    }
    return false;
  }
};

void print_vertex(const vertex& v) { std::cout << "(" << v.x << ", " << v.y << ", " << v.z << ")"; }

void print_edge(const edge& e) {
  print_vertex(e.v);
  std::cout << " -> ";
  print_vertex(e.u);
  std::cout << " : " << e.weight << "\n";
}

void print_ds(const disjoint_set& ds) {
  for (const auto& [v, d] : ds.nodes) {
    std::cout << "Vertex: ";
    print_vertex(v);
    std::cout << "; Parent: ";
    print_vertex(d.parent);
    std::cout << std::endl;
  }
  std::cout << std::endl;
}

void print_counter(const std::unordered_map<vertex, u64>& counter) {
  for (const auto& [v, c] : counter) {
    std::cout << "Vertex: ";
    print_vertex(v);
    std::cout << "; Count: " << c << std::endl;
  }
  std::cout << std::endl;
};

u64 solve_part1(const vector<vertex>& vs, const vector<edge>& edges) {
  disjoint_set ds(vs);
  std::priority_queue<edge, vector<edge>, std::greater<edge>> pq(edges.begin(), edges.end());

  for (auto i{0}; i < 1000; i++) {
    edge e = pq.top();
    // std::cout << "size: " << pq.size() << ", i = ";
    // std::cout << i << ": ";
    // print_edge(e);
    ds.union_by_rank(e.v, e.u);
    pq.pop();
  }
  // print_ds(ds);

  std::unordered_map<vertex, u64> counter{};
  for (const auto& [v, d] : ds.nodes) {
    auto parent_v = d.parent;
    auto root_v = ds.find(parent_v);
    // while (root_v != parent_v) {
    //   root_v = ds.find(parent_v);
    // }
    if (auto count = counter.find(root_v); count != counter.end()) {
      count->second += 1;
      // std::cout << count->second << "\n";
    } else {
      counter.insert({root_v, 1});
    }
  }

  u64 first = 0, second = 0, third = 0;
  for (const auto& [v, c] : counter) {
    // std::cout << c << "\n";
    if (c > first) {
      third = second;
      second = first;
      first = c;
    } else if (c > second) {
      third = second;
      second = c;
    } else if (c > third) {
      third = c;
    }
  }

  return first * second * third;
}

u64 solve_part2(const vector<vertex>& vs, const vector<edge>& edges) {
  disjoint_set ds(vs);
  std::priority_queue<edge, vector<edge>, std::greater<edge>> pq(edges.begin(), edges.end());

  vertex prev_v, prev_u;
  u64 mst{};
  auto vs_count = vs.size();
  while (mst < vs_count - 1) {
    edge e = pq.top();
    if (ds.union_by_rank(e.v, e.u)) {
      mst++;
    }
    prev_v = e.v;
    prev_u = e.u;
    pq.pop();
  }

  return prev_v.x * prev_u.x;
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
  u64 sum{};
  vector<vertex> vertices{};
  while (std::getline(input, line)) {
    line = trim(line);
    auto coord_s{split(line, ",")};
    assert(coord_s.size() == 3);
    u64 x = std::stoull(coord_s[0]);
    u64 y = std::stoull(coord_s[1]);
    u64 z = std::stoull(coord_s[2]);
    vertices.emplace_back(vertex{x, y, z});
  }

  vector<edge> edges{};
  for (auto i{0}; i < vertices.size() - 1; i++) {
    for (auto j{i + 1}; j < vertices.size(); j++) {
      edge e{vertices.at(i), vertices.at(j)};
      edges.emplace_back(e);
    }
  }

  sum = solve_part2(vertices, edges);

  std::cout << sum << std::endl;
  return 0;
};
