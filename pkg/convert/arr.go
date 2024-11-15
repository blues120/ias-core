package convert

// ToArr 通过函数 func(*A) *B，将 []*A 转换成 []*B
// eg: ToArr[ent.Camera, biz.Camera](CameraEntToBiz, arr)
func ToArr[T any, V any](fn func(*T) *V, arr []*T) []*V {
	list := make([]*V, len(arr))
	for idx, item := range arr {
		list[idx] = fn(item)
	}
	return list
}
