/*
 * build:
 *   c++ -std=c++11 303.cpp -o 303
 */
#include <iostream>
#include <vector>

class NumArray {
private:
    std::vector<int> subtotal;
public:
    NumArray(std::vector<int> &nums) {
        int n = 0;
        for (auto i: nums) {
            n += i;
            subtotal.push_back(n);
        }
    }

    int sumRange(int i, int j) {
        return (i == 0) ? subtotal[j] : subtotal[j]-subtotal[i-1];
    }
};

int main() {
    std::vector<int> v = {-2, 0, 3, -5, 2, -1};
    NumArray numArray(v);
    std::cout << numArray.sumRange(0, 2) << std::endl; // -> 1
    std::cout << numArray.sumRange(2, 5) << std::endl; // -> -1
    std::cout << numArray.sumRange(0, 5) << std::endl; // -> -3
}
