# Morse Code API in GO

Create a ```.env``` like the ```.env.example```

Run in the command line:

```sh
go get
```

```sh
go mod tidy
```

For run the project:

```go
go run main.go
```

## Doocker build and run

```sh
docker build -t morse-code-go-api:latest .
```

```sh
docker run -dp PORT:ENVPORT morse-code-go-api:latest
```

## Apis

<!-- ### <span style="color:#25881e">Encode</span> -->
- ### Encode

<!-- #### <span style="color:#FFF151">POST</span> -->
#### $\color{#FFF151}POST$

> localhost:8090/encode

```json
{
    "text":"Hello"
}
```

<!-- #### <span style="color:#32DE84">GET</span> -->
#### $\color{#32DE84}GET$

> localhost:8090/encode/`hello`

- ### Decode
  
#### ${\color{#FFF151}POST}$

> localhost:8090/decode

```json
{
    "text":".... . .-.. .-.. ---"
}
```

#### ${\color{#32DE84}GET}$

<!-- >localhost:8090/decode/<span style="color:#32DE84">.... . .-.. .-.. ---</span> -->
> localhost:8090/encode/`.... . .-.. .-.. ---`
