package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
	"gorgonia.org/vecf64"
)

func main() {
	// Read data from CSV file
	X, y := readCSV("Data.csv")

	// Convert y to one-hot encoding
	// Assuming y is already categorical labels

	// Split data into training and testing sets
	XTrain, XTest, yTrain, yTest := splitData(X, y, 0.2)

	// Define network parameters
	inputSize := len(XTrain[0])
	hiddenSize := 64
	outputSize := len(yTrain[0])

	// Create Gorgonia graph
	g := gorgonia.NewGraph()

	// Define input and output nodes
	X := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(-1, inputSize), gorgonia.WithValue(tensor.New(tensor.Of(tensor.Float64), tensor.WithShape(len(XTrain), inputSize), tensor.WithBacking(XTrain))))
	y := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(-1, outputSize), gorgonia.WithValue(tensor.New(tensor.Of(tensor.Float64), tensor.WithShape(len(yTrain), outputSize), tensor.WithBacking(yTrain))))

	// Define model parameters
	wHidden := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(inputSize, hiddenSize), gorgonia.WithInit(gorgonia.GlorotU(1)))
	bHidden := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(1, hiddenSize), gorgonia.WithInit(gorgonia.Zeroes()))

	wOutput := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(hiddenSize, outputSize), gorgonia.WithInit(gorgonia.GlorotU(1)))
	bOutput := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(1, outputSize), gorgonia.WithInit(gorgonia.Zeroes()))

	// Define operations for forward pass
	hidden := gorgonia.Must(gorgonia.Add(gorgonia.Must(gorgonia.Mul(X, wHidden)), bHidden))
	hiddenActivation := gorgonia.Must(gorgonia.Rectify(hidden))

	output := gorgonia.Must(gorgonia.Add(gorgonia.Must(gorgonia.Mul(hiddenActivation, wOutput)), bOutput))
	prediction := gorgonia.Must(gorgonia.SoftMax(output))

	// Define loss function
	loss := gorgonia.Must(gorgonia.Mean(gorgonia.Must(gorgonia.Neg(gorgonia.Must(gorgonia.Mul(y, gorgonia.Must(gorgonia.Log(prediction))))))))

	// Define the solver
	solver := gorgonia.NewRMSPropSolver(gorgonia.WithLearnRate(0.001))

	// Define the gradients
	grads, err := gorgonia.Gradient(loss, wHidden, bHidden, wOutput, bOutput)
	if err != nil {
		log.Fatal(err)
	}

	// Define the trainer
	m := gorgonia.NewTapeMachine(g, gorgonia.BindDualValues(wHidden, bHidden, wOutput, bOutput))
	defer m.Close()

	// Train the model
	epochs := 50
	for epoch := 0; epoch < epochs; epoch++ {
		if err := m.RunAll(); err != nil {
			log.Fatal(err)
		}

		if err := solver.Step(gorgonia.NodesToValueGrads(grads)); err != nil {
			log.Fatal(err)
		}

		// Clear the gradients
		for _, grad := range grads {
			grad.Zero()
		}

		// Print the loss
		fmt.Printf("Epoch %d/%d, Loss: %.4f\n", epoch+1, epochs, loss.Value().Data().(float64))
	}

	// Evaluate the model
	correct := 0
	total := 0
	for i, xTest := range XTest {
		yTrue := yTest[i]
		yPred := predict(m, xTest)
		if yTrue == yPred {
			correct++
		}
		total++
	}

	// Print accuracy
	accuracy := float64(correct) / float64(total) * 100
	fmt.Printf("Accuracy: %.2f%%\n", accuracy)
}

// Function to read data from CSV file
func readCSV(filename string) ([][]float64, [][]float64) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var X [][]float64
	var y [][]float64

	for _, record := range records {
		UltrasonicValue, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		buttonZeroState, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		buttonOneState, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		buttonTwoState, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}
		// Add more feature parsing as needed

		label, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			log.Fatal(err)
		}

		X = append(X, []float64{UltrasonicValue, buttonZeroState, buttonOneState, buttonTwoState})
		y = append(y, []float64{label})
	}

	return X, y
}

// Function to split data into training and testing sets
func splitData(X, y [][]float64, testRatio float64) ([][]float64, [][]float64, [][]float64, [][]float64) {
	rand.Seed(time.Now().UnixNano())
	n := len(X)
	nTest := int(float64(n) * testRatio)
	idx := rand.Perm(n)
	XTrain := make([][]float64, n-nTest)
	yTrain := make([][]float64, n-nTest)
	XTest := make([][]float64, nTest)
	yTest := make([][]float64, nTest)
	for i, j := range idx[:n-nTest] {
		XTrain[i] = X[j]
		yTrain[i] = y[j]
	}
	for i, j := range idx[n-nTest:] {
		XTest[i] = X[j]
		yTest[i] = y[j]
	}
	return XTrain, XTest, yTrain, yTest
}

// Function to predict labels using the trained model
func predict(m *gorgonia.VM, x []float64) int {
	gorgonia.Let(gorgonia.NewTensor(m, tensor.Float64, 2, gorgonia.WithShape(1, len(x)), gorgonia.WithValue(tensor.New(tensor.Of(tensor.Float64), tensor.WithShape(1, len(x)), tensor.WithBacking(x)))), gorgonia.NodeFromGraph(m.Graph().Roots()[0].Children()[0]))
	if err := m.RunAll(); err != nil {
		log.Fatal(err)
	}
	prediction := m.Graph().Roots()[0].Value().Data().([]float64)
	return vecf64.Argmax(prediction)
}
