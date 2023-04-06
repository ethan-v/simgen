# Simgen

Simple and power full CLI for generate code from definition in a json file

## Build project

```shell
make build
```

## Usage

Easy to setup, follow these steps:

### Step 1: Download pre-build simgen for Linux.

1. Download latest release here: https://github.com/ethan-v/simgen/releases

2. Create a directory to copy simgen file and the templates.

```shell
mkdir <your_project_path>/.sim
chmod +x .sim/simgen
```

### Step 2: Define templates and params by yourself.

Example

**.sim/form.tpl.html**

```html
<form>
    {{range $key, $value := .}}
        <div class="mb-4">
            <label for="{{$key}}">{{$value.Label}}</label>
            <input type="{{$value.Type}}" id="{{$key}}" name="{{$key}}" placeholder="{{$value.Placeholder}}">
        </div>
    {{end}}
</form>
```

**.sim/form.tpl.json***

```json
{
    "name": {
        "type": "text",
        "label": "Name",
        "placeholder": "Enter your name"
    },
    "email": {
        "type": "email",
        "label": "Email"
    },
    "password": {
        "type": "password",
        "label": "Password"
    }
}
```

### Step 3: Run command to generate from template files

```shell
./.sim/simgen -template=./.sim/form.tpl.html -data=./.sim/form.tpl.json -output=../generated_form.html

# OR
cd .sim
./simgen -template=./form.tpl.html -data=./form.tpl.json -output=../generated_form.html
```

Check file: generated_form.html

Enjoy it!
