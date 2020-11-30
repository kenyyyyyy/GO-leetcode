**530. 二叉搜索树的最小绝对差**

[leetcode题目地址](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/)
  
给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。

1. 因为是一颗搜索树,所以根据中序遍历的结果来看是一个升序的数组，所以答案只会在中序遍历结果的相邻两个元素中产生。  