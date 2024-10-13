# Email Template Project

This project generates a dynamic HTML email template using Go Templating. The template includes a logo, dynamic list of user cards, and a footer.

## Structure

- `/templates/test-email-template.html`: The HTML template for the email.
- `/main.go`: The Go file for rendering the template with dynamic data.
- `/data/emailData.json`: (Optional) JSON data file with dynamic data.

## How to Run

1. Clone the repository:
git clone git@github.com:kandyba/test-email-go-templating.git

2. Navigate to the project directory:
cd test-email-go-templating

3. Run the Go program:
go run main.go
