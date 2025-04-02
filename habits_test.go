package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"
)

// TestDataFile is used for testing purposes, so it doesn't modify user's real data
const TestDataFile = "test_habits_data.json"

// setupTestEnv prepares a test environment
func setupTestEnv(t *testing.T) func() {
	// Save the original data file path
	originalDataFilePath := dataFilePath
	
	// Set the test data file path
	dataFilePath = TestDataFile
	
	// Delete the test data file if it exists
	os.Remove(dataFilePath)
	
	// Return a cleanup function
	return func() {
		// Delete the test data file
		os.Remove(dataFilePath)
		
		// Restore the original data file path
		dataFilePath = originalDataFilePath
	}
}

// TestLoadSaveData tests the ability to load and save data
func TestLoadSaveData(t *testing.T) {
	cleanup := setupTestEnv(t)
	defer cleanup()
	
	// Create test data
	testData := &DataFile{
		Habits: []Habit{
			{
				Name:         "Test Habit 1",
				ShortName:    "th1",
				DatesTracked: []string{"2023-01-01", "2023-01-02"},
				ReminderInfo: map[string]interface{}{},
			},
			{
				Name:         "Test Habit 2",
				ShortName:    "th2",
				DatesTracked: []string{"2023-01-01"},
				ReminderInfo: map[string]interface{}{},
			},
		},
	}
	
	// Save the test data
	err := saveData(testData)
	if err != nil {
		t.Fatalf("Failed to save test data: %v", err)
	}
	
	// Load the data back
	loadedData, err := loadData()
	if err != nil {
		t.Fatalf("Failed to load test data: %v", err)
	}
	
	// Compare the loaded data with the original test data
	if len(loadedData.Habits) != len(testData.Habits) {
		t.Errorf("Expected %d habits, got %d", len(testData.Habits), len(loadedData.Habits))
	}
	
	for i, habit := range testData.Habits {
		if i >= len(loadedData.Habits) {
			t.Errorf("Missing habit at index %d", i)
			continue
		}
		
		loadedHabit := loadedData.Habits[i]
		if habit.Name != loadedHabit.Name {
			t.Errorf("Expected habit name %s, got %s", habit.Name, loadedHabit.Name)
		}
		
		if habit.ShortName != loadedHabit.ShortName {
			t.Errorf("Expected short name %s, got %s", habit.ShortName, loadedHabit.ShortName)
		}
		
		if !reflect.DeepEqual(habit.DatesTracked, loadedHabit.DatesTracked) {
			t.Errorf("Expected dates tracked %v, got %v", habit.DatesTracked, loadedHabit.DatesTracked)
		}
	}
}

// TestAddHabit tests adding a new habit
func TestAddHabit(t *testing.T) {
	cleanup := setupTestEnv(t)
	defer cleanup()
	
	// Initial data should be empty
	df, err := loadData()
	if err != nil {
		t.Fatalf("Failed to load initial data: %v", err)
	}
	
	if len(df.Habits) != 0 {
		t.Errorf("Expected 0 habits initially, got %d", len(df.Habits))
	}
	
	// Add a habit
	commandAdd([]string{"Test", "Habit"}, df)
	
	// Load data again to verify
	df, err = loadData()
	if err != nil {
		t.Fatalf("Failed to load data after adding habit: %v", err)
	}
	
	// Verify the habit was added
	if len(df.Habits) != 1 {
		t.Errorf("Expected 1 habit after adding, got %d", len(df.Habits))
	}
	
	if df.Habits[0].Name != "Test Habit" {
		t.Errorf("Expected habit name 'Test Habit', got '%s'", df.Habits[0].Name)
	}
}

// TestMarkHabitDone tests marking a habit as done
func TestMarkHabitDone(t *testing.T) {
	cleanup := setupTestEnv(t)
	defer cleanup()
	
	// Create initial data with a habit
	df := &DataFile{
		Habits: []Habit{
			{
				Name:         "Test Habit",
				ShortName:    "th",
				DatesTracked: []string{},
				ReminderInfo: map[string]interface{}{},
			},
		},
	}
	
	// Save the initial data
	err := saveData(df)
	if err != nil {
		t.Fatalf("Failed to save initial data: %v", err)
	}
	
	// Mark the habit as done
	commandDone([]string{"1"}, df)
	
	// Load data again to verify
	df, err = loadData()
	if err != nil {
		t.Fatalf("Failed to load data after marking habit done: %v", err)
	}
	
	// Verify the habit was marked as done
	if len(df.Habits[0].DatesTracked) != 1 {
		t.Errorf("Expected 1 date tracked, got %d", len(df.Habits[0].DatesTracked))
	}
	
	// The date should be today in YYYY-MM-DD format
	today := time.Now().Format("2006-01-02")
	if df.Habits[0].DatesTracked[0] != today {
		t.Errorf("Expected date tracked to be %s, got %s", today, df.Habits[0].DatesTracked[0])
	}
}

