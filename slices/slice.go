package slices

import (
	"slices"

	"github.com/floatyun/gocollections/numbers"
)

// ToMap 传入slices和从元素提取key的函数keyF,变成map(key 由keyF提取的，val 原来的slice的元素)
// 通常使用场景 ent查出来的结构体数组，想构造一个 结构体id为key，val是结构体的map
// 例子. s []*ent.User --> m map[int64]*ent.User 其中map的key是ent.User.ID 则用法如下:
// ToMap(s, func(v *ent.User) int64 {return v.ID})
func ToMap[K comparable, V any](s []V, keyF func(v V) K) map[K]V {
	m := make(map[K]V, len(s))
	for _, v := range s {
		m[keyF(v)] = v
	}
	return m
}

func KeyToIndex[K comparable, V any](s []V, keyF func(v V) K) map[K]int {
	m := make(map[K]int, len(s))
	for i, v := range s {
		m[keyF(v)] = i
	}
	return m
}

// ToMapByKVF 传入slices和 kvF, 变成map(key,val 由kvF提取的)
func ToMapByKVF[K comparable, V1 any, V2 any](s []V1, kvF func(v V1) (K, V2)) map[K]V2 {
	m := make(map[K]V2, len(s))
	for i := range s {
		k, v := kvF(s[i])
		m[k] = v
	}
	return m
}

// GroupByKF 将k(key)一致的元素聚合成一个列表。最终形成key到元素列表的映射。
//
//	@Description: 传入slice和kvF.对每个元素使用kF计算map的key值：k。将k(key)一致的元素聚合成一个列表。最终形成map[k][]v。使用场景：多为根据的切片元素的某个字段值，进行元素的聚合分类。
//	@param s
//	@param kF 由元素获取key值的函数
//	@return map[K][]V
func GroupByKF[K comparable, V any](s []V, kF func(v V) K) map[K][]V {
	m := make(map[K][]V)
	for i := range s {
		k := kF(s[i])
		m[k] = append(m[k], s[i])
	}
	return m
}

// GroupByKVF 传入slice和kvF.对每个元素，有kvF计算得到k,v 计算出来k一致的所有v聚合成一个列表[]v,最终形成一个map[k][]v。
//
//	@Description: 类似于GroupByKF，不同的是，聚合的不是原始的切片元素，而是由元素的切片元素派生出来的数据。常见的场景是：一个列表，需要按照某一个字段为key，将另一个字段的取值都聚合成一个列表。
//	@param s
//	@param kvF 由元素获取到的key和要聚合的数据value的函数。
//	@return map[K][]V2
func GroupByKVF[K comparable, V1 any, V2 any](s []V1, kvF func(v V1) (K, V2)) map[K][]V2 {
	m := make(map[K][]V2)
	for i := range s {
		k, v := kvF(s[i])
		m[k] = append(m[k], v)
	}
	return m
}

// Conv 将一个类型slice转化为另一个类型的结构体。
//
// convF 是转换的函数
func Conv[V1 any, V2 any, S ~[]V1](s S, convF func(V1) V2) []V2 {
	if s == nil {
		return nil
	}
	res := make([]V2, 0, len(s))
	for _, v := range s {
		res = append(res, convF(v))
	}
	return res
}

// ConvIfOk 将一个类型slice转化为另一个类型的结构体。类似于Conv，但是转换后的列表仅包含转换OK的元素。
// convF函数返回两个值，第一个是转换后的元素，第二个是转换是否成功。
// 对于转换后的结果不需要的元素，返回false；对于需要的元素，返回true。
func ConvIfOk[V1 any, V2 any, S ~[]V1](s S, convF func(V1) (V2, bool)) []V2 {
	if s == nil {
		return nil
	}
	res := make([]V2, 0, len(s))
	for _, v := range s {
		if v2, ok := convF(v); ok {
			res = append(res, v2)
		}
	}
	return res
}

// ConvInts [T,F] 将一个具体的整型切片([]T)转换为另一个具体整型([]F)的切片。T是目标类型，F是来源类型
func ConvInts[T numbers.AllInt, F numbers.AllInt, S ~[]F](s S) []T {
	return Conv(s, numbers.ConvInt[T, F])
}

// ToAny 用于将一个具体的类型的T []T slice转化为[]any，也就是[]interface{}
// 因为有些包的入参必须是[][]any
func ToAny[T any](s []T) []any {
	p := make([]any, 0, len(s))
	for i := range s {
		p = append(p, s[i])
	}
	return p
}

// To2DAny 所有将一个具体的2维[][]T slice转化为为 [][]any
// 因为有些包的入参必须是[][]any
func To2DAny[T any](s [][]T) [][]any {
	p := make([][]any, 0, len(s))
	for i := range s {
		p = append(p, ToAny(s[i]))
	}
	return p
}

