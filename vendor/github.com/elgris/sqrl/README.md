# sqrl - fat-free version of squirrel - fluent SQL generator for Go

**Non thread safe** fork of [squirrel](http://github.com/lann/squirrel). The same handy fluffy helper, but with extra letters removed :)

```go
import "github.com/elgris/sqrl"
```

[![GoDoc](https://godoc.org/github.com/elgris/sqrl?status.png)](https://godoc.org/github.com/elgris/sqrl)
[![Build Status](https://travis-ci.org/elgris/sqrl.png?branch=master)](https://travis-ci.org/elgris/sqrl)

**Requires Go 1.8 and higher**

## Inspired by

- [squirrel](https://github.com/lann/squirrel)
- [dbr](https://github.com/gocraft/dbr)

## Why to make good squirrel lighter?

Ask [benchmarks](https://github.com/elgris/golang-sql-builder-benchmark) about that ;). Squirrel is good, reliable and thread-safe with it's immutable query builder. Although immutability is nice, it's resource consuming and sometimes redundant. As authors of `dbr` say: "100% of our application code was written without the need for this".

## Why not to use dbr then?

Although, `dbr`'s query builder is proven to be much [faster than squirrel](https://github.com/tyler-smith/golang-sql-benchmark) and even faster than [sqrl](https://github.com/elgris/golang-sql-builder-benchmark), it doesn't have all syntax sugar. Especially I miss support of JOINs, subqueries and aliases.

## Usage

**sqrl is not an ORM.**, it helps you build SQL queries from composable parts.
**sqrl is non thread safe**. SQL builders change their state, so using the same builder in parallel is dangerous.

It's very easy to switch between original squirrel and sqrl, because there is no change in interface:

```go
import sq "github.com/elgris/sqrl" // you can easily use github.com/lann/squirrel here

users := sq.Select("*").From("users").Join("emails USING (email_id)")

active := users.Where(sq.Eq{"deleted_at": nil})

sql, args, err := active.ToSql()

sql == "SELECT * FROM users JOIN emails USING (email_id) WHERE deleted_at IS NULL"
```

```go
sql, args, err := sq.
    Insert("users").Columns("name", "age").
    Values("moe", 13).Values("larry", sq.Expr("? + 5", 12)).
    ToSql()

sql == "INSERT INTO users (name,age) VALUES (?,?),(?,? + 5)"
```

Like [squirrel](https://github.com/lann/squirrel), sqrl can execute queries directly:

```go
stooges := users.Where(sq.Eq{"username": []string{"moe", "larry", "curly", "shemp"}})
three_stooges := stooges.Limit(3)
rows, err := three_stooges.RunWith(db).Query()

// Behaves like:
rows, err := db.Query("SELECT * FROM users WHERE username IN (?,?,?,?) LIMIT 3", "moe", "larry", "curly", "shemp")
```

Build conditional queries with ease:

```go
if len(q) > 0 {
    users = users.Where("name LIKE ?", q)
}
```

### MySQL-specific functions

#### [Multi-table delete](https://dev.mysql.com/doc/refman/5.7/en/delete.html)

```go
sql, args, err := sq.Delete("a1", "a2").
    From("z1 AS a1").
    JoinClause("INNER JOIN a2 ON a1.id = a2.ref_id").
    Where("b = ?", 1).
    ToSql()
```

### PostgreSQL-specific functions

#### [JSON values](https://www.postgresql.org/docs/9.3/static/functions-json.html)

Package [pg](https://godoc.org/github.com/elgris/sqrl/pg) contains JSON and JSONB operators that use json.Marshal to serialize values and cast them to appropriate column type.

```go
sql, args, err := sq.Insert("posts").
    Columns("content", "tags").
    Values("Lorem Ipsum", []string{"foo", "bar"}).
    ToSql()
```

## License

Sqrl is released under the
[MIT License](http://www.opensource.org/licenses/MIT).
