package com.xie.sword2offer;

/**
 * @author xie4ever
 * @date 2021/5/25 0:18
 */
public class JZ1 {

    public boolean Find(int target, int[][] array) {
        for (int i = 0; i < array.length; i++) {
            int[] tmp = array[i];
            for (int j = 0; j < tmp.length; j++) {
                if (target < tmp[j]) {
                    break;
                }
                if (target == tmp[j]) {
                    return true;
                }
            }
        }
        return false;
    }
}
