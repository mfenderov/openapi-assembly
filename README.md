# OpenAPI Assembly
A small CLI wrapper on top of the awesome [bdpiprava/scalar-go](https://github.com/bdpiprava/scalar-go) package.   
The project serves a single purpose - a simple CLI to assemble OpenAPI Specification from segmented files.   
For more details look no further than [reading-from-segmented-files](https://github.com/bdpiprava/scalar-go?tab=readme-ov-file#reading-from-segmented-files).

## Installation:
```bash
go install github.com/mfenderov/openapi-assembly@latest
```

## Usage
### `openapi-assembly -h`
```bash
Usage of openapi-assembly:
  -i string
        input specification (default "docs/api.yaml") # atm supports only *.yml and *.yaml extensions
  -o string
        output specification (default "api.json")     # atm supports only *.json extension
```
### Example:
```openapi-assebly -i test-resources/docs/api.yaml -o tmp/api.json```



