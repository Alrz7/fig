# FIG (Flexible In-memory ConFIGuration)

**Note:** FIG was originally developed for personal use. Ongoing development is driven by the needs of the projects where it is integrated.

FIG is a lightweight, file-backed configuration library for Go. It keeps your configuration in memory as native Go structs and automatically synchronises changes to JSON files. You can group related settings into **Handlers** and **Fields**, making it easy to manage complex application configurations.

## Features

- **In‑memory & persistent** – work with plain Go pointers; changes are saved to JSON.
- **Handler & Field grouping** – organise configuration into logical sections.
- **Automatic restore** – load saved state at startup and reconcile with your code.
- **Pointer‑based tracking** – FIG watches the variables you give it (no manual getters/setters).
- **Simple API** – `Set`, `Save`, `Restore` – that’s all you need.
- **JSON only** (for now) – human‑readable and easy to edit.

## Installation

```bash
go get github.com/Alrz7/fig
```

## Quick Start

The example below shows a typical setup – two configuration sections (`app_api` and `Appconf_second`) managed by a single handler.

```go
package main

import (
	"fmt"
	"github.com/Alrz7/fig"
)

func main() {
	// Define your configuration structs
	type api struct {
		Url  string `json:"url"`
		Port int    `json:"port"`
	}
	type car struct {
		Name  string `json:"name"`
		Model string `json:"model"`
		Year  int    `json:"year"`
	}

	// Create a handler (the main orchestrator)
	appConfig := fig.CreateNewHandler("./config/", "appConfig")

	// Create two fields linked to the handler
	apiField := appConfig.NewField("./config/", "app_api")
	carField := appConfig.NewField("./config/", "Appconf_second")

	// Prepare your variables (must be pointers)
	googleAPI := &api{}
	porche911 := &car{
		Name:  "porche911",
		Model: "porche",
		Year:  2003,
	}

	// Register them with FIG
	apiField.Set("google", googleAPI)
	carField.Set("porche911", porche911)

	// Modify the in‑memory values
	googleAPI.Port = 5050

	// Restore previous state (or create files if missing)
	// It is recommended to call Restore after all Set operations.
	appConfig.PanicRestore()

	// Now configuration is ready – use googleAPI, porche911 directly
	fmt.Printf("API port: %d\n", googleAPI.Port)
}
```

### What happens behind the scenes?

- `CreateNewHandler("./config/", "appConfig")` creates a handler that will manage a file `appConfig.json` inside `./config/`.
- `NewField` creates a field – each field has its own JSON file (`app_api.json`, `Appconf_second.json`).
- `Set(key, pointer)` stores the pointer internally. FIG tracks the value through the pointer.
- `PanicRestore()`:
  - If the JSON file exists, it reads the file and overwrites the fields of your structs (preserving any extra keys you added in code).
  - If the file does not exist, it saves the current state (so the file is created).
  - If a key exists in the file but not in your current `Set` call, FIG emits a warning (that key is ignored).

After restore, you can continue working with your configs – any changes can be saved later with `Save()`.

## API Overview

### `Handler`

- `CreateNewHandler(dir, name string) *Handler` – creates a new handler. `dir` must exist.
- `(h *Handler) NewField(dir, name string) *Field` – creates a new field linked to this handler.
- `(h *Handler) Save() error` – saves all linked fields and the handler info.
- `(h *Handler) Restore() error` – restores all linked fields.
- `(h *Handler) PanicSave()` / `PanicRestore()` – same as above, but panic on error.

### `Field`

- `CreateNewField(dir, name string) *Field` – creates an unlinked field (not recommended for most use cases).
- `(f *Field) Set(key string, newValue any)` – stores a pointer under `key`. If the field has already been restored, it also triggers an immediate save.
- `(f *Field) Pop(key string) any` – removes a key and returns its value.
- `(f *Field) Save() error` – saves only this field.
- `(f *Field) Restore() error` – restores only this field.

### `Topic` (internal map)

You normally don’t use `Topic` directly. It is a `map[string]any` that holds all key‑value pairs of a field.

## Example Output

After running the quick start example or `fig/core/fig_test.go`, three sample JSON files are created:

**`./config/app_api.json`**
```json
{
  "info": {
    "dir": "./config/",
    "name": "app_api",
    "format": ".json",
    "linked_to_Handler": true,
    "last_time_modified": "2026-03-27T22:57:15.116582276+03:30"
  },
  "data": {
    "google": { "url": "", "port": 5050 }
  }
}
```

**`./config/Appconf_second.json`**
```json
{
  "info": {
    "dir": "./config/",
    "name": "Appconf_second",
    "format": ".json",
    "linked_to_Handler": true,
    "last_time_modified": "2026-03-27T22:57:15.11705572+03:30"
  },
  "data": {
    "porche911": { "Name": "porche911", "Model": "porche", "Year": 2003 }
  }
}
```

**`./config/appConfig.json`** (handler metadata)
```json
{
  "dir": "./config/",
  "name": "appConfig",
  "format": ".json",
  "last_time_modified": "2026-03-27T22:57:15.119952079+03:30",
  "fileds_info": {
    "Appconf_second": {
      "dir": "./config/",
      "name": "Appconf_second",
      "format": "json",
      "linked_to_Handler": true,
      "last_time_modified": "2026-03-27T22:57:15.11705572+03:30"
    },
    "app_api": {
      "dir": "./config/",
      "name": "app_api",
      "format": "json",
      "linked_to_Handler": true,
      "last_time_modified": "2026-03-27T22:57:15.116582276+03:30"
    }
  }
}
```

## Important Notes

1. **Pointers only** – `Set()` expects a pointer to your variable. Passing a value will not work.
2. **Directory must exist** – FIG does not automatically create directories; you need to ensure `dir` exists beforehand.
3. **Restore order** – always call `Restore()` (or `PanicRestore()`) **after** all `Set()` calls for a field. This ensures that saved data correctly overrides your default values.
4. **Concurrency** – FIG is not thread‑safe yet, If you access the same field from multiple goroutines, add your own synchronisation.
5. **Error handling** – Use the `Panic` variants (`PanicRestore`, `PanicSave`) for quick prototyping. In production, handle errors returned by `Save()` and `Restore()`.
6. **Logger** – FIG uses an internal logger (from `github.com/Alrz7/fig/echo`). Errors are printed automatically when you use `Panic` methods.

## License

FIG is open source and available under the [MIT License](LICENSE).