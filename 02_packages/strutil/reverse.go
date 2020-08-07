package strutil


// function, which takes a string as 
// argument and return the reverse of string. 
func Reverse(str string) (result string) { 
    for _, v := range str { 
        result = string(v) + result 
    } 
    return
} 