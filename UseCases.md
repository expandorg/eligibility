## GetAllFilters()
```
SELECT * from filters;
```

```
{
  languages: [
    {id: 1, value: "english"},
    ...
  ],
  interests: [
    {id: 2, value: "machine learning"},
    ...
  ],
  education: [
    {id: 3, value: "master degree"},
    ...
  ]
}
```

## GetJobFilters(jobID int)

```
Select * from filters AS f inner join filters_jobs AS fj on f.`id` = fj.`filter_id` where fj.`job_id` = ?
```

```
{
  languages: [
    {id: 1, value: "english"},
    ...
  ],
  interests: [
    {id: 2, value: "machine learning"},
    ...
  ],
  education: [
    {id: 3, value: "master degree"},
    ...
  ]
}
```

## GetJobWhiteList(jobID int)

```
{
  white_list: ["email1@gmail.com", "email2@gmail.com"]
}
```

## CreateJobFilters(jobID int, filters []Filter, profile Profile)

```
// filters
INSERT INTO filters_jobs (filter_id, job_id)
values
(?, ?)

//profile
INSERT INTO worker_profiles (worker_id, age, city, locality, country)
values
(?, ?, ?, ?, ?)
```


```
{
  filters: [1, 3, 5],
  age: {  
    value: 30
    operation: ">="
  },
  location: {
    value: [city: "", locality: "", country: ""]
    operation: "city"
  }
}
```

## GetEligibleWorkerCount(filters []Filter)

```
TODO
```

```
{
  count: 3
}
```

## GetWorkerProfile(workerID int)

```
select * from worker_profiles where worker_id=?

select * from worker_profiles where worker_id=1Select * from filters AS f inner join filters_workers AS fj on f.`id` = fj.`filter_id` where fj.`worker_id` = ?
```

```
{
  languages: [
    {id: 1, value: "english"},
    ...
  ],
  interests: [
    {id: 2, value: "machine learning"},
    ...
  ],
  education: [
    {id: 3, value: "master degree"},
    ...
  ]
}
```

## CreateWorkerProfile(workerID int, profile []Profile)

```

```

```
{
  filters: [1, 3, 5],
  date_of_birth: 9/9/1978,
  location: {
    city: "San Francisco",
    locality: "California",
    country: "USA"
  }
}
```

## GetEligibleJobsForWorker(workerID int)

```
{
  job_ids: [3, 5, 6]
}
```

## IsWorkerEligibleForJob(workerID int, jobID int)

```
Select * from filters AS f inner join filters_jobs AS fj on f.`id` = fj.`filter_id` where fj.`job_id` = ?

select * from worker_profiles where worker_id=?
Select * from filters AS f inner join filters_workers AS fj on f.`id` = fj.`filter_id` where fj.`worker_id` = ?

```

```
{
  eligible: true
}
```