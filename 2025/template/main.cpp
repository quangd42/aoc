#include "../common/common.h"
#include <fstream>
#include <iostream>

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
  while (std::getline(input, line)) {
    // do something
  }
  std::cout << sum << std::endl;
  return 0;
};
