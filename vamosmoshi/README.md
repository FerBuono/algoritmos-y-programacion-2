# VamosMoshi

This repository contains the VamosMoshi project, which is designed to handle graph-based operations for various tasks such as finding routes, generating itineraries, and creating KML files for visualization.

## Project Structure

- **data/**: Contains data files used by the scripts.
- **docs/**: Contains documentation files.
- **libs/**: Contains the main Python scripts for the project.

## How to Run

To run the main script, use the following command:
```sh
python3 scripts/vamosmoshi.py <path_to_sedes_file>
```

## Files

### Libs
- `vamosmoshi.py`: Main entry point for running the program.
- `funciones_auxiliares.py`: Contains auxiliary functions.
- `funciones_grafos.py`: Contains graph-related functions.
- `funciones_vamosmoshi.py`: Contains main functions for VamosMoshi operations.
- `grafo.py`: Defines the Grafo class used throughout the project.

### Data
- `destino.pj`: Contains destination data.
- `mapa.kml`: KML file for map visualization.
- `qatar.pj`: Another set of destination data.
- `recomendacion_ejemplo.txt`: Example recommendation data.
- `viaje.kml`: KML file for travel visualization.