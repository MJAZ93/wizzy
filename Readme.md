
# Wizzy
A directory-based, easy-to-use code generator that integrates seamlessly with your project.

![Project Logo](build/img/logo.png)

![Preview](build/img/preview.gif)

## How to Use

### Folder Structure
- Each folder represents a new template.
- Every folder must include a `template.json` file to define the template's rules and parameters.

See test directory.

### template.json
The `template.json` file includes:
1. **Description (`desc`)**: Metadata about the template.
2. **Rules (`rules`)**: Defines how files are created or updated. Rules can also reference other templates (`template/template`), starting always with the temp.
3. **Parameters (`parameters`)**: Inputs required by the template, such as:
    - **`free`**: Simple inline text.
    - **`select`**: A list of options.
    - **`formatted`**: Textarea-like input.

#### Example `template.json`:
```json
{
  "desc": {
    "name": "Template Name",
    "description": "Template Description"
  },
  "rules": [
    {
      "rule": "{{name}}Activity.java -> ./features/{{feature}}",
      "condition": ""
    },
    {
      "rule": "template -> package",
      "condition": "name==Login"
    }
  ],
  "parameters": [
    {
      "id": "name",
      "desc": "The name of the service. Ex: BananaService",
      "regex": "ˆ*Service",
      "type": "free",
      "required": true
    },
    {
      "id": "type",
      "desc": "The type of data exchange",
      "regex": "",
      "type": "select",
      "required": false,
      "options": ["GET", "POST", "PUT", "DELETE"]
    },
    {
      "id": "objectOut",
      "desc": "The output object",
      "type": "formatted",
      "required": true,
      "condition": "data_exchange==BODY"
    }
  ]
}
``` 

### Files
- Each folder should include the template files used for creation or modification.

#### File Types
1. **`.n` (New Files)**: Used when the target file does not exist.
2. **`.e` (Edit Files)**: Used when updating an existing file.

For example: `Main.java.e`, `Main.java.n`, `Main.svelte.e`. 

#### Syntax for `.n` Files
- Use `{{PARAM_NAME}}` to reference a parameter.
- Use `{%if(condition)%}...{%endif%}` to apply conditions.

Example `.n` File:
```java
package co.vm.features.{{feature}};

import co.vm.features.apis.Api;
// New imports come here

public class {{feature}}Services {
    // Service {{name}} {{feature}}
    public void {%if(type=="GET")%}fetch{{name}}{%endif%}{%if(type=="POST")%}post{{name}}{%endif%}() {
        return {{objectOut}};
    }

    // New services come here
}
```

#### Syntax for `.e` Files
- Use `@@` followed by a regex to define insertion points.
- Use `-@@` to mark the end of a block.

Example `.e` File:
```java
@@// New Services come here
// Service {{name}} {{feature}}
public void {%if(type=="GET")%}fetch{{name}}{%endif%}{%if(type=="POST")%}post{{name}}{%endif%}() {
    // ...
}
-@@

@@// New Declarations come here
// {{name}} {{feature}}
-@@
```

## Installation

### Manual Installation
- Add the `wizzy` binary to the root of your project.
- Create a `.wizzy` folder and start defining your templates and rules.

### One-Liner Installation (macOS)
Run the following command in your terminal:
```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/MJAZ93/wizzy/main/build/remote-mac.sh)"
```

## License
Wizzy is distributed under the [Fair Code License](https://faircode.io/).

## Donate - Support Development
Help Wizzy grow by supporting us on [Patreon](https://www.patreon.com/MJAZ) or [Ko-Fi](https://ko-fi.com/afonsomatlhombe).

### Work in Progress
Stay tuned for updates and enhancements!