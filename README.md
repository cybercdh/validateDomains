# validateDomains

Reads in a list of domains from `stdin` and outputs those with a valid TLD according the latest list from [ICANN](https://data.iana.org/TLD/tlds-alpha-by-domain.txt).

## Installation

Assuming you have [Go installed](https://go.dev/doc/install) you can run the following:

```bash
go install https://github.com/cybercdh/validateDomains@latest
```

## Usage

```bash
echo foo.example.com | validateDomains
or
cat domains.txt | validateDomains
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)