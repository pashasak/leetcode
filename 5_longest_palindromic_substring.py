#!/bin/python3

import unittest

class Solution:
    def longestPalindrome(self, s:str) -> str:
        p = ''
        def palindrome(s: str, l: int, r: int) -> str:
            while l >= 0 and r < len(s) and s[l] == s[r]:
                l -= 1
                r += 1
            return s[l+1:r]
        
        for i in range(len(s)):
            p1 = palindrome(s, i, i+1)
            p2 = palindrome(s, i, i)
            p = max([p, p1, p2], key = len)
        return p


if __name__ == '__main__':
    print(Solution().longestPalindrome('babad'))
    print(Solution().longestPalindrome('cbbd'))
