import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

public class SolutionTest {
    @Test
    void levelOrder() {
        TreeNode r = new TreeNode(3);
        r.left = new TreeNode(9);
        TreeNode rs = new TreeNode(20);
        r.right = rs;
        rs.left = new TreeNode(15);
        rs.right = new TreeNode(7);
        Solution s = new Solution();
        List<List<Integer>> result = s.levelOrder(r);
        List<List<Integer>> want = new ArrayList<>();
        want.add(Arrays.asList(3));
        want.add(Arrays.asList(9,20));
        want.add(Arrays.asList(15,7));
        assertEquals(result, want);
    }
}