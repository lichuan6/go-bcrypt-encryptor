# go-bcrypt-encryptor

go-bcrypt-encryptor is a Go package for password encryption and comparison using the bcrypt algorithm.

## Installation

Get and install the go-bcrypt-encryptor package from GitHub using the following command:

```bash
go get github.com/lichuan6/go-bcrypt-encryptor
```

## Usage

Import the `go-bcrypt-encryptor` package:

```go
import "github.com/lichuan6/go-bcrypt-encryptor"
```

Create a new `BcryptEncryptor` instance:

```go
encryptor := bcryptencryptor.NewBcryptEncryptor(bcrypt.DefaultCost)
```

### Encrypting Passwords

```go
password := "myPassword"
encryptedPassword, err := encryptor.Encrypt(password)
if err != nil {
    // Handle error
}
```

### Comparing Passwords

```go
password := "myPassword"
err := encryptor.Compare(encryptedPassword, password)
if err != nil {
    // Passwords do not match
}
```

### Hash and Password Comparison

```go
password := "myPassword"
err := encryptor.CompareHashAndPassword(encryptedPassword, password)
if err != nil {
    // Passwords do not match
}
```

## Testing

Run the following command to execute the tests for the `go-bcrypt-encryptor` package:

```
go test github.com/lichuan6/go-bcrypt-encryptor
```

## Contribution

If you find any bugs or have suggestions for improvement, please create an issue on GitHub or submit a pull request. We welcome and appreciate your contributions!

## License

The go-bcrypt-encryptor package is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
