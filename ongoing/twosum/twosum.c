#include <stdio.h>
#include <stdlib.h>

typedef struct buf
{
	int* index[2];
}buf;

int buf_has_value(buf* tmp, int index)
{
	buf* ptr = &tmp[index];
	printf("%p\n", ptr);
	if(ptr->index[0] != NULL)
	{
		return 1;
	}
	return 0;
}

void buf_set_value(buf* tmp, int index, int value)
{
	int i;
	for(i=0; i<2; i++)
	{
		if(tmp[index].index[i] == NULL)
		{
			int* a = malloc(sizeof(int));
			*a = value;
			tmp[index].index[i] = a;
			printf("HASH %d[%d] = %d(%d)\n", index, i, value, *a);
			break;
		}
	}

	return ;
}

int buf_get_value(int* nums, buf* tmp, int index, int value, int target)
{
	buf* ptr = &tmp[index];
	int i;
	for(i=0; i<2; i++)
	{
		int tmp_index = *ptr->index[i];
		int tmp_value = nums[tmp_index];
		if(tmp_value + value == target)
			return i;
	}

	return -1;
}

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* twoSum(int* nums, int numsSize, int target) 
{
	int* ret_array = calloc(sizeof(int), 2);
	buf* hash = calloc(sizeof(buf), 100000);
	
	int i;
	for(i=0; i<numsSize; i++)
	{
		// 2 3 9 11, 11
		// -1 0 1, 0
		// 2 3 4, 6
		// -1 -2 -3, -4
		// hash index is offset
		// value is i
		int number = nums[i];
		int offset = abs(target - number);
		int leak = abs(target - (target-number));
		printf("i=%d target=%d num=%d offset=%d leak=%d\n", i, target, number,offset, leak);
		if(buf_has_value(hash, leak) == 1)
		{
			ret_array[0] = buf_get_value(nums, hash, leak, number, target)+1;
			ret_array[1] = i+1;
			break;
		}

		buf_set_value(hash, offset, i);
	}

	return ret_array;
}

int main()
{
	int input[] = {572,815,387,418,434,530,376,190,196,74,830,561,973,771,640,37,539,369,327,51,623,575,988,44,659,48,22,776,487,873,486,169,499,82,128,31,386,691,553,848,968,874,692,404,463,285,745,631,304,271,40,921,733,56,883,517,99,580,55,81,232,971,561,683,806,994,823,219,315,564,997,976,158,208,851,206,101,989,542,985,940,116,153,47,806,944,337,903,712,138,236,777,630,912,22,140,525,270,997,763,812,597,806,423,869,926,344,494,858,519,389,627,517,964,74,432,730,843,673,985,819,397,607,34,948,648,43,212,950,235,995,76,439,614,203,313,180,760,210,813,920,229,615,730,359,863,678,43,293,978,305,106,797,769,3,700,945,135,430,965,762,479,152,121,935,809,101,271,428,608,8,983,758,662,755,190,632,792,789,174,869,622,885,626,310,128,233,82,223,339,771,741,227,131,85,51,361,343,641,568,922,145,256,177,329,959,991,293,850,858,76,291,134,254,956,971,718,391,336,899,206,642,254,851,274,239,538,418,21,232,706,275,615,568,714,234,567,994,368,54,744,498,380,594,415,286,260,582,522,795,261,437,292,887,405,293,946,678,686,682,501,238,245,380,218,591,722,519,770,359,340,215,151,368,356,795,91,250,413,970,37,941,356,648,594,513,484,364,484,909,292,501,59,982,686,827,461,60,557,178,952,218,634,785,251,290,156,300,711,322,570,820,191,755,429,950,18,917,905,905,126,790,638,94,857,235,889,611,605,203,859,749,874,530,727,764,197,537,951,919,24,341,334,505,796,619,492,295,380,128,533,600,160,51,249,5,837,905,747,505,82,158,687,507,339,575,206,28,29,91,459,118,284,995,544,3,154,89,840,364,682,700,143,173,216,290,733,525,399,574,693,500,189,590,529,972,378,299,461,866,326,43,711,460,426,947,391,536,26,579,304,852,158,621,683,901,237,22,225,59,52,798,262,754,649,504,861,472,480,570,347,891,956,347,31,784,581,668,127,628,962,698,191,313,714,893};
	int* ans = twoSum(input, sizeof(input)/sizeof(int), 101);
	printf("%d %d\n", ans[0], ans[1]);
	free(ans);
	return 0;
}
