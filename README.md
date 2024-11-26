# GitHub_user_activity

A simple CLI tool to fetch user's github activity in terminal. 

---

### Roadmap Project Challenge

This project was developed as part of the [GitHub User Activity from Roadmap.sh](https://roadmap.sh/projects/github-user-activity), designed to enhance skills in building practical applications.

---


## Installation

To install and run the github user activity CLI, ensure you have Golang installed on your system. Then follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/samrat-rm/GitHub_user_activity.git
   ```

2. Navigate to the project directory:

   ```bash 
   cd github_user_activity
   ```

3. Build the application:

   ```bash
   go build -o github-activity
   ```

---

## Usage

Here’s how to use the log the user's github activity :

```bash
./github-activity <username>
```

or you can move your ./github-activity file for global access using 

```bash 
sudo mv github-activity /usr/local/bin/
```
and now we can use the command anywhere in your local 

```bash
./github-activity <username>
```

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
