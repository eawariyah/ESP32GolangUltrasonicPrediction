<?php
// Get the request body
$request_body = file_get_contents('php://input');

// Decode the JSON data
$data = json_decode($request_body, true);

if (
    isset($data["distance"]) && 
    isset($data["buttonZeroState"]) &&
    isset($data["buttonOneState"]) &&
    isset($data["buttonTwoState"]) &&
    isset($data["buttonThreeState"])
) {
    // Add timestamp to the data
    $data["timestamp"] = date(DATE_RFC2822);

    // Insert the data into the CSV file as a new row
    $file = fopen("data.csv", "a");
    fputcsv($file, $data);
    fclose($file);

    // Print result
    echo "Data Inserted Successfully!";
} else {
    echo "Waiting for data to insert...";
}
?>
