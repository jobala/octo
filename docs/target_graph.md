## Target

A target is the unit of execution.It is what gets executed.

```go
type Target struct {
    // id is the target's identifier of the form package#task
    id   string
    cwd  string
    task string

    // List of "dependency specs" like "^build", "build", "foo#build"
    depSpecs: string[];

    //  Dependencies of the target - these are the targets that must be complete before the target can be complete
    dependencies: string[];

    // Dependents of the target - these are the targets that depend on this target
    dependents: string[];
}
```

## Target Graph

A target graph is obtained from the package graph and the tast pipeline. For example if we have the following package
graph and pipeline

```json
Package Dependencies

{
    packageA: {
        dependsOn: [packageB]
    }
    packageB: {
        dependsOn: []
    }
}
```

```json
Pipeline

{
    build: {
        dependsOn: ['^build']
    },
    test: {
        dependsOn: ['build']
    }
}
```

This will generate the following task graph

```json
{
    packageB#build: [
        dependsOn: []
    ]
    packageB#test: [
        dependsOn: [packageB#build]
    ]
    packageA#build: [
        dependsOn: [packageB#build]
    ]
    packageA#test: [packageA#build]
}
```
