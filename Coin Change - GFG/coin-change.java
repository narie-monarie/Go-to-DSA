//{ Driver Code Starts
// Initial Template for Java

import java.io.*;
import java.util.*;
class GfG {
    public static void main(String args[]) {
        Scanner sc = new Scanner(System.in);
        int t = sc.nextInt();
        while (t-- > 0) {
            int sum = sc.nextInt();
            int N = sc.nextInt();
            int coins[] = new int[N];
            for (int i = 0; i < N; i++) coins[i] = sc.nextInt();
            Solution ob = new Solution();
            System.out.println(ob.count(coins, N, sum));
        }
    }
}


// } Driver Code Ends


// User function Template for Java

class Solution {

    public long count(int coins[], int N, int amount) {

        int n= coins.length;

        long [][]dp= new long[n][amount+1];

        for(long arr[]:dp){

            Arrays.fill(arr,-1);

        }

        

        return f(n-1,amount,coins,dp);

        

    }

    

    static long f(int ind,int target,int arr[],long dp[][]){

        

        if(ind==0){

            if(target%arr[0]==0) return 1;

            else return 0;

        }

        if(dp[ind][target]!=-1)return dp[ind][target];

        

        long nottake=f(ind-1,target,arr,dp);

        long take= 0;

        if(arr[ind]<=target){

            take=f(ind,target-arr[ind],arr,dp);

        }

        return dp[ind][target]=take+nottake;

    }

}