#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int lengthOfLongestSubstring(char *s)
{
    if (!s)
        return 0;
    int strLen = strlen(s);
    int maxLen = 0;
    for (int i = 0, j, pos = 0; i < strLen; ++i)
    {
        for (j = pos; j < i; ++j)
        {
            if (s[i] == s[j])
            {
                if (maxLen < i - pos)
                    maxLen = i - pos;
                pos = j + 1;
                break;
            }
        }
        if (strLen - 1 == i)
            if (maxLen < i - pos + 1)
                maxLen = i - pos + 1;
    }
    return maxLen;
}

int main(void)
{
    char *s = "abcabcbb";
    int result = lengthOfLongestSubstring(s);

    printf("max: %d", result);
}
