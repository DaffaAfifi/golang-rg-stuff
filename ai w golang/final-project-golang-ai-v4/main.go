package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type AIModelConnector struct {
	Client *http.Client
}

type Inputs struct {
	Table map[string][]string `json:"table"`
	Query string              `json:"query"`
}

type Response struct {
	Answer      string   `json:"answer"`
	Coordinates [][]int  `json:"coordinates"`
	Cells       []string `json:"cells"`
	Aggregator  string   `json:"aggregator"`
}

func CsvToSlice(data string) (map[string][]string, error) {
	r := csv.NewReader(strings.NewReader(data))
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	result := make(map[string][]string)
	headers := records[0]
	for i := 1; i < len(records); i++ {
		for j, value := range records[i] {
			result[headers[j]] = append(result[headers[j]], value)
		}
	}

	return result, nil
}

func (c *AIModelConnector) ConnectAIModel(payload interface{}, token string) (Response, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return Response{}, err
	}

	modelURL := "https://api-inference.huggingface.co/models/google/tapas-base-finetuned-wtq"
	req, err := http.NewRequest("POST", modelURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return Response{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Response{}, err
	}

	return response, nil
}

func (c *AIModelConnector) generateRecommendation(query, token string) (string, error) {
	apiURL := "https://api-inference.huggingface.co/models/openai-community/gpt2"

	payload := struct {
		Inputs string `json:"inputs"`
	}{
		Inputs: query,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response []struct {
		GeneratedText string `json:"generated_text"`
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	if len(response) > 0 {
		return response[0].GeneratedText, nil
	}

	return "", nil
}

func main() {
	// Read csv file
	file, err := os.Open("data-series.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvData, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Convert csv to string
	var csvDataStr string
	for _, record := range csvData {
		csvDataStr += strings.Join(record, ",") + "\n"
	}

	// Convert csv string to map
	tableData, err := CsvToSlice(csvDataStr)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize connector
	client := &http.Client{}
	connector := AIModelConnector{Client: client}

	token := os.Getenv("HUGGINGFACE_TOKEN")
	if token == "" {
		log.Fatal("HUGGINGFACE_TOKEN is not set")
	}

	// Interactive CLI
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Welcome to the AI-Powered Smart Home Energy Management System!")
		fmt.Println("Please select an option:")
		fmt.Println("1. Use Tapas Model")
		fmt.Println("2. Use GPT-2 Model")
		fmt.Println("3. Exit")
		fmt.Println("==============================")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			fmt.Println("Enter your query for the Tapas Model (or 'back' to return to the main menu):")
			for {
				query, _ := reader.ReadString('\n')
				query = strings.TrimSpace(query)

				if query == "back" {
					fmt.Println("==============================")
					break
				}

				payload := Inputs{
					Table: tableData,
					Query: query,
				}

				response, err := connector.ConnectAIModel(payload, token)
				if err != nil {
					log.Println(err)
					continue
				}

				fmt.Printf("Answer: %s\n", response.Answer)
				fmt.Printf("Coordinates: %v\n", response.Coordinates)
				fmt.Printf("Cells: %v\n", response.Cells)
				fmt.Printf("Aggregator: %s\n", response.Aggregator)
				fmt.Println("Enter your next query (or 'back' to return to the main menu):")
				fmt.Println("==============================")
			}
		case "2":
			fmt.Println("Enter your query for the GPT-2 Model (or 'back' to return to the main menu):")
			for {
				query, _ := reader.ReadString('\n')
				query = strings.TrimSpace(query)

				if query == "back" {
					fmt.Println("==============================")
					break
				}

				recommendation, err := connector.generateRecommendation(query, token)
				if err != nil {
					log.Println(err)
					continue
				}

				fmt.Printf("Recommendation: %s\n", recommendation)
				fmt.Println("Enter your next query (or 'back' to return to the main menu):")
				fmt.Println("==============================")
			}
		case "3":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
