# Sample app: Rejoiner with Apollo client

## Structure and Stack

- **[servers/github](./servers/github)**: gRPC server
    - Go([grapi](https://github.com/izumin5210/grapi))
- **[servers/qiita](./servers/qiita)**: gRPC server
    - Ruby
- **[bff-rejoiner](./bff-rejoiner)**: GraphQL server
    - Kotlin([Ktor](https://ktor.io/))
    - [Rejoiner](https://github.com/google/rejoiner)
