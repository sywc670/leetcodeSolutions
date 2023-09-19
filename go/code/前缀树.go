package code

import "slices"

// 208. 实现 Trie (前缀树)
// 空间浪费比较大，可以用map来优化
// 较为现实的前缀树实现 https://geektutu.com/post/gee-day3.html
type Trie struct {
	isEnd    bool
	children [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, r := range word {
		r -= 'a'
		if node.children[r] == nil {
			node.children[r] = &Trie{}
		}
		node = node.children[r]
	}
	node.isEnd = true
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, r := range prefix {
		r -= 'a'
		if node.children[r] != nil {
			node = node.children[r]
		} else {
			return nil
		}
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}

// map优化版
type TrieMap struct {
	isEnd    bool
	children map[byte]*TrieMap
}

func ConstructorTrieMap() TrieMap {
	return TrieMap{children: make(map[byte]*TrieMap)}
}

func (t *TrieMap) Insert(word string) {
	node := t
	for _, r := range word {
		r -= 'a'
		b := byte(r)
		if node.children[b] == nil {
			node.children[b] = &TrieMap{children: make(map[byte]*TrieMap)}
		}
		node = node.children[b]
	}
	node.isEnd = true
}

func (t *TrieMap) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (t *TrieMap) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}

func (t *TrieMap) SearchPrefix(prefix string) *TrieMap {
	node := t
	for _, r := range prefix {
		r -= 'a'
		b := byte(r)
		if node.children[b] == nil {
			return nil
		}
		node = node.children[b]
	}
	return node
}

// 1268. 搜索推荐系统
func suggestedProducts(products []string, searchWord string) (ans [][]string) {
	t := &TrieProduct{}
	for _, p := range products {
		t.Insert(p)
	}
	for i := 1; i <= len(searchWord); i++ {
		ans = append(ans, t.SearchPrefix(searchWord[:i]))
	}
	return
}

type TrieProduct struct {
	children [26]*TrieProduct
	words    []string
}

func (t *TrieProduct) Insert(word string) {
	node := t
	for _, r := range word {
		r -= 'a'
		if node.children[r] == nil {
			node.children[r] = &TrieProduct{}
		}
		node = node.children[r]
		node.words = append(node.words, word)
	}
}

func (t *TrieProduct) SearchPrefix(prefix string) []string {
	node := t
	for _, r := range prefix {
		r -= 'a'
		if node.children[r] == nil {
			return nil
		}
		node = node.children[r]
	}
	slices.Sort(node.words)
	if len(node.words) > 3 {
		return node.words[:3]
	}
	return node.words
}
