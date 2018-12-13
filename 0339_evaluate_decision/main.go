package main

func compute(graph map[string]map[string]float64, first, second string, visited map[string]bool, value float64) float64 {
	nexts := graph[second]
	for next, ratio := range nexts {
		if _, ok := visited[next]; !ok {
			if next == first {
				return value * ratio
			} else {
				visited[next] = true
				res := compute(graph, first, next, visited, value * ratio)
				if res != -1 {
					return res
				}
				delete(visited, next)
			}
		}
	}
	return -1
}

func buildGraph(equations [][]string, values []float64) map[string]map[string]float64 {
	graph := map[string]map[string]float64{}
    for i, equation := range equations {
    	first := equation[0]
    	second := equation[1]
    	value := values[i]
    	if _, ok := graph[first]; !ok {
    		graph[first] = map[string]float64{}
    	}
    	if _, ok := graph[second]; !ok {
    		graph[second] = map[string]float64{}
    	}
    	graph[first][second] = 1/value
    	graph[second][first] = value
    }
    return graph
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := buildGraph(equations, values)
    res := []float64{}
    for _, query := range queries {
    	first := query[0]
    	second := query[1]
    	value := -1.0
		visited := map[string]bool{}
		value = compute(graph, first, second, visited, 1)
    	res = append(res, value)	
    }
    return res
}

func calcEquation2(equations [][]string, values []float64, queries [][]string) []float64 {
    graph := buildGraph(equations, values)
    for i := range graph {
        graph[i][i] = 1.0
    	for j := range graph {
    		for k := range graph {
    			ratio1, ok1 := graph[j][i]
    			ratio2, ok2 := graph[i][k]
    			if ok1 && ok2 {
    				graph[j][k] = ratio2 * ratio1
    			}
    		}
    	}
    }
    res := []float64{}
    for _, query := range queries {
    	first := query[0]
    	second := query[1]
        value := -1.0
        if _, ok := graph[second][first]; ok {
    		value = graph[second][first]
    	}
        res = append(res, value)
    }
    return res
}

func main() {

}