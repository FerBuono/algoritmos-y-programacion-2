# TP2 Algogram

## Overview
Algogram is a social media simulation application implemented in **Go**. It allows users to create and manage posts and users, simulating a basic social media platform. 

## Prerequisites
Make sure you have Go installed on your system. You can download it from [Go's official website](https://golang.org/dl/).

## Running the Project

### Steps to Run

1. **Clone the repository** (if you haven't already):
    ```bash
    git clone https://github.com/FerBuono/algoritmos-y-programacion-2.git
    cd algoritmos-y-programacion-2
    cd algogram
    ```
2. **Initialize the Go module** (if not already initialized):
    ```bash
      go mod tidy
    ```
3. **Prepare the input files**:
Ensure you have the usuarios.txt file in the correct format in the same directory as main.go.
    ```bash
    # Example contents of usuarios.txt
    user1
    user2
    user3
    ```
4. **Run the project**:
    ```bash
    go run main.go usuarios.txt
    ```

### Expected Arguments
- **usuarios.txt**: A text file containing a list of users.

### How to interact
Once the program is running, you can interact with it using the following commands:
- **Login**: This command logs in a user.
  ```bash
  login <user>
  # Example:
  login user1
  ```
- **Logout**: This command logs out the current user.
  ```bash
  logout
  ```
- **Publish a post**: This command allows a user to create a post.
  ```bash
  publicar <content>
  # Example:
  publicar "Hello, world!"
  ```
- **View the next post in the feed**: This command allows the user to view the next post in their feed.
  ```bash
  ver_siguiente_feed
  ```
- **Like a post**: This command allows a user to like a post.
  ```bash
  likear_post <post_id>
  # Example:
  likear_post 1
  ```
- **Show likes on a post**: This command shows the users who liked a post.
  ```bash
  mostrar_likes <post_id>
  # Example:
  mostrar_likes 1
  ```
## Running Tests
The `pruebas.sh` script automates the process of running tests. Here's how you can use it:

1. **Navigate to the project directory**:
    ```bash
    cd /algogram
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
