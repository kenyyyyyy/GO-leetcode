**剑指 Offer 68 - I. 二叉搜索树的最近公共祖先**

[leetcode题目地址](https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xian-lcof/)
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

因为这道题没有Go版本的所以就用C++来写了，因为这颗树是搜索树，所以其实就很简单了，如果p、q在某一个节点的两边或者等于p、q其中一个，那么这个节点就是这两个节点的最近公共祖先。
否则看p、q和当前节点的关系就只需要遍历一边的子树就行了。  

```
struct TreeNode {
     int val;
     TreeNode *left;
     TreeNode *right;
     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 };
 class Solution {
 public:
     TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
         TreeNode* ans = nullptr;
         if(root == nullptr){
             return nullptr;
         }
          // 如果root等于p,q其中一个，就直接返回root
         if (root->val==p->val || root->val==q->val){
             return root;
         }
         // 如果root在p,q中间，就直接返回root
         if (root->val>p->val&&root->val<q->val || root->val>q->val&&root->val<p->val){
             return root;
         }
         // 否则递归遍历左子树或者右子树
         if (root->val>p->val&&root->val>q->val){
             ans=lowestCommonAncestor(root->left,p,q);
         }
         if (root->val<p->val&&root->val<q->val){
             ans=lowestCommonAncestor(root->right,p,q);
         }
         return ans;
     }
 };
```

