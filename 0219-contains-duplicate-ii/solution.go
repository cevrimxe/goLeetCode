func containsNearbyDuplicate(nums []int, k int) bool {
    // |i-j| <=k 
    
    window := map[int]struct{}{} // struct{} = 0 byte


    for right:= 0; right< len(nums); right++{

        // o eleman varsa true dön
        if _, ok:= window[nums[right]]; ok{
            return true
        }

        // yoksa ekle
        window[nums[right]] = struct{}{}

        // window boyutu k'ye büyük eşit olursa windowa ilk ekleneni çıkar
        if right >= k {
            delete(window, nums[right-k])
        }
    }

    return false
}
