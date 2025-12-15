class MyPow{
    public static double f(double x, int n){
        if(n==0) return 1.0;
        if(n<0){
            x = 1/x;
            n = -1*n;
        }
        if(n%2!=0) return x * f(x, n-1);
        return f(x*x, n/2);
    }
    public static void main(String[] args){
        System.out.println("2^10 : "+f(2.000000, 10));
    }
}