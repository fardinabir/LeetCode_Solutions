class Solution {
public:
    struct Job {
    int startTime;
    int endTime;
    int profit;

    // Constructor for easy initialization
    Job(int start, int end, int p) : startTime(start), endTime(end), profit(p) {}
};

    int dp[100005]; // Added a semicolon to end the line

    int jobScheduling(std::vector<int>& startTime, std::vector<int>& endTime, std::vector<int>& profit) {
        memset(dp, -1, sizeof(dp)); // Added "sizeof(dp)" for array size
        std::vector<Job> jobs;
    for (size_t i = 0; i < startTime.size(); ++i) {
        jobs.emplace_back(startTime[i], endTime[i], profit[i]);
    }

    // Sort the vector of jobs based on startTime
    std::sort(jobs.begin(), jobs.end(), [](const Job& a, const Job& b) {
        return a.startTime < b.startTime;
    });

    for (int i=0;i<jobs.size();i++) {
        startTime[i] = jobs[i].startTime;
        endTime[i] = jobs[i].endTime;
        profit[i] = jobs[i].profit;
    }
    return findMaxProfit(startTime, endTime, profit, 0);
    }
    
    int findMaxProfit(std::vector<int>& startTime, std::vector<int>& endTime, std::vector<int>& profit, int ind) {
        if (dp[ind] != -1) {
            return dp[ind];
        }
        if (ind >= startTime.size()) {
            return 0;
        }
        int path1, path2, p;
        path1 = findMaxProfit(startTime, endTime, profit, ind + 1);
        
        p = binarySearchUpperBound(startTime, endTime[ind]);
        path2 = findMaxProfit(startTime, endTime, profit, p) + profit[ind];
        // cout<<"*****Profit : "<<profit[ind]<<" path1: "<<path1<<" path2: "<<path2<<" p: "<<p<<endl;
        dp[ind] = max(path1, path2);
        return dp[ind];
    }
    
    
    int binarySearchUpperBound(std::vector<int>& numbers, int targetValue) {
        auto lowerBoundIt = std::lower_bound(numbers.begin(), numbers.end(), targetValue);
        auto upperBoundIt = std::upper_bound(numbers.begin(), numbers.end(), targetValue);

        if (lowerBoundIt != upperBoundIt) {
            return distance(numbers.begin(), lowerBoundIt);
        } else {
            return distance(numbers.begin(), upperBoundIt);
        }
    }
};