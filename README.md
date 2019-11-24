# status-code-check
Status code check of URLs

| Flag    | Name  | Example Input            |
| :-----: |:-----:| :----------------------- |
| -u      | url   | https://google.com       |
| -p      | path  | "./path/to/filename.txt" |


## Usage 

### Check the status code of a URL
#### Terminal
```
go run main.go -u https://google.com
```

#### Output
```
200 | https://google.com
```

---
### Check the status codes of all URLs in a file
#### filename.txt *(File Content)*
```
https://jsonplaceholder.typicode.com/posts
https://jsonplaceholder.typicode.com/comments
https://jsonplaceholder.typicode.com/albums
https://jsonplaceholder.typicode.com/photos
https://jsonplaceholder.typicode.com/todos
https://jsonplaceholder.typicode.com/users
```

#### Terminal
```
go run main.go -p "path/to/filename.txt"
```

#### Output
```
200 | https://jsonplaceholder.typicode.com/posts
200 | https://jsonplaceholder.typicode.com/comments
200 | https://jsonplaceholder.typicode.com/albums
200 | https://jsonplaceholder.typicode.com/photos
200 | https://jsonplaceholder.typicode.com/todos
200 | https://jsonplaceholder.typicode.com/users
```