# fantastic-succotash

This is a simple application to send Mpesa STK Push requests using the Daraja API to Kenyan Members of Parliament.

## Installation

To run this application, you need to have Go installed on your machine. You can download and install Go from the [official website](https://go.dev/).

1. Clone the repository

   ```bash
   git clone https://github.com/0x6flab/fantastic-succotash
   ```

2. Run `go mod download` to download the dependencies

   ```bash
   go mod download
   ```

3. Create a `.env` file in the root directory and add the following environment variables

   ```bash
   MPESA_CONSUMER_KEY=
   MPESA_CONSUMER_SECRET=
   MPESA_PASS_KEY=
   ```

4. Run the application

   ```bash
   go run main.go
   ```
