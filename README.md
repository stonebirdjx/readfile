# readfile
A read file package for Go
# Install
```
  go get -u github.com/stonebirdjx/readfile
```
# Example

```
    fileName := "test.txt"
    fr ,err := NewReader(fileName)
    if err != nil{
        log.Fatal(err)
    }
    go fr.Scanner()
    for line := range fr.ContentChan{
        fmt.Println(line)
    }

    // Output:
    // line 1
    // line 2
    // line 3
    // ....
```