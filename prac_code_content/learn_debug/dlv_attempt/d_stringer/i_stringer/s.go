/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/10/9 23:23 10月
 **/
package i_stringer

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

