# **题解**

1.定义函数twosum:
```go
func twosum(nums []int,target int)[]int{
  var NUMS []int
  for i:=0,i<len(nums);i++{
  for j:=i+1;j<len(nums);j++{
  if(nums[i]+nums[j]==target){
     NUMS[0]=i
     NUMS[1]=j
     }
   }  
 }
  return NUMS
}
```
2.再调用twosum函数
```go
func main(){
  var target,length int
  var nums,NUMS []int
  fmt.Println("请输入数组：")
  fmt.Scanf("%d",&length)
  for k:=0;k<length;k++{
    fmt.Scanf("%d",&nums[k])
   }     
  fmt.Println("请输入两数和：")
  fmt.Scanf("%d",&target)
  NUMS=twosum(nums,target)
  if(len(NUMS)!=0){
    fmt.Println(NUMS[ : ]
   }
}

