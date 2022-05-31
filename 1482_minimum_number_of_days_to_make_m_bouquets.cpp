#include <iostream>
#include <vector>
#include <string>

using namespace std;

class Solution {
    public:
        int minDays(vector<int>& bloomDay, int m, int k){
            int n = bloomDay.size(), left = 1, right = 1e9;
            if (m * k > n) return -1;
            while (left < right){
                int mid = (left + right) / 2;
                int bouquets = 0, flowers = 0;
                for (int daysneed : bloomDay){
                    if (daysneed <= mid){
                        flowers = flowers + 1;
                        if (flowers == k){
                            bouquets = bouquets + 1;
                            flowers = 0;
                        }
                    } 
                    else {
                        flowers = 0;
                    }
                }
                if (bouquets >= m){
                    right = mid;
                }
                else {
                    left = mid + 1;
                }
            }
            return left


        }
};

int main(){
    Solution s;
    vector<int> v {1, 10, 3, 10, 2};
    int output = s.minDays(v, 3, 1);
    cout <<  output << endl;
}