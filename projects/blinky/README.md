# Blink

Example that loads a WASM module `blink.wasm` that can blink an LED.

## How to run

Create a new project based on this example:

```
mecha new example.com/modules/blinky github.com/hybridgroup/mechanoid-templates/blinky
```

Now install and build the needed WASM module:

```
cd blinky
mecha new module example.com/modules/blink github.com/hybridgroup/mechanoid-templates/modules/blink
mecha build
```

Now you can build and flash your board:

```
mecha flash pybadge -monitor
```
