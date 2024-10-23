
#题解
定义函数twosum:
```go
func twosum(nums []int,target int)[]int{
var NUMS []int
for i:=0,i<len(nums);i++{
for j:=i+1;j<len(nums);j++{
if(nums[i]+nums[j]==target){
NUMS={i,j}
}
}
}
return NUMS
}
```
再调用twosum函数


