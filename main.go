package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)


type EventObject struct {
	TargetState string `json:"targetState"`
	// ignore other keys - just care about target state
}

type State struct {
	Type string `json:"stateType"`
	OnEvent map[string][]EventObject `json:"onEvent"`
}

type StateMachine struct {
	InitialState string         `json:"initialState"`
	States     map[string]State `json:"states"`
}


func generateDOT(sm StateMachine) string {
	var dot strings.Builder
	
	dot.WriteString("digraph NavigationGraph {\n")
	dot.WriteString("  rankdir=LR; // Layout from Left to Right (e.g., A -> B -> C)\n")
	dot.WriteString("  node [shape=box, style=\"filled,rounded\", fontname=\"Inter\"];\n\n")

	for name, stateData := range sm.States {

		var color string
		if stateData.Type != "screen" {
			continue
		}
		
		if name == sm.InitialState {
			color = "#90EE90"
		} else {
			color = "#5594f1ff"
		}
		
		label := fmt.Sprintf("%s", name)
		dot.WriteString(fmt.Sprintf("  \"%s\" [label=\"%s\", fillcolor=\"%s\"];\n", name, label, color))
	}
	


	for source, stateData := range sm.States {
		for eventName, eventArr := range stateData.OnEvent {
        
            for _, eventObj := range eventArr {
                target := eventObj.TargetState 
			
                dot.WriteString(fmt.Sprintf("  %s -> %s [label=\"%s\", color=\"#0047AB\"];\n", source, target, eventName))
            }
		}
	}


	dot.WriteString("}\n")
	return dot.String()
}

func getOutputFormat(outputPath string) string {
	ext := strings.ToLower(filepath.Ext(outputPath))
	if len(ext) > 0 {
		return ext[1:] 
	}
	return "png" 
}

func main() {
	inputPath := flag.String("f", "", "Path to the input JSON file containing the state machine definition.")
	outputPath := flag.String("o", "navigation_graph.png", "Output path for the generated image file (e.g., graph.svg or graph.png).")
	flag.Parse() 
	if *inputPath == "" {
		fmt.Println("Error: Missing required flag -f. Please provide the path to the JSON file.")
		flag.Usage()
		os.Exit(1)
	}

	jsonContent, err := ioutil.ReadFile(*inputPath)
	if err != nil {
		fmt.Printf("Error reading input file %s: %v\n", *inputPath, err)
		os.Exit(1)
	}

	var sm StateMachine 
	if err := json.Unmarshal(jsonContent, &sm); err != nil {
		fmt.Printf("Error parsing JSON state machine: %v\n", err)
		os.Exit(1)
	}

	dotSource := generateDOT(sm)

	// 4. Execute Graphviz 'dot' command

	dotFilePath := filepath.Join(os.TempDir(), "temp_nav_graph.dot")
	if err := ioutil.WriteFile(dotFilePath, []byte(dotSource), 0644); err != nil {
		fmt.Printf("Error writing temporary DOT file: %v\n", err)
		os.Exit(1)
	}
	defer os.Remove(dotFilePath) 

	outputFormat := getOutputFormat(*outputPath)
    
	cmd := exec.Command("dot", fmt.Sprintf("-T%s", outputFormat), dotFilePath, "-o", *outputPath)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("\nGraphviz execution failed! (Is the 'dot' command installed?)\n")
		fmt.Printf("Error: %v\n", err)
		fmt.Printf("Graphviz Output:\n%s\n", string(output))
		os.Exit(1)
	}

	fmt.Printf("Successfully generated graph: %s\n", *outputPath)
}
