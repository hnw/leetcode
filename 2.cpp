/*
 * build:
 *   c++ -std=c++11 2.cpp -o 2
 */
#include <iostream>

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *n) : val(x), next(n) {}
};

std::ostream & operator<<(std::ostream & stream, ListNode* const & l) {
    if (l != nullptr) {
        stream << l->val;
        if (l->next != nullptr) {
            stream << " -> " << l->next;
        }
    }
    return stream;
}

class Solution {
private:
    ListNode* addTwoNumbersAux(ListNode* l1, ListNode* l2, int carry) {
        if (carry == 0) {
            if (l1 == nullptr) { return l2; }
            if (l2 == nullptr) { return l1; }
        }

        int v = carry;
        if (l1 != nullptr) {
            v += l1->val;
            l1 = l1->next;
        }
        if (l2 != nullptr) {
            v += l2->val;
            l2 = l2->next;
        }
        ListNode *ret = new ListNode(v % 10);
        ret->next = addTwoNumbersAux(l1, l2, v/10);
        return ret;
    }
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        return addTwoNumbersAux(l1, l2, 0);
    }
};

int main() {
    ListNode *l1 = new ListNode(2, new ListNode(4, new ListNode(3)));
    ListNode *l2 = new ListNode(5, new ListNode(6, new ListNode(4)));
    Solution sol;

    std::cout << sol.addTwoNumbers(l1,l2) << std::endl;
}
