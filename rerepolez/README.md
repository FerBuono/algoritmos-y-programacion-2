# TP1 ReRePolez Voting System

## Overview
This project is a voting system implemented in Go. It includes modules for handling voter registration, party management, and vote processing. 


## Prerequisites
Make sure you have Go installed on your system. You can download it from [Go's official website](https://golang.org/dl/).

## Running the Project

### Steps to Run

1. **Clone the repository** (if you haven't already):
    ```bash
    git clone https://github.com/FerBuono/algoritmos-y-programacion-2.git
    cd algoritmos-y-programacion-2
    cd rerepolez
    ```
2. **Initialize the Go module** (if not already initialized):
    ```bash
      go mod tidy
    ```

3. **Prepare the input files**:
Ensure you have the `partidos.csv` and `padron.txt` files in the correct format in the same directory as `main.go`.
    ```bash
    # Example contents of partidos.csv
    President_Party,President1,President2,President3
    Governor_Party,Governor1,Governor2,Governor3
    Mayor_Party,Mayor1,Mayor2,Mayor3

    # Example contents of padron.txt
    12345678
    87654321
    11223344
    ```
4. **Run the project**:
    ```bash
    go run main.go partidos.csv padron.txt
    ```

### Expected Arguments
- **partidos.csv**: A CSV file containing a list of political parties and their candidates.
- **padron.txt**: A text file containing a list of voters.

### How to interact
Once the program is running, you can interact with it using the following commands:
- **Enter a voter**: this command adds a voter to the queue.
  ```bash
  ingresar <padron>
  # Example:
  ingresar 12345678
  ```
- **Vote for a candidate**: This command allows a voter to vote for a candidate in a specified position.
  ```bash
  votar <Presidente/Gobernador/Intendente> <partido>
  # Example:
  votar Presidente 1
  ```
  - **Finalize a vote**: This command finalizes the voting process for the current voter.
  ```bash
  fin-votar
  ```
- **Undo the last vote**: This command undoes the last action performed by the current voter.
  ```bash
  deshacer
  ```

## Running Tests
The `pruebas.sh` script automates the process of running tests. Here's how you can use it:

1. **Navigate to the project directory**:
    ```bash
    cd /rerepolez
    ```
2. **Make the script executable** (if not already executable):
    ```bash
    chmod +x tests/pruebas.sh
    ```

3. **Make the main.go file executable**:
    ```bash
    chmod +x chmod +x main.go
    ```
4. **Run the script**:
    ```bash
    ./tests/pruebas.sh main.go
    ```