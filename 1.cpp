/*
 * build:
 *   c++ -std=c++11 1.cpp -o 1
 */
#include <iostream>
#include <vector>
#include <map>

class Solution {
public:
    std::vector<int> twoSum(std::vector<int>& nums, int target) {
        std::map<int, int> m;

        int size = nums.size();
        for (int i = 0; i < size; i++) {
            m[nums[i]] = i;
        }
        for (int i = 0; i < size; i++) {
            std::map<int, int>::iterator it;
            it = m.find(target-nums[i]);
            if (it != m.end() && i != it->second) {
                return std::vector<int>{i,it->second};
            }
        }
        return std::vector<int>();
    }
};

int main() {
    std::vector<int> v = {2, 7, 11, 15}, ret;
    Solution sol;
    ret = sol.twoSum(v, 9);
    if (ret.size() == 2) {
        std::cout << "[" << ret[0] << "," << ret[1] << "]" << std::endl; // -> ,0,1]
    }
}

