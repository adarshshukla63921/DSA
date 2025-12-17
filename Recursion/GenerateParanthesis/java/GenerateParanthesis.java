import java.util.*;

class GenerateParanthesis{
    public static void generateParanthesis(
        int idx,
        String s,
        int open,
        int close,
        List<String> list,
        int n
    ){
        if(open > n) return;

        if(open==close && open+close==2*n){
            list.add(s);
            return;
        }
        generateParanthesis(idx+1, s.concat("("), open+1, close, list, n);
        if(open>close){
            generateParanthesis(idx+1, s.concat(")"), open, close+1, list, n);
        }
    }
    public static void main(String[] args) {
        List<String> list = new ArrayList<>();
        generateParanthesis(0,"",0,0,list,3);
        System.out.println(list);
    }
}