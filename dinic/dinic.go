package dinic

type Edge struct {
	from int
	to   int
	cap  int
	flow int
}

type Dinic struct {
	inf    int
	need   int
	n      int
	source int
	sink   int
	g      [][]int
	edges  []Edge
	d      []int
	ptr    []int
	q      []int
}

func (sf *Dinic) Init(_n, _s, _t int) {
	sf.inf = 1e9
	sf.need = 1
	sf.n = _n
	sf.source = _s
	sf.sink = _t
	sf.g = make([][]int, sf.n)
	sf.edges = make([]Edge, 0)
	sf.d = make([]int, sf.n)
	sf.ptr = make([]int, sf.n)
	sf.q = make([]int, 0)
}

func (sf *Dinic) AddEdge(from, to, fcap, bcap int) {
	sf.g[from] = append(sf.g[from], len(sf.edges))
	sf.edges = append(sf.edges, Edge{from, to, fcap, 0})
	sf.g[to] = append(sf.g[to], len(sf.edges))
	sf.edges = append(sf.edges, Edge{to, from, bcap, 0})
	for sf.need <= fcap/2 {
		sf.need *= 2
	}
}

func (sf *Dinic) bfs() bool {
	for i := 0; i < sf.n; i++ {
		sf.d[i] = -1
	}
	sf.q = append(sf.q, sf.source)
	sf.d[sf.source] = 0
	for len(sf.q) > 0 {
		i := sf.q[0]
		sf.q = sf.q[1:]
		for _, id := range sf.g[i] {
			e := &sf.edges[id]
			if sf.d[e.to] == -1 && e.cap-e.flow >= sf.need {
				sf.d[e.to] = sf.d[i] + 1
				sf.q = append(sf.q, e.to)
			}
		}

	}
	return sf.d[sf.sink] != -1
}

func (sf *Dinic) dfs(v int, flow int) int {
	if flow == 0 {
		return 0
	}
	if v == sf.sink {
		return flow
	}
	for j := &sf.ptr[v]; *j < len(sf.g[v]); *j = *j + 1 {
		e := &sf.edges[sf.g[v][*j]]
		be := &sf.edges[sf.g[v][*j]^1]
		if sf.d[e.to] != sf.d[v]+1 || e.cap-e.flow < sf.need {
			continue
		}
		pushed := sf.dfs(e.to, min(flow, e.cap-e.flow))
		if pushed > 0 {
			e.flow += pushed
			be.flow -= pushed
			return pushed
		}
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func (sf *Dinic) MaxFlow() int {
	var res int = 0
	for sf.need > 0 {
		if !sf.bfs() {
			sf.need /= 2
			continue
		}
		for i := 0; i < sf.n; i++ {
			sf.ptr[i] = 0
		}
		for pushed := sf.dfs(sf.source, sf.inf); pushed > 0; pushed = sf.dfs(sf.source, sf.inf) {
			res += pushed
		}
	}
	return res
}
