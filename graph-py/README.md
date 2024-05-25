# Python Graph Implementation

This repository contains a Python implementation of a graph data structure along with various graph algorithms and a Union-Find data structure.

## Structure

- **grafo.py**: Contains the `Grafo` class which implements the graph data structure.
- **union_find.py**: Contains the `UnionFind` class which implements the union-find data structure.
- **ejercicios.py**: Contains various graph algorithms and example usages of the `Grafo` class.

## Methods in `Grafo` Class

- `agregar_vertice(v)`: Adds a vertex `v` to the graph.
- `eliminar_vertice(v)`: Removes vertex `v` from the graph.
- `agregar_arista(v1, v2, p)`: Adds an edge between vertices `v1` and `v2` with weight `p`.
- `eliminar_arista(v1, v2)`: Removes the edge between vertices `v1` and `v2`.
- `peso(v1, v2)`: Returns the weight of the edge between `v1` and `v2`.
- `obtener_vertices()`: Returns a list of all vertices in the graph.
- `adyacentes(v)`: Returns a list of vertices adjacent to `v`.
- `vertice_aleatorio()`: Returns a random vertex from the graph.
- `contiene_arista(v1, v2)`: Checks if there is an edge between `v1` and `v2`.

## Methods in `UnionFind` Class

- `find(v)`: Finds the root of the set containing `v`.
- `union(u, v)`: Unites the sets containing `u` and `v`.

## Example Usage

```python
from grafo import Grafo
from union_find import UnionFind

# Create a graph
grafo = Grafo(es_dirigido=False, lista_vertices=["A", "B", "C", "D"])
grafo.agregar_arista("A", "B", 1)
grafo.agregar_arista("A", "C", 2)
grafo.agregar_arista("B", "D", 3)

# Print vertices and their adjacencies
print("Vertices:", grafo.obtener_vertices())
for v in grafo.obtener_vertices():
    print(f"{v}: {grafo.adyacentes(v)}")

# Create a Union-Find structure
uf = UnionFind(["A", "B", "C", "D"])
uf.union("A", "B")
print("A and B are in the same set:", uf.find("A") == uf.find("B"))
```