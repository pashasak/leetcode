#!/bin/python3
#
#   Array    Two Pointers    Sorting
#
#   Two Sum, 3Sum Closest, 4Sum
#
from typing import List, Tuple, Set

class Solution:
    def threeSum(self, nums:List[int]) -> List[List[int]]:
        nums.sort()
        output = set()
        for i, v in enumerate(nums):
            output = output | self.twoSum(nums[i+1:], -v)
        return list(map(list, output))

    def twoSum(self, nums:List[int], target) -> Set[Tuple[int]]:
        seen = {}
        output = set()
        for i, v in enumerate(nums):
            r = target - v
            if r in seen:
                output.add((v, r, -target))
            else:
                seen[v] = i
        return output



if __name__ == '__main__':
    print(Solution().threeSum([-1, 0, 1, 2, -1, 4]))

    #print(Solution().longestPalindrome('cbbd'))
