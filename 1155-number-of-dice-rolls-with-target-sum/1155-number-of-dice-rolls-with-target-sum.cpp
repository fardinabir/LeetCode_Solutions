class Solution {
public:
    int numRollsToTarget(int n, int k, int target) {
        const int MOD = 1000000007;

        if (n * k < target) {
            return 0;
        }

        vector<vector<long long>> dp(n + 1, vector<long long>(target + 1, 0));

        dp[0][0] = 1;

        for (int i = 1; i <= n; i++) {
            for (int j = i; j <= min(i * k, target); j++) {
                for (int t = 1; t <= min(k, j); t++) {
                    dp[i][j] = (dp[i][j] + dp[i - 1][j - t]) % MOD;
                }
            }
        }

        return static_cast<int>(dp[n][target]);
    }
};
