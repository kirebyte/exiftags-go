# exiftags-go

This project was a special request but I fell in love with the design and I wish to continue making it grow.
Looking forward to make it a Go port initiative of exiftags (https://github.com/ejohnst/exiftags)
Architecture based on MVC and intended to be used "the unix way".

## Installation

```go install github.com/kirebyte/exiftags-go```

## Usage/Examples

```
   --recursive, -r           if any of the file parameters is a directory then crawl inside looking for more images (default: false)
   --format value, -f value  set output format to: [csv|html|json|raw]. Example: --format csv (default: "raw")
   --help, -h                show help
```
The way to comply with the original project requirements is using:

```
# "-r" enables nested directories crawling, "-f csv" for csv output. Prints on STDOUT by default. 
exiftags-go -r -f csv images/* 2>/dev/null > gps.csv

# "-r" enables nested directories crawling, "-f html" for html output. Prints on STDOUT by default. 
exiftags-go -r -f html images/* 2>/dev/null > gps.html

# JSON format, just for fun :)
exiftags-go -r -f json  images/* 2>/dev/null > gps.json
```

## Running Tests

``` go test ./...```

## Contributing

Project structure is still a work in progress. 

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Authors

- [@kirebyte](https://www.github.com/kirebyte)

## Acknowledgements

- DO