package Week_07

//字典树解法，需要背诵默写下来

type node struct {
	child [26]*node
	word  bool
}
type WordDictionary struct {
	child *node
}

func Constructor() *WordDictionary {
	return &WordDictionary{child: &node{}}

}

func (this *WordDictionary) AddWord(word string) {
	n := this.child
	l := len(word)
	for i := 0; i < l; i++ {
		c := word[i] - 97
		if n.child[c] == nil {
			n.child[c] = &node{}
		}
		n = n.child[c]
	}
	n.word = true

}

func (this *WordDictionary) Search(word string) bool {
	w := []byte(word)
	return this.search(w, this.child, 0, len(w))
}

func (this *WordDictionary) search(word []byte, n *node, s, l int) bool {
	//fmt.Println(word,s,l)

	for i := s; i < l; i++ {

		c := word[i] - 97
		if n.child[c] == nil {
			return false
		}
		n = n.child[c]

	}
	return n.word
}

func (this *WordDictionary) StartsWithAndWord(prefix string) (bool, bool) {
	n := this.child
	l := len(prefix)
	for i := 0; i < l; i++ {
		//fmt.Println(string(prefix[i]))
		if n.child[prefix[i]-97] == nil {
			return false, false
		}
		n = n.child[prefix[i]-97]

	}

	return true, n.word
}

var dic *WordDictionary
var dir [][]int
var m int
var n int
var ret map[string]bool

func findWords(board [][]byte, words []string) []string {
	dir = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	m = len(board)
	n = len(board[0])
	path := make([]byte, 100)
	ret = make(map[string]bool)
	dic = Constructor()

	l := len(words)
	for i := 0; i < l; i++ {
		dic.AddWord(words[i])
	}

	for j := 0; j < m; j++ {
		for k := 0; k < n; k++ {
			backtracking(board, j, k, path, 0, dic.child)
		}
	}
	re := make([]string, len(ret))
	i := 0
	for k, _ := range ret {
		re[i] = k
		i++

	}
	//fmt.Println(dic.StartsWithAndWord("oath"))
	return re

}
func backtracking(board [][]byte, x, y int, path []byte, index int, trie *node) {

	char := board[x][y]
	path[index] = char
	//fmt.Println(string(char), trie.child[char-97] == nil, index+1, trie)
	if trie.child[char-97] == nil {
		return
	}
	if trie.child[char-97].word {
		s := string(path[0 : index+1])
		ret[s] = true
	}
	o := board[x][y]
	board[x][y] = 0
	for i := 0; i < 4; i++ {
		r := dir[i][0] + x
		c := dir[i][1] + y
		if r < 0 || r >= m || c < 0 || c >= n {
			continue
		}
		if board[r][c] != 0 {
			backtracking(board, r, c, path, index+1, trie.child[char-97])
		}
	}
	board[x][y] = o

}
