package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) (ans int) {
    num := [128]int{}
    left := 0
    for right,y := range s{
        num[y]++
        for num[y] > 1{
            num[s[left]] --
            left++
        }
        ans = max(ans,right-left+1)
    }
    return ans
}