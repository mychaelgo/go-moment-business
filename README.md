# go-moment-business
Moment business in Go

# Reference
(JS) https://github.com/jmeas/moment-business

# Install
```
go get github.com/mychaelgo/go-moment-business
```

# API
```
AddWeekDays(t time.Time, amount int) (time.Time)
```
Add week days to a date, modifying the original date. Returns the date with skipping the weekend.
