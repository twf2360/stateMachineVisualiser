# State Machine Graph Generator (JSON to DOT)

This project provides a Go utility that takes a custom-formatted JSON file defining a state machine and converts it into a Graphviz DOT language file. It then executes the `dot` command (from Graphviz) to automatically generate a visual representation (e.g., a PNG or SVG image) of the state machine's navigational flow.

## Prerequisites (For End Users)


### 1. Graphviz (Mandatory)

Graphviz is the visualization software that reads the DOT language output by the program and draws the final graph image. The `dot` command must be accessible from your shell's PATH.

* **Installation (macOS/Linux):**

macOS (Homebrew)

brew install graphviz

Debian/Ubuntu (apt)

sudo apt install graphviz


## Usage

Assuming you have downloaded or build the executable and your state machine definition is `demo.json`.

### 1. Executing 
Run the executable, specifying the input file (`-f`) and the desired output file (`-o`).

| Flag | Description | Required | Example | 
| :--- | :--- | :--- | :--- | 
| **-f** | Path to the input JSON definition file. | Yes | `-f my_machine.json` | 
| **-o** | Path and filename for the output image. Determines format (e.g., `.png`, `.svg`). | No | `-o output_flow.svg` | 

**Example Command**

The provided demo.json when ran with `./main -f  demo.json -o demo.svg` will generate the provided demo.svg

## Building the Executable

If you need to compile the program yourself, you will need the Go toolchain. Refer to the separate **`BUILD.md`** file for detailed instructions.

## Input JSON Structure

The tool expects a specific, nested structure to correctly map states, events, and transitions.

### Key Fields

| Field Name | Location | Description | 
| :--- | :--- | :--- | 
| **`initialState`** | Root object | The name of the first state (node) in the graph. | 
| **`states`** | Root object | A map where keys are the state names (nodes). | 
| **`stateType`** | Inside a State object | The type of the state (e.g., `screen`, `activity`). Used for styling the node. | 
| **`onEvent`** | Inside a State object | A map where keys are event names (edge labels). | 
| **`targetState`** | Inside the nested Event object array | The name of the destination state for the transition. | 
