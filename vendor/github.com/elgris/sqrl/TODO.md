- Replace fmt.Errorf with errors.New() and compare results
- Replace calls like sql.WriteString(strings.Join(b.Options, " ")) with adding data in loop. Thus remove extra memory allocation. Don't forget benchmarks. Something like:
```
func (b *buffer) WriteStrings(s []string, separator string)
```
- Refactor StatementBuilderType to something like: 
```
type StatementBuilderType struct {
    placeholderFormat PlaceholderFormat
    runWith           BaseRunner

    dataStruct Sqlizer // dataStruct depends of concrete builder
}
```