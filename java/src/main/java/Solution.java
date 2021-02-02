import java.util.*;

public class Solution {
    public List<List<Integer>> levelOrder(TreeNode root) {
        List<List<Integer>> res = new ArrayList<>();
        if (root == null)
            return res;
        Queue<TreeNode> q = new LinkedList<>();
        q.add(root);
        while (!q.isEmpty()) {
            List<Integer> v = new ArrayList<>();
            for (int i = q.size(); i > 0; --i) {
                TreeNode t = q.poll();
                v.add(t.val);
                if (t.left!=null) q.add(t.left);
                if (t.right!=null) q.add(t.right);
            }
            res.add(v);
        }
        return res;
    }
}