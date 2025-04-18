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

```sf
docker build -t morse-code-go-api:latest .
```

```sh
docker run -dp PORT:ENVPORT morse-code-go-api:latest
```

## Apis

<!-- ### <span style="color:#25881e">Encode</span> -->
- ### <span >Encode</span>

#### <span style="color:#FFF151">POST</span>

> localhost:8090/encode

```json
{
    "text":"Hello"
}
```

#### <span style="color:#32DE84">GET</span>

> localhost:8090/encode/<span style="color:#32DE84">hello</span>

- ### Decode

#### <span style="color:#FFF151">POST</span>

>localhost:8090/decode

```json
{
    "text":".... . .-.. .-.. ---"
}
```

#### <span style="color:#32DE84">GET</span>

>localhost:8090/decode/<span style="color:#32DE84">.... . .-.. .-.. ---</span>
