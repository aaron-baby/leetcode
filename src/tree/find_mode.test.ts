
import { findMode } from './find_mode';
import {TreeNode} from "./tree_node";
import 'jest-extended'

test('passes when modes match', () => {
    let leaf = new TreeNode(2);
    let r = new TreeNode(2, leaf);
    let t = new TreeNode(1, null, r);
    let want:number[] = [2];
    let got = findMode(t);
    expect(got).toIncludeSameMembers(want);

    let t2 = new TreeNode(2,new TreeNode(1))
    expect(findMode(t2)).toIncludeSameMembers([1,2])
});