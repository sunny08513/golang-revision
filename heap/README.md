# Heap

https://www.geeksforgeeks.org/heap-data-structure/
https://www.geeksforgeeks.org/introduction-to-heap-data-structure-and-algorithm-tutorials/


Parent:- (i-1)/2
Child:- Left(2*i + 1), Right(2*i + 2)

# Priority Queue in c++ 
```
 int findKthLargest(vector<int>& nums, int k) {
          priority_queue<int, vector<int>, //Min Heap 
          greater<int>> pq;
          for(int i= 0;i<nums.size();i++) {
              pq.push(nums[i]);
              if(pq.size() > k) {
                 pq.pop();
              }
          }
          return pq.top();
    }
```