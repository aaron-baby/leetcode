/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */
import {TreeNode} from "./tree_node";
export function findMode(root: TreeNode | null): number[] {
    let o:{
        max: number
        cnt: number
    } = {
        max:0,
        cnt:1
    };
    let pre:TreeNode[] = [];
    let res:number[] = [];
    inOrder(root, pre,o,res);
    return res;
}

function inOrder(node: TreeNode|null, pre: TreeNode[], o:{ max: number; cnt: number }, res: number[]): number[]|undefined{
    if (node==null) {
        return;
    }
    inOrder(node.left, pre, o, res);
    let preItem = pre.pop()

    if (preItem!=undefined) {
        o.cnt = (node.val == preItem.val) ? o.cnt + 1 : 1;
    }
    if (o.cnt >= o.max) {
        if (o.cnt>o.max) {
            res.length = 0
        }
        res.push(node.val)
        o.max = o.cnt
    }
    pre.push(node)
    inOrder(node.right, pre, o, res);
}