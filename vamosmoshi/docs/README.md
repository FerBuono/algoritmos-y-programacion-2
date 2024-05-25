
# Documentation

This directory contains additional documentation for the VamosMoshi project.

## Contents

- Detailed descriptions of the algorithms used.
- Explanation of the data formats.
- User guides and tutorials.

## Algorithms

- Dijkstra's Algorithm for shortest path finding.
- Prim's Algorithm for minimum spanning tree.
- Hierholzer's Algorithm for finding Eulerian cycles.
- Fleury's Algorithm for finding Eulerian paths.

## Data Formats

### PJ Files

PJ files contain city and route data formatted as follows:

- First line: Number of cities.
- Following lines: City name and coordinates.
- Subsequent lines: Routes between cities and their distances.

### KML Files

KML files are used for visualizing routes on a map. They follow the standard KML format used by Google Earth.

## User Guides

### Running the Program

To run the program, use the main script `vamosmoshi.py` with the path to the desired PJ file as an argument.

Example:
```sh
python3 scripts/vamosmoshi.py data/destino.pj
```

### Commands
The program accepts the following commands via standard input:

- `ir, origin, destination, output_file.kml`
- `itinerario, input_file.txt`
- `viaje, origin, output_file.kml`
- `reducir_caminos, output_file.pj`
- `salir`