// Filter 原地挑选。remainF是一个谓词函数,对返回true的元素进行保留，false删除。 是对官方的slices.DeleteFunc包进行了简单包装。
func Filter[S ~[]E, E any](s S, remainF func(E) bool) S {
	delF := func(v E) bool { return !remainF(v) }
	return slices.DeleteFunc(s, delF)
}

// Pick 挑选pred为true的元素。和Filter的区别是，Filter是原地修改Slice，Pick是创建一个新切片。
func Pick[S ~[]E, E any](s S, pred func(E) bool) S {
	ok := make([]E, 0, len(s))
	for _, v := range s {
		if pred(v) {
			ok = append(ok, v)
		}
	}
	return ok
}

func Split[S ~[]E, E any](s S, pred func(E) bool) (ok, not S) {
	for _, v := range s {
		if pred(v) {
			ok = append(ok, v)
		} else {
			not = append(not, v)
		}
	}
	return ok, not
}

func SplitWithIgnore[S ~[]E, E any](s S, predWithIgnore func(E) (ok bool, ignore bool)) (ok, not S) {
	for _, v := range s {
		if b, ignore := predWithIgnore(v); ignore {
			continue
		} else if b {
			ok = append(ok, v)
		} else {
			not = append(not, v)
		}
	}
	return ok, not
}

// SafeSlice s[l:r]避免索引越界的安全版本。
//
//	@Description: 具体地说，l,r小于0的时候，会自动调整为0；r大于s的长度的时候，会自动地调整为切片的长度。
//	@param s
//	@param l
//	@param r
//	@return S
func SafeSlice[S ~[]E, E any](s S, l, r int) S {
	if s == nil {
		return []E{}
	}
	l, r = numbers.Max(l, 0), numbers.Max(r, 0)
	r = numbers.Min(r, len(s))
	if l > r {
		return []E{}
	} else {
		return s[l:r]
	}
}

// GetPageIndexRange 获取分页index下标范围，左闭右开。预设输入的参数正常。
func GetPageIndexRange(page, size int) (int, int) {
	return (page - 1) * size, page * size
}

func markUntilFirstDuplicateItem[K comparable](a []K, s map[K]struct{}) {
	for _, x := range a {
		if _, has := s[x]; has {
			return
		}
		s[x] = struct{}{}
	}
}

// Deduplicate 列表去重并保持元素第一次出现的相对顺序。保证原切片不变。
//
//	@Description: Warning: 原列表没有重复元素的时候，返回的列表和原列表共用相同的底层空间。
//	@param a
//	@return []K 如果没有重复元素，返回原切片(共用底层空间)，否则返回去重后的新切片(和元切片不共用底层空间)。
func Deduplicate[K comparable](a []K) []K {
	if len(a) == 0 {
		return a
	}
	s := make(map[K]struct{})
	markUntilFirstDuplicateItem(a, s)
	if len(a) == len(s) {
		return a
	}

	j := len(s)
	b := make([]K, len(s))
	copy(b, a[:j])

	for i := j + 1; i < len(a); i++ {
		if _, has := s[a[i]]; has {
			continue
		}
		s[a[i]] = struct{}{}
		b = append(b, a[i])
	}

	return b
}

// DeduplicateInPlace 类似于Deduplicate,但是原地去重，同样保证元素第一次出现时的相对顺序。
//
//	@param a
//	@return []K
func DeduplicateInPlace[K comparable](a []K) []K {
	if len(a) == 0 {
		return a
	}

	s := make(map[K]struct{})
	markUntilFirstDuplicateItem(a, s)
	if len(a) == len(s) {
		return a
	}

	j := len(s)
	for i := j + 1; i < len(a); i++ {
		if _, has := s[a[i]]; has {
			continue
		}
		s[a[i]] = struct{}{}
		a[j] = a[i]
		j++
	}
	return a[:j]
}

// EnsureNonNil 确保slice不为nil.如果s是nil，返回空切片；否则直接返回。
func EnsureNonNil[S ~[]E, E any](s S) S {
	if s == nil {
		return make(S, 0)
	}
	return s
}

func ForEach[E any, S ~[]E](s S, f func(e E)) {
	for _, v := range s {
		f(v)
	}
}

func ForEach2Args[E any, S ~[]E](s S, f func(e E, i int)) {
	for i, v := range s {
		f(v, i)
	}
}

func ForEach3Args[E any, S ~[]E](s S, f func(e E, i int, s S)) {
	for i, v := range s {
		f(v, i, s)
	}
}
