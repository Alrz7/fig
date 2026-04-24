# FIG (Flexible In-memory ConFIGuration)

FIG is a simple, lightweight and flexible Go library designed to simplify the management of configuration data in your applications. It allows you to define, load, save, and synchronize configuration structures using JSON files with ease. Whether you're managing simple key-value pairs or complex nested data, FIG provides a robust solution.

## ✨ Overview

FIG empowers you to manage your application's configuration in a structured and dynamic way. It focuses on:

-   **JSON-Based Configuration:** Define your configuration directly in JSON files.
-   **Type-Safe Structures:** Map your JSON data to Go structs for type safety and ease of use.
-   **Flexible Handlers:** Create multiple independent configuration handlers for different parts of your application.
-   **Dynamic Data Management:** Load, update, and save configuration values on-the-fly.
-   **Pointer-Based Updates:** Pass configuration values as pointers, allowing for real-time modifications.
-   **Automatic Synchronization:** Ensures your configuration is saved and restored reliably.

## 🚀 Getting Started

FIG is designed to be intuitive and easy to integrate into your Go projects. Follow these steps to start managing your configuration:

### 1. Create a New Handler

Initialize a new configuration handler by specifying the directory where your JSON configuration files will be stored and the name of the main configuration file.

```go
// "./" is the directory, "testConfig" is the config file name
appConfig := core.CreateNewHandeler("./", "testConfig")
```

### 2. Define Your Configuration Structure

Create Go structs that mirror the structure of your JSON configuration. Use struct tags (`json:"fieldName"`) to map JSON keys to your Go struct fields.

```go
// Example: A driver's configuration
type driver struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
	Job    string `json:"job"`
}

// Example: A company's configuration
type Company struct {
	ID        int               `json:"id"`
	Name      string            `json:"name"`
	Founded   string            `json:"founded"`
	IsActive  bool              `json:"is_active"`
	Employees []driver          `json:"employees"`
	Revenue   float64           `json:"revenue"`
	Departments map[string]int    `json:"departments"`
	Headquarters struct {
		City  string `json:"city"`
		State string `json:"state"`
	} `json:"headquarters"`
}
```

### 3. Pre-Declare Your Instances (Optional)

You can declare instances of your configuration structs. These can be initially empty or populated with default values.

```go
// Pre-declare instances (can be empty or have default values)
miles := driver{
	Name:   "Miles",
	Age:    23,
	Gender: "male",
	Job:    "driver",
}

Uber := Company{
	ID:       8403453853045,
	Name:     "Uber",
	Founded:  "2001/12/july",
	IsActive: true,
	Employees: []driver{miles},
	Revenue:  352094284,
	Departments: map[string]int{"Florida": 4, "Newyork": 3, "vegas": 8},
	Headquarters: struct {
		City  string `json:"city"`
		State string `json:"state"`
	}{City: "miami", State: "Florida"},
}

// Let's assume 'porche911' and 'name' (for CEO) are also defined elsewhere
// For example:
// porche911 := SomeOtherStruct{ ... }
// name := "Alice" // simple string example
```

### 4. Add Instances as Pointers to the Handler

Use the `Set` method of your handler to register your configuration instances. Crucially, you must pass them as **pointers**. This allows FIG to track and manage changes to these variables.

```go
initConfig := func() {
	// Add your pre-declared instances as pointers
	appConfig.Set("Uber", &Uber)
	appConfig.Set("porche911", &porche911) // Assuming porche911 is defined
	appConfig.Set("CEO", &name)          // Assuming name is defined as a string pointer or similar
}

// Call the init function to set up your configurations
initConfig()
```

**Important:** The `defer appConfig.PanicRestore(appConfig)` line is essential. It ensures that your handler automatically saves and synchronizes the data after the first load and whenever operations are completed. This provides a robust mechanism for persisting changes.

### 5. Modify and Save Configurations

You can pass your configuration data (as pointers) anywhere in your application and modify them. FIG will only save these changes to the JSON file when you explicitly call the `Save()` function on your handler.

```go
// Example: Modifying the Uber configuration
func updateUberRevenue(newRevenue float64) {
	// Retrieve the pointer to the Uber configuration
	uberConfigPtr, err := appConfig.Get("Uber")
	if err != nil {
		// Handle error
		return
	}

	// Assert the type to your Company struct (or the correct type)
	// Use type assertion or a safer method if available in your core library
	if company, ok := uberConfigPtr.(*Company); ok {
		company.Revenue = newRevenue
		fmt.Printf("Uber revenue updated to: %.2f\n", company.Revenue)

		// Save the changes back to the JSON file
		if err := appConfig.Save(); err != nil {
			// Handle save error
			fmt.Printf("Error saving config: %v\n", err)
		} else {
			fmt.Println("Configuration saved successfully.")
		}
	} else {
		fmt.Println("Error: Could not assert type for Uber configuration.")
	}
}

// Call the update function
// updateUberRevenue(400000000)
```

## 💡 Ideas and Enhancements

FIG is already quite powerful, but here are a few ideas for potential enhancements or alternative use cases:

*   **Advanced Validation:** Implement built-in validation rules for configuration fields (e.g., minimum/maximum values, required fields, regex patterns) that can be triggered during `Set` or `Save`.
*   **Environment-Specific Configurations:** Support loading configurations based on the environment (e.g., `config.dev.json`, `config.prod.json`) and merging them.
*   **Hot Reloading:** Introduce a mechanism for automatically detecting changes in the config file and reloading them into the application without requiring a restart.
*   **Custom Data Types:** Extend support for custom data types beyond basic JSON primitives, perhaps with custom marshaling/unmarshaling logic.
*   **Configuration Templating:** Allow for variables and environment expansion within the JSON configuration itself (e.g., using Go's `text/template` or `html/template` packages).
*   **Decryption/Encryption:** Integrate support for encrypting sensitive configuration values (like API keys or passwords) within the JSON file and decrypting them upon loading.
*   **Centralized Configuration Service Integration:** Adapt FIG to work with or fetch configurations from external services like Consul, etcd, or AWS Parameter Store.
*   **Monitoring and Auditing:** Log changes made to the configuration and who made them, especially in distributed systems.

## 📄 License

This project is licensed under the [MIT License](LICENSE) - see the [LICENSE](LICENSE) file for details.
