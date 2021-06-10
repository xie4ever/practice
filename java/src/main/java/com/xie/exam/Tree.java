package com.xie.exam;

/**
 * @author xie4ever
 * @date 2021/5/20 14:37
 */
public class Tree {

    class TreeNode {
        public TreeNode(int value) {
            this.value = value;
        }

        int value;
        TreeNode left;
        TreeNode right;
    }

    @org.junit.jupiter.api.Test
    public void Test() {
        int sum = 3;

        TreeNode root = new TreeNode(1);
        TreeNode node1 = new TreeNode(2);

        root.left = node1;

        System.out.println(find(root, sum));
    }

    public boolean find(TreeNode node, int value) {
        if (node == null) {
            return false;
        }
        int tmp = value - node.value;
        if (tmp < 0) {
            return false;
        }
        if (tmp == 0 && node.left == null && node.right == null) {
            return true;
        }
        return find(node.left, tmp) || find(node.right, tmp);
    }
}
