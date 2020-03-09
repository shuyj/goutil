#Caffine for go simple

~~~
    cache/cache_test.go : TestAsyncNBlockCache
    key := fmt.Sprintf("key_%d", i)
    val, rtime, err := GetWithTime(key)

    if err != nil {
        fmt.Printf("Get error = %v\n", err)
        continue
    }
    if rtime <= 10*time.Millisecond {
        // expired
        // reload asynchronous
        common.Async(func() {
            // get data with key
            fmt.Printf("reload key=%s\n", key)
            SetEx(key, i, 10*time.Second)
        })
    }
    // return cached data whether or not expired
    val=val

~~~