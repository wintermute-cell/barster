# barster

`barster` is a content provider for simple, string based status-bars like
[dwm](https://dwm.suckless.org/)s bar, or wayland statusbar projects like
[sewn/dam](https://codeberg.org/sewn/dam). Essentially, `barster` continually
outputs the content of a statusbar as a single string that is then displayed on
the bar.

With `dwm` you would use `barster` like this:
```bash
barster | while read new_status; do xsetroot -name "$new_status"; done
```

With `dam` you would use `barster like this:
```bash
barster | dam
```

## what makes `barster` special?

- Final output composed of individual "modules"
- Collection of [prebuilt modules](./prebuilt) like Battery, Datetime, Nettraffic, ...
- Easy to add new [custom modules](#configuration-adding-modules) tailored to your system
- Modules can update independently on a timer (polling), or even only when
  their data changes (event-driven) for optimal performance

## installation

First, make sure you have [go](https://go.dev/) installed.

```bash
git clone https://github.com/wintermute-cell/barster
cd barster
go build
./barster
```

## configuration, adding modules

To configure, modify `main.go` and recompile using `go build`. The
[prebuilt](./prebuilt) directory contains a (growing) number of prebuilt
modules. To add a new module, add it to the `modules` list in `main.go`:

```go
modules := []pkg.Module{
    // Put the modules here!
}
```


The easiest way to create a custom module is to add a block like this to the `modules`
list in `main.go`:
```go
modules := []pkg.Module{
    {
        Name:     "myModule",      // has to be unique across modules
        Interval: 1 * time.Second, // the module will update every second
        Update: func() string { // the function that returns the module's output
            return "Hello, world! Random number: " + fmt.Sprint(rand.Intn(100))
        },
    },
    // Other modules ...
}
```

### advanced: creating event-driven modules

You should be comfortable with `goroutines` before attempting this!

To save on cpu time, you can set up modules to only re-run when the module
itself wants to. To do this, set the `Interval` to a really high value, and add
a `Ticker`:

```go
modules := []pkg.Module{
    {
        Name:     "eventDrivenModule",
        Interval: 1000 * time.Hour,    // effectively disables periodic updates
        Update: func() string {
            return "Hello World!"
        }, 
        // the ticker is a channel that notifies the
        // statusbar that the module needs to update.
        Ticker: func() chan struct{} {
            ch := make(chan struct{})
            go func() {
                ticker := time.NewTicker(5 * time.Second)
                defer ticker.Stop()
                for {
                    <-ticker.C
                    ch <- struct{}{}
                }
            }()
            return ch
        }(),
    },
    // Other modules ...
}
```

## contributing

You are welcome to contribute new modules you've built! Please make sure they
are as portable as they can be so many people can benefit.
You may also create feature request issues if you are unable to create a module
you need yourself.
