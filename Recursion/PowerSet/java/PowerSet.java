package powerset.java;

import java.util.*;

public class PowerSet {
    public static void generator(
        int idx,
        int[] arr,
        List<Integer> list, 
        List<List<Integer>> ans,
        int n
    ){
        if(idx==n){
            ans.add(new ArrayList<>(list));
            return;
        }
        list.add(arr[idx]);
        generator(idx+1, arr, list, ans, n);
        list.remove(list.size()-1);
        generator(idx+1, arr, list, ans, n);
    }
    public static void main(String[] args) {
        List<Integer> list = new ArrayList<>();
        List<List<Integer>> powerSet = new ArrayList<>();
        int[] arr =  {1,2,3};
        generator(0, arr, list, powerSet, arr.length);
        System.out.println(powerSet);
    }
}
