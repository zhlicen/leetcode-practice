package leetcode

func addStrings(num1 string, num2 string) string {
    l1 := len(num1)
    l2 := len(num2)
    l := l1
    if l2 > l1 {
        l = l2
    } 
    sum := make([]byte, l)
    add := false
    i := l1 - 1 
    j := l2 - 1
    idx := l - 1
    s := byte(0)
    for {
        if i < 0 && j < 0 {
            return string(sum)
        }
        if i >= 0 {
            s += (num1[i] - '0')
            i--
        } 
        if j >= 0 {
            s += (num2[j] - '0')
            j--
        }
        if add {
            s += 1
            add = false
        }
        if s >= 10 {
            add = true
            s -= 10
        }
        sum[idx] = s + '0'
        idx--
        if  i < 0 && j < 0 && add {
            return "1" + string(sum)   
        }
        s = byte(0)
    }
    return string(sum)    
}