### Code generation
* add go generate comment
* add comments to the generated functions
* support url templates
* check if the given client is valid and if not return an error
* normalize endpoint parameter names from snake case to camel case

### Endpoint parameters
* add support for the postField parameter location
* check if required parameters are actually set

### Tigs binary
* add verbose logging like in goldigen
* add forceStdOut flag
* add overwrite flag
* check if generated code compiles (via extra flag?)
* add optional --out parameter
* add --forceExport command to force exporting all operations

### General
* write README
* add travis integration
* add coveralls integration
