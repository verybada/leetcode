class Solution:  
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        stock = None
        profit = 0
        max_len = len(prices)
        
        for index in range(max_len):
            price = prices[index]
            try:
                next_price = prices[index+1]
            except IndexError:
                next_price = None
        
            # find highest and lowest price
            highest = False
            if next_price is not None and next_price < price:
                highest = True

            lowest = False
            if next_price is not None and next_price > price:
                lowest = True
            
            # sell stock on high
            if stock is not None and highest:
                profit += price - stock
                stock = None
            # buy stock on low
            elif stock is None and lowest:
                stock = price
        else:
            # sell stock if we have
            if stock is not None and price > stock:
                profit += price - stock
        
        return profit
