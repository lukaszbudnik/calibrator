# calibrator

Calibrator is a framework for performing basic analytics on multi-dimensional data sets.

Framework expects that data is organised as a slice of slice of strings (`[][]string`). The first record is the header. Header is used for filtering and grouping fields.

An example would be a CSV file (and you can use JSON, XML, SQL, NoSQL, anything that can be turned into a slice of slice of strings).

```
timestamp,group,subgroup,value
2019-02-03T18:25:59+01:00,home,temp,25
2019-02-03T18:25:59+01:00,work,temp,21
2019-02-03T18:26:00+01:00,home,temp,24
2019-02-03T18:26:00+01:00,home,humidity,50
2019-02-03T18:26:00+01:00,home,light,3
```

# Functionality

Here are some examples to get you started. Please review unit tests for the details.

## Filtering

In order to filter a data set by a specific dimension use `Filter` function. The map parameter represents the filter conditions. `Filter` returns a slice of slice of strings (`[][]string`). `Filter` preserves the header row.

```go
condition := map[string]string{"group": "home"}
filtered, err := Filter(records, condition)
```

In fact, you can filter on as many dimensions as you like:

```go
condition = map[string]string{"group": "home", "subgroup": "temp"}
filtered, err := Filter(records, condition)
```

## Grouping

Another useful function is grouping data sets by specific dimension.

`Group` function returns a map of slice of slice of strings (`map[string][][]string`). Map keys are the values of the specified dimension. `Group` preserves the header row.

```go
dimension := "group"
grouped, err := Group(records, dimension)
// then use like this
grouped["home"]
grouped["work"]
```

There is also a `RecursiveGroup` function which can perform a multi-level grouping. It returns a slice of `RecGroup` struct which looks like this:

```go
type RecGroup struct {
	Dimension string
	Records   [][]string
	Subgroups []*RecGroup
}
```

Each `RecGroup` contains a slice of pointers to its subgroups. Just like the simple `Group` the `RecursiveGroup` also preserves the header row in the `Records` field.

An example would be:

```go
results, err := RecursiveGroup(records, []string{"group", "subgroup"})
// records grouped by dimension "group"
results[0].Dimension
results[0].Records
// records grouped by dimension "group" and "subgroup"
results[0].Subgroups[0].Dimension
results[0].Subgroups[0].Records
// iterate over all slices (top level results and Subgroups for all groups)
```

## Summarize

The next most useful function is `Summarize`. This function is especially handy when you just start analysing your data set and have very little or no knowledge of it. The dimension must be convertible to float64 otherwise error is returned.

The `Summarize` function returns a `Summary` struct which looks like this:

```go
type Summary struct {
	Samples   int             // contains number of analysed samples
	Variance  float64         // variance of specified dimension
	Min       float64         // min value of the specified dimension
	Max       float64         // max value of the specified dimension
	Mean      float64         // mean value of the specified dimension
	Median    float64         // median value of the specified dimension
	Quartiles stats.Quartiles // Q1, Q2, and Q3 quartiles of the specified dimension
	Outliers  stats.Outliers  // outliers values of the specified dimension
}
```

You can either summarize the whole data set or individual groups:

```go
// get summary of all records
dimension := "value"
all, err := Summarize(records, dimension)
// get summary for g1 group
group := "group"
grouped, err := Group(records, group)
summary, err := Summarize(grouped["g1"], dimension)
```

## Outliers

`Outliers` function returns outliers records (not only values like in the `Summarize`) for the specified dimension. The dimension must be convertible to float64 otherwise error is returned.

To find outliers in the whole data set use:

```go
dimension := "value"
outliers, err := Outliers(records, dimension)
```

To find outliers in individual groups use:

```go
grouped, err := Group(records, "group")
dimension := "value"
outliers, err := Outliers(grouped["g1"], dimension)
```

## Leaders

`Leaders` function is complementary to `Outliers`, but returns only the records whose values of the specified dimension are below first quartile. The dimension must be convertible to float64 otherwise error is returned.

To find leaders in the whole data set use:

```go
dimension := "value"
leaders, err := Leaders(records, dimension)
```

To find outliers in individual groups use:

```go
grouped, err := Group(records, "group")
dimension := "value"
leaders, err := Leaders(grouped["g1"], dimension)
```

## Stability

`Stability` is the function that computes a variance of given dimension and based on passed threshold returns either true or false. The dimension must be convertible to float64 otherwise error is returned.

Note: it is your responsibility to define the threshold and calibrator will not do this for you. You need to analyse the data and then pick a threshold value that best reflects the stability point. Should you need a little bit more information about your data use the `Summarize` function.

To check if data set in a specific group is stable use:

```go
grouped, err := Group(records, "group")
dimension := "value"
threshold := 10
stable, err := Stability(grouped["g1"], dimension, threshold)
```
