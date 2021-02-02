import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class StrStrTest {
    @Test
    void strStr() {
        String haystack = "hello", needle = "ll";
        StrStr s = new StrStr();
        assertEquals(2, s.strStr(haystack, needle));

        haystack = "aaaaa";
        needle = "bba";
        assertEquals(-1, s.strStr(haystack, needle));

        haystack = "a";
        needle = "a";
        assertEquals(0, s.strStr(haystack, needle));

        haystack = "mississippi";
        needle = "a";
        assertEquals(-1, s.strStr(haystack, needle));

        haystack = "mississippi";
        needle = "issip";
        assertEquals(4, s.strStr(haystack, needle));

        haystack = "mississippi";
        needle = "mississippi";
        assertEquals(0, s.strStr(haystack, needle));

        haystack = "abc";
        needle = "c";
        assertEquals(2, s.strStr(haystack, needle));
    }
}