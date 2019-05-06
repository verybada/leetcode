# Definition for an interval.
#class Interval(object):
#    def __init__(self, s=0, e=0):
#        self.start = s
#        self.end = e

class Solution(object):
    def insert(self, intervals, newInterval):
        """
        :type intervals: List[Interval]
        :type newInterval: Interval
        :rtype: List[Interval]
        """
        result = list()
        tmp = newInterval
        for i in intervals:
            if tmp.end < i.start:
                result.append(tmp)
                tmp = i
            elif tmp.start > i.end:
                result.append(i)
            else:
                tmp.start = tmp.start if tmp.start < i.start else i.start
                tmp.end = tmp.end if tmp.end > i.end else i.end
    
        result.append(tmp)       
        return result
