package main
import "fmt"
func twosum(nums []int,target int)[]int{
	var NUMS []int
	for i:=0,i<len(nums)；i++{
		for j:=i+1;j<len(nums);j++{
			if(nums[i]+nums[j]==target){
				NUMS[0]=i
				NUMS[1]=j
			}
			}
		}
		}
		return NUMS 
	}
func main(){
	var target,length int
	var nums,NUMS []int
	fmt.Println("请输入数组长度：")
	fmt.Scanf("%d",&length)
	fmt.Println("请输入数组：")
	for k:=0;k<length;k++{
		fmt.Scanf("%d",&nums[k])
	}
	fmt.Println("请输入两数和：")
	fmt.Scanf("%d",&target)
	NUMS=twosum(nums,target)
	if(len(NUMS)!=0){
		fmt.Println(NUMS[ : ])
	}
}


