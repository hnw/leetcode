/*
 * build:
 *   c++ -std=c++11 4.cpp -o 4
 */
#include <iostream>
#include <vector>

class Solution {
private:
    double findMedianSortedArrayAux(std::vector<int>& nums, int medianI, int medianNum) {
        if (medianNum == 1) {
            return (double)nums[medianI];
        } else {
            return (nums[medianI]+nums[medianI+1]) / 2.0;
        }
    }
public:
    double findMedianSortedArrays(std::vector<int>& nums1, std::vector<int>& nums2) {
        double median = 0.0;
        int i1 = 0, i2 = 0;
        int nums1Len = nums1.size();
        int nums2Len = nums2.size();
        int mergedLen = nums1Len + nums2Len;
        int medianI = (mergedLen - 1) / 2;
        int medianNum = (mergedLen%2 == 0) ? 2 : 1;

        if (nums1Len == 0) {
            return findMedianSortedArrayAux(nums2, medianI, medianNum);
        }
        if (nums2Len == 0) {
            return findMedianSortedArrayAux(nums1, medianI, medianNum);
        }

        // skip to median
        for (int i = 0; i < medianI; i++) {
            if (nums1[i1] <= nums2[i2]) {
                i1++;
                if (i1 >= nums1Len) {
                    return findMedianSortedArrayAux(nums2, medianI-nums1Len, medianNum);
                }
            } else {
                i2++;
                if (i2 >= nums2Len) {
                    return findMedianSortedArrayAux(nums1, medianI-nums2Len, medianNum);
                }
            }
        }

        if (nums1[i1] <= nums2[i2]) {
            median = (double)nums1[i1];
            i1++;
            if (i1 >= nums1Len && medianNum == 2) {
                return (median + (double)nums2[i2]) / 2;
            }
        } else {
            median = (double)nums2[i2];
            i2++;
            if (i2 >= nums2Len && medianNum == 2) {
                return (median + (double)nums1[i1]) / 2;
            }
        }

        if (medianNum == 2) {
            if (nums1[i1] <= nums2[i2]) {
                median += (double)nums1[i1];
            } else {
                median += (double)nums2[i2];
            }
            median /= 2.0;
        }
        return median;
    }
};

int main() {
    std::vector<int> nums1 = {1, 2, 3, 4};
    std::vector<int> nums2 = {5, 6, 7};
    Solution sol;

    std::cout << sol.findMedianSortedArrays(nums1, nums2) << std::endl;
}


