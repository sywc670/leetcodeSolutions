package code

// lc 208. 实现 Trie (前缀树)
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
