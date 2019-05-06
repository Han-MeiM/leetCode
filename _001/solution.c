#include <stdio.h>
#include <stdlib.h>

int *twoSum(int* nums, int numsSize, int target, int* returnSize)
{
    *returnSize = 2;
    static int a[2]={0};
	for (int i = 0; i < numsSize - 1; i++)
	{
		for (int j = i+1; j < numsSize; j++)
		{
			if (nums[i] + nums[j] == target)
			{
				a[0] = i;
				a[1] = j;
				return a;
			}
		}
	}
	return 0;
}

int main(void)
{
    int nums[] = { 3, 2, 3 };
    int count = sizeof(nums) / sizeof(*nums);
    int target = 6;
    int returnSize;
    int *indexes = twoSum(nums, count, target, &returnSize);
    if (indexes != NULL) {
        printf("%d %d\n", indexes[0], indexes[1]);
    } else {
        printf("Not found\n");
    }

    return 0;
}