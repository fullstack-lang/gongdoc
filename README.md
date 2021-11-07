# Gongdoc
A gong stack for editing UML diagrams of gong code

### compile application
> cd ng; npm i; ng build;

### launch application

at the root of the repository (requires go >= 1.16)
> go run main.go

### expected result

launch browser on http://localhost:8080

![result](gongdoc.png)

# Import this stack

add to dev dependencies of your package.json
```json
        "backbone": "^1.4.0",
        "jointjs": "^3.2.0",
        "jquery": "^3.5.1",
        "lodash": "^4.17.20",
        "@types/backbone": "^1.4.5",
        "@types/jointjs": "^2.0.0",
        "@types/jquery": "^3.5.3",
        "@types/lodash": "^4.14.162"
```

and add the following path in your tsconfig.json

```json
      "jointjs": [
        "./node_modules/jointjs"
      ],
      "@types/jointjs": [
        "./node_modules/@types/jointjs"
      ],
```