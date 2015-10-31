### YAML parser
* allow importing other yaml files in one yaml file

### Code generation
* add support for extending commands
* add go generate comment
* add comments to the generated functions
* add comment to generated type
* support url templates
* generated code must have nice format
* check if the given client is valid and if not return an error

### Endpoint parameters
* implement default method: GET
* add support for the postField parameter location
* check if required parameters are actually set

### Tigs binary
* add verbose logging like in goldigen
* add forceStdOut flag
* add overwrite flag
* check if generated code compiles (via extra flag?)
* add optional --out parameter
* add required package parameter
* add --forceExport command to force exporting all operations

### General
* write README
* add travis integration
* add coveralls integration
