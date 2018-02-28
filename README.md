### CFPB PlatformOps EC 


## Installation

 - Install Go 
 - Set up GOPATH
 - Clone this repo to $GOPATH/src
 
----

## Evidence Collection (EC)

**Description**:  Evidence Collection project is written in Golang to support the automated process of collecting 
evidence from a set of baselines. 

**Project code structure**:

 ```
 - models/
        - baseline.go
        - control.go
        - ecGlobalVars.go
        - manifest.go
 - services/
        - commandExe.go
        - crudBaselineControl.go
        - readExcelBaseline.go
 - ec_agent.go
 - parse_Excel_convertTo_Json.go
 - parse_Excel_loadTo_Sql.go
```

**Convention**:
- Class names should be nouns in UpperCamelCase.
- Methods should be verbs in lowerCamelCase. 
- Local variables, instance variables, and class variables are also written in lowerCamelCase.
- Constants should be written in uppercase characters separated by underscores.
- 4 spaces for an indentation.
- When in doubt, see [Effective Go](https://golang.org/doc/effective_go.html#mixed-caps)

## Dependencies

- github.com/tealeg/xlsx - Excel parser
- github.com/lib/pq - PostgreSQL Go driver

## Configuration

- Please contact __DL_CFPB_Platform_Operations_Team_ for configuration setup.

## Usage


## How to test the software


## Known issues


## Getting help


## Getting involved


----

## Open source licensing info
1. [TERMS](TERMS.md)
2. [LICENSE](LICENSE)
3. [CFPB Source Code Policy](https://github.com/cfpb/source-code-policy/)


----

## Credits and references
1. GoLang
2. Git
3. [Design Patterns - The Gang of Four](https://www.amazon.com/Design-Patterns-Object-Oriented-Addison-Wesley-Professional-ebook/dp/B000SEIBB8)
