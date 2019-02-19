package filetree

type TreeStackCacheKey struct {
	bottomTreeStart, bottomTreeStop, topTreeStart, topTreeStop int
}

type TreeStackCache struct {
	refTrees []*FileTree
	cache    map[TreeStackCacheKey]*FileTree
}

func (cache *TreeStackCache) Get(bottomTreeStart, bottomTreeStop, topTreeStart, topTreeStop int) *FileTree {
	key := TreeStackCacheKey{bottomTreeStart, bottomTreeStop, topTreeStart, topTreeStop}
	if value, exists := cache.cache[key]; exists {
		return value
	} else {

	}
	value := cache.buildTree(key)
	cache.cache[key] = value
	return value
}

func (cache *TreeStackCache) buildTree(key TreeStackCacheKey) *FileTree {
	newTree := StackTreeRange(cache.refTrees, key.bottomTreeStart, key.bottomTreeStop)
	for idx := key.topTreeStart; idx <= key.topTreeStop; idx++ {
		newTree.CompareAndMark(cache.refTrees[idx])
	}
	return newTree
}

func (cache *TreeStackCache) Build() {
	var bottomTreeStart, bottomTreeStop, topTreeStart, topTreeStop int

	// case 1: layer compare (top tree SIZE is fixed (BUT floats forward), Bottom tree SIZE changes)
	for selectIdx := 0; selectIdx < len(cache.refTrees); selectIdx++ {
		bottomTreeStart = 0
		topTreeStop = selectIdx

		if selectIdx == 0 {
			bottomTreeStop = selectIdx
			topTreeStart = selectIdx
		} else {
			bottomTreeStop = selectIdx - 1
			topTreeStart = selectIdx
		}

		cache.Get(bottomTreeStart, bottomTreeStop, topTreeStart, topTreeStop)
	}

	// case 2: aggregated compare (bottom tree is ENTIRELY fixed, top tree SIZE changes)
	for selectIdx := 0; selectIdx < len(cache.refTrees); selectIdx++ {
		bottomTreeStart = 0
		topTreeStop = selectIdx
		if selectIdx == 0 {
			bottomTreeStop = selectIdx
			topTreeStart = selectIdx
		} else {
			bottomTreeStop = 0
			topTreeStart = 1
		}

		cache.Get(bottomTreeStart, bottomTreeStop, topTreeStart, topTreeStop)
	}
}

func NewFileTreeCache(refTrees []*FileTree) TreeStackCache {

	return TreeStackCache{
		refTrees: refTrees,
		cache:    make(map[TreeStackCacheKey]*FileTree),
	}
}
