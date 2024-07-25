# Merge Intervals
Given an array of intervals where ```intervals[i] = [starti, endi]```, merge all overlapping intervals and return an array of the non-overlapping intervals that cover all the intervals in the input.


Leetcode link: https://leetcode.com/problems/merge-intervals/description/
 
**Examples:**</br>
**Example 1:**</br>
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]</br>
Output: [[1,6],[8,10],[15,18]]</br>
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].

**Example 2:**</br>
Input:intervals = [[1,4],[4,5]]</br>
Output: [[1,5]]</br>
Explanation: Intervals [1,4] and [4,5] are considered overlapping.</br>

**Constraints:**</br>
- ```1 <= intervals.length <= 10^4```
- ```intervals[i].length == 2```
- ```0 <= starti <= endi <= 10^4```
- ```The input array is not guaranteed to be sorted```
 

**Scoring**</br>
- Solutions must pass all test cases, including edge cases; tests used for grading/scoring may not be the same tests provided as examples.
- Fast execution times will be rewarded, but solutions must, first and foremost, be correct and complete.
- If more than three languages are submitted, the fastest solution from each of the top three will be considered.

**Submission**
- All solutions must be submitted either via an MR into this GitHub repo or through email by Aug 3rd at 23:59:59


