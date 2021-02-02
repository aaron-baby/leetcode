public class StrStr {
    public int strStr(String haystack, String needle) {
        int m = haystack.length(), n = needle.length();
        if (n == 0) {
            return 0;
        }
        for (int i = 0; i <= m - n; i++) {
            int j = 0;
            // move cursor to first match
            if (haystack.charAt(i) != needle.charAt(j)) {
                continue;
            }
            for (; j <= n - 1; j++) {
                if (haystack.charAt(i + j) != needle.charAt(j)) {
                    break;
                }
            }
            if (j == n) {
                return i;
            }
        }
        return -1;
    }
}
