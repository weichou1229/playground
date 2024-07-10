use std::collections::HashMap;

pub struct Solution {}

impl Solution {
    pub fn remove_element(nums: &mut Vec<i32>, val: i32) -> i32 {
        if nums.len() == 0 {
            return 0;
        }
        if nums.len() == 1 && nums[0]== val {
            return 0;
        }
        if nums.len() == 1 && nums[0]!= val {
            return 1;
        }
        nums.retain(|&x| x != val);
        nums.len() as i32
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_1() {
        assert_eq!(2, Solution::remove_element(&mut vec![3, 2, 2, 3], 3));
        assert_eq!(5, Solution::remove_element(&mut vec![0,1,2,2,3,0,4,2], 2));
    }
}
