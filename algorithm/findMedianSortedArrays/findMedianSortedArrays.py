class Solution(object):
    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        >>> obj = Solution()
        >>> obj.findMedianSortedArrays([], [1])
        1.0
        >>> obj.findMedianSortedArrays([], [2, 3])
        2.5
        >>> obj.findMedianSortedArrays([], [1,2,3,4,5])
        3.0
        """
        # Merge two lists and sort them
        nums1.extend(nums2)
        new_list = sorted(nums1)

        devider = 1.0
        listlen = len(new_list)
        # find the middle position of this list
        middle = listlen >> 1
        remain = listlen % 2 # check list is odd or not
        median = new_list[middle-1+remain]
        if remain == 0:
            # Needs to calculate median with 2 numbers when list is not odd
            median += new_list[middle]
            devider = 2.0

        return median/devider