// TestRemoveHabit tests removing a habit
func TestRemoveHabit(t *testing.T) {
	cleanup := setupTestEnv(t)
	defer cleanup()
	
	// Create initial data with two habits
	df := &DataFile{
		Habits: []Habit{
			{
				Name:         "Test Habit 1",
				ShortName:    "th1",
				DatesTracked: []string{},
				ReminderInfo: map[string]interface{}{},
			},
			{
				Name:         "Test Habit 2",
				ShortName:    "th2",
				DatesTracked: []string{},
				ReminderInfo: map[string]interface{}{},
			},
		},
	}
	
	// Save the initial data
	err := saveData(df)
	if err != nil {
		t.Fatalf("Failed to save initial data: %v", err)
	}
	
	// Remove the first habit
	commandDelete([]string{"1"}, df)
	
	// Mock user input for the delete command confirmation
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	
	go func() {
		fmt.Fprintln(w, "y") // Confirm deletion
		w.Close()
	}()
	
	// Delete the habit
	commandDelete([]string{"1"}, df)
	
	// Restore stdin
	os.Stdin = oldStdin
	
	// Load data again to verify
	df, err = loadData()
	if err != nil {
		t.Fatalf("Failed to load data after removing habit: %v", err)
	}
	
	// Verify the habit was removed
	if len(df.Habits) != 1 {
		t.Errorf("Expected 1 habit after removal, got %d", len(df.Habits))
	}
	
	if len(df.Habits) > 0 && df.Habits[0].Name != "Test Habit 2" {
		t.Errorf("Expected remaining habit to be 'Test Habit 2', got '%s'", df.Habits[0].Name)
	}
}

// TestEditHabit tests editing a habit's name
func TestEditHabit(t *testing.T) {
	cleanup := setupTestEnv(t)
	defer cleanup()
	
	// Create initial data with a habit
	df := &DataFile{
		Habits: []Habit{
			{
				Name:         "Test Habit",
				ShortName:    "th",
				DatesTracked: []string{},
				ReminderInfo: map[string]interface{}{},
			},
		},
	}
	
	// Save the initial data
	err := saveData(df)
	if err != nil {
		t.Fatalf("Failed to save initial data: %v", err)
	}
	
	// Edit the habit using the flags
	commandEdit([]string{"1", "--name", "Edited Test Habit"}, df)
	
	// Load data again to verify
	df, err = loadData()
	if err != nil {
		t.Fatalf("Failed to load data after editing habit: %v", err)
	}
	
	// Verify the habit name was edited
	if df.Habits[0].Name != "Edited Test Habit" {
		t.Errorf("Expected habit name to be 'Edited Test Habit', got '%s'", df.Habits[0].Name)
	}
}

// TestImportExport tests the import and export functionality
func TestImportExport(t *testing.T) {
	cleanup := setupTestEnv(t)
	defer cleanup()
	
	// Create test data
	df := &DataFile{
		Habits: []Habit{
			{
				Name:         "Test Habit 1",
				ShortName:    "th1",
				DatesTracked: []string{"2023-01-01"},
				ReminderInfo: map[string]interface{}{},
			},
			{
				Name:         "Test Habit 2",
				ShortName:    "th2",
				DatesTracked: []string{"2023-01-02"},
				ReminderInfo: map[string]interface{}{},
			},
		},
	}
	
	// Save the test data
	err := saveData(df)
	if err != nil {
		t.Fatalf("Failed to save test data: %v", err)
	}
	
	// Export the data with --file flag
	exportFile := "test_export_for_test.json"
	commandExport([]string{"--file", exportFile}, df)
	
	// Clean the data file to simulate a fresh state
	os.Remove(dataFilePath)
	
	// Load the data to verify it's empty
	df, err = loadData()
	if err != nil {
		t.Fatalf("Failed to load clean data: %v", err)
	}
	
	if len(df.Habits) != 0 {
		t.Errorf("Expected 0 habits after cleaning, got %d", len(df.Habits))
	}
	
	// Import the data with --file flag
	commandImport([]string{"--file", exportFile}, df)
	
	// Load data again to verify
	df, err = loadData()
	if err != nil {
		t.Fatalf("Failed to load data after import: %v", err)
	}
	
	// Verify the habits were imported
	if len(df.Habits) != 2 {
		t.Errorf("Expected 2 habits after import, got %d", len(df.Habits))
	}
	
	// Clean up the export file
	os.Remove(exportFile)
}

// TestMain sets up and runs the tests
func TestMain(m *testing.M) {
	// Run the tests
	exitCode := m.Run()
	
	// Clean up any leftover test files
	os.Remove(TestDataFile)
	os.Remove("test_export_for_test.json")
	
	os.Exit(exitCode)
} 